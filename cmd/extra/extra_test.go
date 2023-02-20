package main

import (
	"bench/dal"
	"bench/dal/query"
	"context"
	"log"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/zxq97/gokit/pkg/generate"
)

var (
	q      *query.Query
	wg     sync.WaitGroup
	swg    sync.WaitGroup
	logger *log.Logger
	lock   sync.Mutex
)

const (
	cnt = 100000
)

func addFollowerCount(uid int64) error {
	ctx := context.TODO()
	defer swg.Done()
	return q.Transaction(func(tx *query.Query) error {
		rows, err := tx.WithContext(ctx).ExtraFollower.FindUnSyncRecordByUID(uid)
		if err != nil {
			return err
		}
		var n int64
		ids := make([]int64, len(rows))
		for k, v := range rows {
			ids[k] = v.ID
			if v.Stats == 0 {
				n++
			} else {
				n--
			}
		}
		if err = tx.WithContext(ctx).FollowCount.IncrByFollowerCount(uid, n); err != nil {
			return err
		}
		return tx.WithContext(ctx).ExtraFollower.DeleteRecord(ids)
	})
}

func follow(ctx context.Context, uid, touid int64) error {
	if err := q.Transaction(func(tx *query.Query) error {
		if err := tx.WithContext(ctx).Follow.InsertFollow(uid, touid); err != nil {
			return err
		}
		if err := tx.WithContext(ctx).Follower.InsertFollower(touid, uid); err != nil {
			return err
		}
		if err := tx.WithContext(ctx).FollowCount.IncrFollowCount(uid); err != nil {
			return err
		}
		return tx.WithContext(ctx).ExtraFollower.InsertFollower(touid)
	}); err != nil {
		return err
	}
	if lock.TryLock() {
		swg.Add(1)
		time.AfterFunc(time.Second, func() {
			if err := addFollowerCount(touid); err != nil {
				logger.Println(err)
			}
			lock.Unlock()
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
	logger = log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile)
	q = query.Use(dal.SocialDB)
	ch := make(chan int64, cnt)
	g, err := generate.NewSnowIDGen("relation")
	if err != nil {
		t.Fatal(err)
	}
	now := time.Now()
	go func() {
		for i := 0; i < cnt; i++ {
			ch <- g.Gen()
		}
		close(ch)
	}()
	wg.Add(20)
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
