package social

import (
	"context"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestJobFollowOne(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	newJobConsumer(ctx, 3, followch)
	//newTaskConsumer(ctx, 1, countch)
	wg := sync.WaitGroup{}
	now := time.Now()
	for i := 0; i < mux; i++ {
		wg.Add(1)
		u := g.Gen()
		go func() {
			defer wg.Done()
			serviceFollow(u, vuid)
		}()
	}
	wg.Wait()
	logger.Println(time.Since(now))
	close(followch)
	jobwg.Wait()
	logger.Println(time.Since(now))
	//close(countch)
	//taskwg.Wait()
	//logger.Println(time.Since(now))
	cancel()
}

func TestJobFollowMulti(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	newJobConsumer(ctx, 3, followch)
	newTaskConsumer(ctx, 1, countch)
	rand.Seed(time.Now().UnixMilli())
	wg := sync.WaitGroup{}
	now := time.Now()
	for i := 0; i < mux; i++ {
		wg.Add(1)
		u := g.Gen()
		o := rand.Int63n(100)
		go func() {
			defer wg.Done()
			serviceFollow(u, o)
		}()
	}
	wg.Wait()
	logger.Println(time.Since(now))
	close(followch)
	jobwg.Wait()
	logger.Println(time.Since(now))
	close(countch)
	taskwg.Wait()
	logger.Println(time.Since(now))
	cancel()
}
