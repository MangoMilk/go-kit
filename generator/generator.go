package generator

import (
	"errors"
)

type GeneratorType uint8

const (
	SnowflakeGenerator = GeneratorType(1)
	SnoyflakeGenerator = GeneratorType(2)
)

type IGenerator interface {
	GenID() (int64, error)
}

func SetupGenerator(generator GeneratorType, machine string) error {
	switch generator {
	case SnowflakeGenerator:
		return newSnowflakeIDGenerator(machine)
	case SnoyflakeGenerator:
		return newSnoyflakeIDGenerator(machine)
	}

	return errors.New("not support this generator")
}
