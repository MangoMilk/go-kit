// Copyright © 2024 VincentY. All rights reserved.
// Distribution Lock
// 可重入
// 阻塞
// 公平
// 高可用获取和释放
// 高性能获取和释放

package dlock

import (
	"time"
)

type DLockInterface interface {
	Lock(key string, duration time.Duration) (bool, error)
	Release(key string) (bool, error)
}
