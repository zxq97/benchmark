package social

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/cast"
	"github.com/zxq97/gokit/pkg/mq/kafka"
)

var (
	now time.Time
	n   atomic.Value
	cnt int64
)

func ConsumerFollow(ctx context.Context, msg *kafka.KafkaMessage) error {
	switch msg.EventType {
	case 0:
		if n.CompareAndSwap(0, 1) {
			now = time.Now()
		}
		if cur := atomic.AddInt64(&cnt, 1); cur%10000 == 0 {
			Logger.Println(cur, time.Since(now))
		}
		argc := &Message{}
		if err := proto.Unmarshal(msg.Message, argc); err != nil {
			return err
		}
		if err := followOne(ctx, argc.Uid, argc.TargetId); err != nil {
			return err
		}
		if ok, _ := sendSyncLock(ctx, argc.TargetId); ok {
			if err := P.SendMessage(ctx, Topicsync, cast.FormatInt(argc.TargetId), &Message{Uid: argc.TargetId, TargetId: 0}, 1); err != nil {
				Logger.Println(err)
			}
		}
		//return follow(ctx, argc.Uid, argc.TargetId)
	case 1:

	}
	return nil
}
