package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/zxq97/gokit/pkg/cache/xredis"
	"github.com/zxq97/gokit/pkg/cast"
)

const (
	redisKeyZUserFollower = "rla_foe_%d" // uid
)

var (
	xr *xredis.XRedis
)

func getFollowerList(ctx context.Context, uid, lastid, offset int64) ([]*bizdata.FollowItem, error) {
	key := fmt.Sprintf(redisKeyZUserFollower, uid)
	var (
		zs  []redis.Z
		err error
	)
	if lastid == 0 {
		zs, err = c.redis.ZRevRangeWithScores(ctx, key, 0, offset-1).Result()
	} else {
		zs, err = c.redis.ZRevRangeByMemberWithScores(ctx, key, lastid, offset)
	}
	if err != nil {
		return nil, err
	}
	list := make([]*bizdata.FollowItem, len(zs))
	for k, z := range zs {
		list[k] = &bizdata.FollowItem{
			ToUid:      cast.ParseInt(z.Member.(string), 0),
			CreateTime: int64(z.Score),
		}
	}
	return list, nil
}

func getfollowerlist(ctx context.Context, uid, lastid int64) {

}
