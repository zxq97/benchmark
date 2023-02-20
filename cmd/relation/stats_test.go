package main

import (
	"context"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/zxq97/gokit/pkg/generate"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

var (
	wg     sync.WaitGroup
	swg    sync.WaitGroup
	logger *log.Logger
	lock   sync.Mutex
	db     *gorm.DB
)

const (
	cnt = 20000
)

func init() {
	var err error
	db, err = gorm.Open(mysql.Open("root:GUOan1992@tcp(127.0.0.1:3306)/zzlove_relation?charset=utf8mb4&parseTime=True"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "uid",
		NumberOfShards:      4,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "user_follows", "user_followers", "user_follow_counts")); err != nil {
		panic(err)
	}
	logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
}

func addFollowCount(uid int64, now time.Time) error {
	defer swg.Done()
	return db.Transaction(func(tx *gorm.DB) error {
		sql := "SELECT `id` FROM `user_followers` WHERE `uid` = ? AND `create_at` >= ? AND `stats` = 0"
		var ids []int64
		if err := tx.Raw(sql, uid, now).Scan(&ids).Error; err != nil {
			return err
		}
		if len(ids) == 0 {
			return nil
		}
		sql = "UPDATE `user_follow_counts` SET `follower_count` = `follower_count` + ? WHERE `uid` = ? LIMIT 1"
		res := tx.Exec(sql)
		if err := res.Error; err != nil {
			return err
		} else if res.RowsAffected == 0 {
			sql = "INSERT INTO `user_follow_counts` (`uid`, `follower_count`) VALUES (?, ?)"
			if err = tx.Exec(sql, uid, len(ids)).Error; err != nil {
				return err
			}
		}
		sql = "UPDATE `user_followers` SET `stats` = 1 WHERE `uid` = ? AND `id` IN (?)"
		return tx.Exec(sql, uid, ids).Error
	})
}

func follow(ctx context.Context, uid, touid int64) error {
	now := time.Now()
	if err := db.Transaction(func(tx *gorm.DB) error {
		tx = tx.WithContext(ctx)
		sql := "INSERT INTO `user_follows` (`uid`, `to_uid`) VALUES (?, ?)"
		if err := tx.Exec(sql, uid, touid).Error; err != nil {
			return err
		}
		sql = "INSERT INTO `user_followers` (`uid`, `to_uid`) VALUES (?, ?)"
		if err := tx.Exec(sql, touid, uid).Error; err != nil {
			return err
		}
		sql = "UPDATE `user_follow_counts` SET `follow_count` = `follow_count` + 1 WHERE `uid` = ? LIMIT 1"
		res := tx.Exec(sql, uid)
		if err := res.Error; err != nil {
			return err
		} else if res.RowsAffected == 0 {
			sql = "INSERT INTO `user_follow_counts` (`uid`, `follow_count`) VALUES (?, 1)"
			return tx.Exec(sql, uid).Error
		}
		return nil
	}); err != nil {
		return err
	}
	if lock.TryLock() {
		swg.Add(1)
		time.AfterFunc(time.Second, func() {
			if err := addFollowCount(touid, now); err != nil {
				logger.Println(err)
				lock.Unlock()
			}
		})
	}
	return nil
}

func consumer(ch <-chan int64) {
	for {
		select {
		case uid, ok := <-ch:
			if ok {
				ctx, cancel := context.WithTimeout(context.TODO(), time.Second*3)
				if err := follow(ctx, uid, 1); err != nil {
					logger.Println(uid, err)
				}
				cancel()
			} else {
				wg.Done()
				return
			}
		}
	}
}

func TestFollow(t *testing.T) {
	ch := make(chan int64, cnt)
	g, err := generate.NewSnowIDGen("relation")
	if err != nil {
		panic(err)
	}
	now := time.Now()
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- g.Gen()
		}
		close(ch)
	}()
	wg.Add(10)
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	go func() {
		consumer(ch)
	}()
	wg.Wait()
	logger.Println(time.Since(now))
	cnow := time.Now()
	swg.Wait()
	logger.Println(time.Since(cnow))
}
