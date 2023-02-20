package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/zxq97/gokit/pkg/etcd"
	"github.com/zxq97/gokit/pkg/rpc"
	"github.com/zxq97/relation/api/relation/bff/v1"
)

func main() {
	etcdcli, err := etcd.NewEtcd(&etcd.Config{Addr: []string{"192.168.0.121:2379", "192.168.0.122:2379", "192.168.0.123:2379"}, TTL: 3})
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	conn, err := rpc.NewGrpcConn(ctx, "relation_bff", etcdcli)
	if err != nil {
		panic(err)
	}
	client := v1.NewRelationBFFClient(conn)
	wg := sync.WaitGroup{}
	now := time.Now()
	//_, err = client.Follow(context.TODO(), &v1.FollowRequest{Uid: 1, ToUid: 0})
	for i := 0; i < 5000; i++ {
		u := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err = client.Follow(context.TODO(), &v1.FollowRequest{
				Uid:   int64(u),
				ToUid: 0,
			}); err != nil {
				fmt.Println(u, err)
			}
		}()
	}
	wg.Wait()
	fmt.Println(time.Since(now))
}
