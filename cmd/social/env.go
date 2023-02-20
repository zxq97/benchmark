package social

import (
	"bench/dal"
	"bench/dal/query"
	"log"
	"os"
	"sync"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"github.com/zxq97/gokit/pkg/cache/xredis"
	"github.com/zxq97/gokit/pkg/generate"
)

const (
	mux  = 100000
	vuid = 1

	followField   = "follow"
	followerField = "follower"

	lockkey = "lock_%d"

	redisKeyHRelationCount = "lra_cnt_%d" // uid
)

var (
	c         *cron.Cron
	g         *generate.SnowIDGen
	q         *query.Query
	xr        *xredis.XRedis
	mc        *memcache.Client
	lc        *cache.Cache
	logger    *log.Logger
	followch  chan *asyncFollow
	rebuildch chan *rebuildCache
	countch   chan *syncCount

	jobwg  sync.WaitGroup
	taskwg sync.WaitGroup
)

type asyncFollow struct {
	uid   int64
	touid int64
}

type rebuildCache struct {
	uid    int64
	lastid int64
}

type syncCount struct {
	uid      int64
	duration int64
}

func init() {
	logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	c = cron.New()
	g, _ = generate.NewSnowIDGen("social")
	q = query.Use(dal.SocialDB)
	xr = dal.XRedis
	mc = dal.MC
	lc = dal.LC
	followch = make(chan *asyncFollow, mux)
	rebuildch = make(chan *rebuildCache, mux)
	countch = make(chan *syncCount, mux)
}
