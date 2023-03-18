package main

import (
	"bench/cmd/social"
	"context"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(social.Mux)
	now := time.Now()
	for i := 0; i < social.Mux; i++ {
		u := social.G.Gen()
		go func() {
			defer wg.Done()
			if err := social.HandleFollow(context.TODO(), u, social.Vuid); err != nil {
				social.Logger.Println(err)
			}
		}()
	}
	wg.Add(social.Mux * 5)
	for i := 0; i < social.Mux*5; i++ {
		u := social.G.Gen()
		t := social.G.Gen()
		go func() {
			defer wg.Done()
			if err := social.HandleFollow(context.TODO(), u, t); err != nil {
				social.Logger.Println(err)
			}
		}()
	}
	wg.Wait()
	social.Logger.Println(time.Since(now))
}
