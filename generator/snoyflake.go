package generator

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

const (
	ipUnitCount   = 4
	ipUnitMin     = 0
	ipUnitMax     = 255
	ipUnitBit     = 8
	sequenceStart = 0
)

var (
	Snoyflake *snoyflakeGenerator

	ErrInvalidIPFormat = errors.New("ip format error")
)

type snoyflakeGenerator struct {
	snoyflakeSequence int64
	ip                []int64
}

// @param: machine string [0.0.0.0, 255.255.255.255]
// @return: error
func newSnoyflakeIDGenerator(machine string) error {
	ip := strings.Split(machine, ".")
	if len(ip) != ipUnitCount {
		return ErrInvalidIPFormat
	}

	ipUnits := make([]int64, ipUnitCount)
	for i := 0; i < ipUnitCount; i++ {
		ipUnit, err := strconv.ParseInt(ip[i], 10, 64)
		if err != nil {
			return err
		}

		if ipUnit > ipUnitMax || ipUnit < ipUnitMin {
			return ErrInvalidIPFormat
		}

		ipUnits[i] = ipUnit
	}

	Snoyflake = &snoyflakeGenerator{
		snoyflakeSequence: sequenceStart,
		ip:                ipUnits,
	}

	return nil
}

// 1 + 39 + 8 + 16，正负位、10毫秒时间戳位、序列号位、机器位
// 机器位：当前机器的私有IP的最后两位
func (g *snoyflakeGenerator) GenID() (int64, error) {

	if len(g.ip) != ipUnitCount {
		return 0, ErrInvalidIPFormat
	}

	tenMillisecondTimestamp := time.Now().UnixNano() / 1e7
	tenMillisecondTimestamp = tenMillisecondTimestamp << 24

	g.snoyflakeSequence = (g.snoyflakeSequence + 1) & (256 - 1)

	machine := (g.ip[2] << ipUnitBit) + g.ip[3]

	return tenMillisecondTimestamp + (g.snoyflakeSequence << (2 * ipUnitBit)) + machine, nil
}
