package main

import (
	"context"
	"log"
	"testing"

	"github.com/zxq97/gokit/pkg/etcd"
	"github.com/zxq97/gokit/pkg/rpc"
	"github.com/zxq97/relation/api/relation/bff/v1"
)

func getRelationClient() (v1.RelationBFFClient, error) {
	etcdcli, err := etcd.NewEtcd(&etcd.Config{Addr: []string{"192.168.0.121:2379", "192.168.0.122:2379", "192.168.0.123:2379"}, TTL: 3})
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	conn, err := rpc.NewGrpcConn(ctx, "relation_bff", etcdcli)
	if err != nil {
		return nil, err
	}
	client := v1.NewRelationBFFClient(conn)
	return client, nil
}

func TestFollowerList(t *testing.T) {
	client, err := getRelationClient()
	if err != nil {
		t.Fatal(err)
	}
	var lastid int64
	cnt := 0
	m := make(map[int64]struct{})
	for {
		res, err := client.GetFollowerList(context.TODO(), &v1.ListRequest{Uid: 0, LastId: lastid})
		if err != nil {
			t.Error(err)
		}
		log.Println(res.ItemList[0].Uid, res.ItemList[len(res.ItemList)-1].Uid, cnt)
		if len(res.ItemList) == 0 {
			break
		}
		for _, v := range res.ItemList {
			if _, ok := m[v.Uid]; !ok {
				m[v.Uid] = struct{}{}
			} else {
				t.Fatal("replicate", lastid, v.Uid)
			}
		}
		cnt++
		lastid = res.ItemList[len(res.ItemList)-1].Uid
	}
}
