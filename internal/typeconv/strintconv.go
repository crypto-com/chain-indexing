package typeconv

import (
	"strconv"
)

func MustAtoi32(v string) int32 {
	result, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(result)
}

func MustAtou32(v string) uint32 {
	result, err := strconv.ParseUint(v, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(result)
}

func MustAtou64(v string) uint64 {
	result, err := strconv.ParseUint(v, 10, 64)
	if err != nil {
		panic(err)
	}
	return result
}
