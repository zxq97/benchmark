package main

import (
	"bench/cmd/social"
	"context"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(social.Mux)
	for i := 0; i < social.Mux; i++ {
		u := social.G.Gen()
		//o := g.Gen()
		go func() {
			defer wg.Done()
			if err := social.HandleFollow(context.TODO(), u, social.Vuid); err != nil {
				social.Logger.Println(err)
			}
		}()
	}
	wg.Wait()
}
