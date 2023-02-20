package relation_bench

import (
	"context"

	"github.com/zxq97/gokit/pkg/etcd"
	"github.com/zxq97/gokit/pkg/rpc"
	"github.com/zxq97/relation/api/relation/bff/v1"
)

var (
	client v1.RelationBFFClient
)

func init() {
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
	client = v1.NewRelationBFFClient(conn)
}
