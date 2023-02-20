package social

import (
	"bench/dal/query"
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func syncRecordByUID(ctx context.Context, uid int64) error {
	var cnt int64
	err := q.Transaction(func(tx *query.Query) error {
		rows, err := q.WithContext(ctx).ExtraFollower.FindUnSyncRecordByUID(uid)
		if err != nil || len(rows) == 0 {
			return err
		}
		ids := make([]int64, len(rows))
		for k, v := range rows {
			ids[k] = v.ID
			if v.Stats == 0 {
				cnt++
			} else {
				cnt--
			}
		}
		if cnt > 0 {
			if err = tx.WithContext(ctx).FollowCount.IncrByFollowerCount(uid, cnt); err != nil {
				return err
			}
		} else if cnt < 0 {
			if err = tx.WithContext(ctx).FollowCount.DecrByFollowerCount(uid, -cnt); err != nil {
				return err
			}
		}
		return tx.WithContext(ctx).ExtraFollower.DeleteRecord(ids)
	})
	if err != nil {
		return err
	}
	key := fmt.Sprintf(redisKeyHRelationCount, uid)
	return xr.HIncrByXEX(ctx, key, followerField, cnt, time.Hour*8)
}

func taskConsumer(ctx context.Context, ch <-chan *syncCount) {
	defer taskwg.Done()
	for {
		select {
		case val, ok := <-ch:
			if ok {
				if err := syncRecordByUID(context.TODO(), val.uid); err != nil && err != redis.Nil {
					logger.Println(val.uid, err)
				}
			} else {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func newTaskConsumer(ctx context.Context, n int, ch <-chan *syncCount) {
	taskwg.Add(n)
	for i := 0; i < n; i++ {
		go taskConsumer(ctx, ch)
	}
}
