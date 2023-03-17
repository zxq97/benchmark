package social

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/zxq97/gokit/pkg/mq/kafka"
)

func ConsumerSync(ctx context.Context, msg *kafka.KafkaMessage) error {
	switch msg.EventType {
	case 0:
	case 1:
		if n.CompareAndSwap(0, 1) {
			now = time.Now()
		}
		if cur := atomic.AddInt64(&cnt, 1); cur%1000 == 0 {
			Logger.Println(cur, time.Since(now))
		}
		argc := &Message{}
		if err := proto.Unmarshal(msg.Message, argc); err != nil {
			return err
		}
		return syncCount(ctx, argc.Uid)
	}
	return nil
}
