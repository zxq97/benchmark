package social

import (
	"context"
	"sync"
	"testing"
)

func TestFollow(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(Mux)
	for i := 0; i < Mux; i++ {
		u := G.Gen()
		//o := g.Gen()
		go func() {
			defer wg.Done()
			if err := HandleFollow(context.TODO(), u, Vuid); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}

func BenchmarkFollow(b *testing.B) {

}
