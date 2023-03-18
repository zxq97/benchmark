package social

import (
	"bench/dal"
	"bench/dal/query"
	"log"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/patrickmn/go-cache"
	"github.com/zxq97/gokit/pkg/cache/xredis"
	"github.com/zxq97/gokit/pkg/generate"
	"github.com/zxq97/gokit/pkg/mq/kafka"
)

const (
	Mux  = 100000
	Vuid = 1

	FollowField   = "follow"
	FollowerField = "follower"

	Topic     = "social"
	Topicsync = "social_sync"
)

var (
	G      *generate.SnowIDGen
	Q      *query.Query
	Xr     *xredis.XRedis
	Mc     *memcache.Client
	Lc     *cache.Cache
	Logger *log.Logger
	P      *kafka.Producer
)

func init() {
	Logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	G, _ = generate.NewSnowIDGen("social")
	Q = query.Use(dal.SocialDB)
	Xr = dal.XRedis
	Mc = dal.MC
	Lc = dal.LC
	P = dal.KAFKA
	n.Store(0)
}
