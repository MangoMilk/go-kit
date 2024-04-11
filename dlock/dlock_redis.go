package dlock

import (
	"github.com/go-redis/redis"
	"time"
)

type DLockRedis struct {
	instance *redis.Client
}

func (d *DLockRedis) Lock(key string, duration time.Duration) (bool, error) {
	var (
		newExpiredTime int64 = time.Now().Add(duration).Unix()
		oldExpiredTime int64
		curExpiredTime int64
		err            error
		res            bool
	)

	// try to get lock
	res, err = d.instance.SetNX(key, newExpiredTime, duration).Result()
	if err != nil {
		return false, err
	}

	if res {
		// lock successfully
		return true, nil
	}

	oldExpiredTime, err = d.instance.Get(key).Int64()
	if err != nil {
		return false, err
	}

	if oldExpiredTime > time.Now().Unix() {
		// locking
		return false, nil
	}

	// locked overtime (maybe deadlock)
	curExpiredTime, err = d.instance.GetSet(key, newExpiredTime).Int64()
	if err != nil {
		return false, err
	}

	// try to get lock
	if curExpiredTime == oldExpiredTime {
		if err = d.instance.Expire(key, duration).Err(); err != nil {
			return false, err
		}

		return true, nil
	}

	return false, nil
}

func (d *DLockRedis) Release(key string) (bool, error) {
	oldExpiredTime, err := d.instance.Get(key).Int64()
	if err != nil {
		return false, err
	}

	if oldExpiredTime > time.Now().Unix() {
		res, err := d.instance.Del(key).Result()
		return res > 0, err
	}

	return false, nil
}

func NewRedisDLock(instance *redis.Client) DLockInterface {
	return &DLockRedis{
		instance: instance,
	}
}
