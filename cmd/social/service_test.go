package social

import (
	"context"
	"sync"
	"testing"
)

func TestFollow(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(mux)
	for i := 0; i < mux; i++ {
		u := g.Gen()
		//o := g.Gen()
		go func() {
			defer wg.Done()
			if err := handleFollow(context.TODO(), u, vuid); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkFollow(b *testing.B) {

}
