package dal

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"github.com/zxq97/gokit/pkg/cache/xredis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	SocialDB *gorm.DB
	XRedis   *xredis.XRedis
	MC       *memcache.Client
	LC       *cache.Cache
)

func init() {
	var err error
	SocialDB, err = gorm.Open(mysql.Open("root:GUOan1992@tcp(127.0.0.1:3306)/zzlove_social?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	redcli := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	XRedis = &xredis.XRedis{Cmdable: redcli}
	MC = memcache.New("127.0.0.1:11211")
	LC = cache.New(time.Second, time.Minute*5)
}
