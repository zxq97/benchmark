package dal

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
	"github.com/zxq97/gokit/pkg/cache/xredis"
	"github.com/zxq97/gokit/pkg/mq/kafka"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	SocialDB *gorm.DB
	XRedis   *xredis.XRedis
	MC       *memcache.Client
	LC       *cache.Cache
	KAFKA    *kafka.Producer
)

func init() {
	var err error
	SocialDB, err = gorm.Open(mysql.Open("poketest:n2Xo4Vp25R9wy4T7@tcp(10.203.1.133:3306)/pokekara_event?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	redcli := redis.NewClient(&redis.Options{Addr: "10.203.1.1:6379"})
	XRedis = &xredis.XRedis{Cmdable: redcli}
	MC = memcache.New("10.203.0.27:11211")
	LC = cache.New(time.Second, time.Minute*5)
	KAFKA, err = kafka.NewProducer(&kafka.Config{Addr: []string{"10.203.0.27:9092"}})
	if err != nil {
		panic(err)
	}
}
