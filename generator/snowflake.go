package generator

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

// 优点
// 1.实现简单，速度快
// 2.各位段随意调整，灵活

// 缺点
// 1.依赖系统时间
// 2.趋势递增，非连续绝对递增（也有人认为是优点，连续绝对递增容易被猜到真实新增量）

// 时间回拨问题处理
// 1.简单暴力，等待或报错给业务层处理
// 2.回拨未知 2 bit
// 3.使用阿里云的时间服务器(ntp)进行同步

// 1 + 41 + 10 + 12，正负位、1毫秒时间戳位、机器位、序列号位
// 1 + 41 + 8 + 2 + 12，正负位、1毫秒时间戳位、机器位、回拨位、序列号位

const (
	millisecondTimestampBit = 41
	machineBit              = 10
	sequenceBit             = 12

	maxSequence = (1 << sequenceBit) - 1 // 4096 -1
	minNode     = 1
	maxNode     = 1 << machineBit // 1024
)

var (
	Snowflake *snowflakeGenerator

	ErrNodeOverflow = errors.New("machine bit overflow")
	ErrNodeNotSetup = errors.New("node has not been setup")
)

type snowflakeGenerator struct {
	lock              sync.Mutex
	snowflakeSequence int64
	node              int64
}

// @param: machine string [1, 1024]
// @return: error
func newSnowflakeIDGenerator(machine string) error {
	n, err := strconv.ParseInt(machine, 10, 64)
	if err != nil {
		return err
	}

	if n < minNode || n > maxNode {
		return ErrNodeOverflow
	}

	Snowflake = &snowflakeGenerator{
		snowflakeSequence: 0,
		node:              n << sequenceBit,
	}

	return nil
}

func (g *snowflakeGenerator) GenID() (int64, error) {
	if g.node <= 0 {
		return 0, ErrNodeNotSetup
	}

	millisecondTimestamp := (time.Now().UnixNano() / 1e6) << (machineBit + sequenceBit)

	g.lock.Lock()
	defer g.lock.Unlock()
	g.snowflakeSequence = (g.snowflakeSequence + 1) & maxSequence

	return millisecondTimestamp + g.node + g.snowflakeSequence, nil
}
