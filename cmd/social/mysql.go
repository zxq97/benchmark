package social

import (
	"bench/dal/model"
	"bench/dal/query"
	"context"
	"time"
)

func follow(ctx context.Context, uid, touid int64) error {
	return Q.Transaction(func(tx *query.Query) error {
		if err := tx.Follow.WithContext(ctx).InsertFollow(uid, touid); err != nil {
			return err
		}
		if err := tx.Follower.WithContext(ctx).InsertFollower(touid, uid); err != nil {
			return err
		}
		if err := tx.FollowCount.WithContext(ctx).IncrFollowCount(uid); err != nil {
			return err
		}
		return tx.FollowCount.WithContext(ctx).IncrByFollowerCount(touid, 1)
	})
}

func followOne(ctx context.Context, uid, touid int64) error {
	return Q.Transaction(func(tx *query.Query) error {
		if err := tx.Follow.WithContext(ctx).InsertFollow(uid, touid); err != nil {
			return err
		}
		if err := tx.Follower.WithContext(ctx).InsertFollower(touid, uid); err != nil {
			return err
		}
		if err := tx.FollowCount.WithContext(ctx).IncrFollowCount(uid); err != nil {
			return err
		}
		return tx.ExtraFollower.WithContext(ctx).InsertFollower(touid)
	})
}

func unfollow(ctx context.Context, uid, touid int64) error {
	return Q.Transaction(func(tx *query.Query) error {
		if err := tx.Follow.WithContext(ctx).DeleteFollow(uid, touid); err != nil {
			return err
		}
		if err := tx.Follower.WithContext(ctx).DeleteFollower(touid, uid); err != nil {
			return err
		}
		if err := tx.FollowCount.WithContext(ctx).DecrFollowCount(uid); err != nil {
			return err
		}
		return tx.FollowCount.WithContext(ctx).DecrByFollowerCount(touid, 1)
	})
}

func getFollowCount(ctx context.Context, uids []int64) (map[int64]*model.FollowCount, error) {
	res, err := Q.WithContext(ctx).FollowCount.FindUsersRelationCount(uids)
	if err != nil {
		return nil, err
	}
	m := make(map[int64]*model.FollowCount, len(uids))
	for i := range res {
		m[res[i].UID] = res[i]
	}
	return m, nil
}

func syncCount(ctx context.Context, uid int64) error {
	return Q.Transaction(func(tx *query.Query) error {
		list, err := tx.ExtraFollower.WithContext(ctx).FindUnSyncRecordByUID(uid, 1000)
		if err != nil {
			return err
		}
		var amount int64
		ids := make([]int64, len(list))
		for k := range list {
			ids[k] = list[k].ID
			if list[k].Stats == 0 {
				amount++
			} else {
				amount--
			}
		}
		Logger.Println(len(ids), amount)
		if amount != 0 {
			if err = tx.FollowCount.WithContext(ctx).IncrByFollowerCount(uid, amount); err != nil {
				return err
			}
		}
		return tx.ExtraFollower.WithContext(ctx).DeleteRecord(ids)
	})
}

func CronSync(ctx context.Context) error {
	list, err := Q.ExtraFollower.WithContext(ctx).FindUnSyncRecord(10000)
	if err != nil || len(list) == 0 {
		return err
	}
	for i := 0; i < len(list); i += 1000 {
		ids := make([]int64, 0, 1000)
		m := make(map[int64]int64, 1000)
		left := i
		right := left + 1000
		if right > len(list) {
			right = len(list)
		}
		for j := left; j < right; j++ {
			ids = append(ids, list[j].ID)
			if list[j].Stats == 0 {
				m[list[j].UID]++
			} else {
				m[list[j].UID]--
			}
		}
		txCtc, cancel := context.WithTimeout(ctx, time.Second*3)
		if err = Q.Transaction(func(tx *query.Query) error {
			for k, v := range m {
				if v == 0 {
					continue
				}
				if err = tx.FollowCount.WithContext(txCtc).IncrByFollowerCount(k, v); err != nil {
					return err
				}
			}
			return tx.ExtraFollower.WithContext(txCtc).DeleteRecord(ids)
		}); err != nil {
			Logger.Println("cron sync", ids, m, err)
		}
		cancel()
	}
	return nil
}
