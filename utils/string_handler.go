package utils

import (
	"fmt"
	"strconv"
)

const (
	floatPrec    = -1
	floatBitSize = 64
)

func ToString(val interface{}) string {
	switch val := val.(type) {
	case string:
		return val
	case int8:
		return strconv.FormatInt(int64(val), 10)
	case int16:
		return strconv.FormatInt(int64(val), 10)
	case int32:
		return strconv.FormatInt(int64(val), 10)
	case int64:
		return strconv.FormatInt(val, 10)
	case float32:
		return strconv.FormatFloat(float64(val), 'g', floatPrec, floatBitSize)
	case float64:
		return strconv.FormatFloat(val, 'g', floatPrec, floatBitSize)
	case uint8:
		return strconv.FormatUint(uint64(val), 10)
	case uint16:
		return strconv.FormatUint(uint64(val), 10)
	case uint32:
		return strconv.FormatUint(uint64(val), 10)
	case uint64:
		return strconv.FormatUint(uint64(val), 10)
	case bool:
		return strconv.FormatBool(val)
	default:
		return fmt.Sprintf("%+v", val)
	}
}
