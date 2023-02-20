package social

import (
	"bench/dal/query"
	"context"
	"fmt"
	"time"
)

func sendMsg(ctx context.Context, uid int64) {
	key := fmt.Sprintf(lockkey, uid)
	if xr.SetNX(ctx, key, time.Now().UnixMilli(), time.Second).Val() {
		countch <- &syncCount{uid: uid}
	}
}

func followOne(ctx context.Context, uid, touid int64) error {
	err := q.Transaction(func(tx *query.Query) error {
		if err := q.WithContext(ctx).Follow.InsertFollow(uid, touid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).Follower.InsertFollower(touid, uid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).FollowCount.IncrFollowCount(uid); err != nil {
			return err
		}
		return q.WithContext(ctx).ExtraFollower.InsertFollower(touid)
	})
	if err != nil {
		return err
	}
	sendMsg(ctx, touid)
	return nil
}

func unfollowOne(ctx context.Context, uid, touid int64) error {
	err := q.Transaction(func(tx *query.Query) error {
		if err := q.WithContext(ctx).Follow.DeleteFollow(uid, touid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).Follower.DeleteFollower(touid, uid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).FollowCount.DecrFollowCount(uid); err != nil {
			return err
		}
		return q.WithContext(ctx).ExtraFollower.DeleteFollower(touid)
	})
	if err != nil {
		return err
	}
	sendMsg(ctx, touid)
	return nil
}

func follow(ctx context.Context, uid, touid int64) error {
	return q.Transaction(func(tx *query.Query) error {
		if err := q.WithContext(ctx).Follow.InsertFollow(uid, touid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).Follower.InsertFollower(touid, uid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).FollowCount.IncrFollowCount(uid); err != nil {
			return err
		}
		return q.WithContext(ctx).FollowCount.IncrByFollowerCount(touid, 1)
	})
}

func unfollow(ctx context.Context, uid, touid int64) error {
	return q.Transaction(func(tx *query.Query) error {
		if err := q.WithContext(ctx).Follow.DeleteFollow(uid, touid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).Follower.DeleteFollower(touid, uid); err != nil {
			return err
		}
		if err := q.WithContext(ctx).FollowCount.DecrFollowCount(uid); err != nil {
			return err
		}
		return q.WithContext(ctx).FollowCount.DecrByFollowerCount(touid, 1)
	})
}

func jobConsumer(ctx context.Context, ch <-chan *asyncFollow) {
	defer jobwg.Done()
	for {
		select {
		case val, ok := <-ch:
			if ok {
				ctx, cancel := context.WithTimeout(ctx, time.Second)
				if err := follow(ctx, val.uid, val.touid); err != nil {
					logger.Println(val, err)
				}
				cancel()
			} else {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func newJobConsumer(ctx context.Context, n int, ch <-chan *asyncFollow) {
	jobwg.Add(n)
	for i := 0; i < n; i++ {
		go jobConsumer(ctx, ch)
	}
}
