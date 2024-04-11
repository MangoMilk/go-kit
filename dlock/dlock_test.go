package dlock

import (
	"github.com/go-redis/redis"
	"testing"
	"time"
)

func TestRedisDLock(t *testing.T) {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123465",
		DB:       0,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		t.Fatal(err)
	}

	lockKey := "Test"
	dLock := NewRedisDLock(rdb)
	res, err := dLock.Lock(lockKey, time.Second*3)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(res)
}
