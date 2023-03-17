package social

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/zxq97/gokit/pkg/mq/kafka"
)

func TestTaskC(t *testing.T) {
	co := cron.New(cron.WithSeconds())
	_, err := co.AddFunc("*/1 * * * *", func() {
		_ = CronSync(context.TODO())
	})
	c, ch, err := kafka.NewConsumer(&kafka.Config{[]string{"127.0.0.1:9092"}}, []string{Topicsync}, "task", "task1name", ConsumerSync, 1, 1, time.Second*3)
	if err != nil {
		panic(err)
	}
	c.Start()
	co.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	select {
	case <-sig:
		if err = c.Stop(); err != nil {
			t.Error(err)
		}
		co.Stop()
		<-ch
	}
}
