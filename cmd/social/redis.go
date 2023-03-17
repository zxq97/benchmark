package social

import (
	"context"
	"time"

	"github.com/zxq97/gokit/pkg/cast"
)

func sendSyncLock(ctx context.Context, uid int64) (bool, error) {
	return Xr.SetNX(ctx, cast.FormatInt(uid), "", time.Second).Result()
}
