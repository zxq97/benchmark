package social

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestServiceFollowOne(t *testing.T) {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.TODO())
	taskwg.Add(1)
	go taskConsumer(ctx, countch)
	now := time.Now()
	for i := 0; i < 1000; i++ {
		u := g.Gen()
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()
			if err := followOne(ctx, u, vuid); err != nil {
				logger.Println(u, err)
			}
		}()
	}
	wg.Wait()
	logger.Println(time.Since(now))
	cancel()
	taskwg.Wait()
}

func TestServiceFollow(t *testing.T) {
	wg := sync.WaitGroup{}
	now := time.Now()
	for i := 0; i < 1000; i++ {
		u := g.Gen()
		o := rand.Int63n(100)
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, cancel := context.WithTimeout(context.TODO(), time.Second)
			defer cancel()
			if err := follow(ctx, u, o); err != nil {
				logger.Println(u, o, err)
			}
		}()
	}
	wg.Wait()
	logger.Println(time.Since(now))
}
