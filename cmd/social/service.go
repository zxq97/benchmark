package social

import (
	"context"

	"github.com/zxq97/gokit/pkg/cast"
)

func HandleFollow(ctx context.Context, uid, touid int64) error {
	//m, err := getFollowCount(ctx, []int64{uid})
	//if err != nil {
	//	return err
	//}
	//if c, ok := m[uid]; ok && c.FollowCount > 3000 {
	//	return nil
	//}
	return P.SendMessage(ctx, "social", cast.FormatInt(uid), &Message{Uid: uid, TargetId: touid}, 0)
}

func handleUnfollow(ctx context.Context, uid, touid int64) error {
	return nil
}
