package main

import (
	"bench/cmd/social"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/zxq97/gokit/pkg/mq/kafka"
)

func main() {
	c, ch, err := kafka.NewConsumer(&kafka.Config{[]string{"10.203.0.27:9092"}}, []string{social.Topic}, "job", "jobc1name", social.ConsumerFollow, 10, 10, time.Second)
	if err != nil {
		panic(err)
	}
	c.Start()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT)

	select {
	case <-sig:
		if err = c.Stop(); err != nil {
			social.Logger.Println(err)
		}
		<-ch
	}
}
