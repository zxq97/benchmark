package social

import (
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/zxq97/gokit/pkg/mq/kafka"
)

func TestJobFollowC1(t *testing.T) {
	c, ch, err := kafka.NewConsumer(&kafka.Config{[]string{"127.0.0.1:9092"}}, []string{Topic}, "job", "jobc1name", ConsumerFollow, 10, 10, time.Second)
	if err != nil {
		panic(err)
	}
	c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	select {
	case <-sig:
		if err = c.Stop(); err != nil {
			t.Error(err)
		}
		<-ch
	}
}
