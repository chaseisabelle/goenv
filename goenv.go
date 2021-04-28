package goenv

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func Exists(key string) bool {
	_, has := os.LookupEnv(key)

	return has
}

func String(key string, def string) string {
	str, has := os.LookupEnv(key)

	if !has {
		return def
	}

	return str
}

func Bool(key string, def bool) bool {
	tmp := "0"

	if def {
		tmp = "1"
	}

	return String(key, tmp) != "0"
}

func Int(key string, def int) (int, error) {
	return strconv.Atoi(String(key, fmt.Sprintf("%d", def)))
}

func Int32(key string, def int32) (int32, error) {
	tmp, err := Int(key, int(def))

	return int32(tmp), err
}

func Int64(key string, def int64) (int64, error) {
	tmp, err := Int(key, int(def))

	return int64(tmp), err
}

func Uint(key string, def uint) (uint, error) {
	tmp, err := Uint64(key, uint64(def))

	return uint(tmp), err
}

func Uint8(key string, def uint8) (uint8, error) {
	tmp, err := Uint64(key, uint64(def))

	return uint8(tmp), err
}

func Uint16(key string, def uint16) (uint16, error) {
	tmp, err := Uint64(key, uint64(def))

	return uint16(tmp), err
}

func Uint32(key string, def uint32) (uint32, error) {
	tmp, err := Uint64(key, uint64(def))

	return uint32(tmp), err
}

func Uint64(key string, def uint64) (uint64, error) {
	return strconv.ParseUint(String(key, fmt.Sprintf("%d", def)), 10, 64)
}

func Float32(key string, def float32) (float32, error) {
	tmp, err := Float64(key, float64(def))

	return float32(tmp), err
}

func Float64(key string, def float64) (float64, error) {
	return strconv.ParseFloat(String(key, fmt.Sprintf("%f", def)), 64)
}

func Bytes(key string, def []byte) []byte {
	return []byte(String(key, string(def)))
}

func Duration(key string, def int, as time.Duration) (time.Duration, error) {
	i, err := Int(key, def)

	return time.Duration(i) * as, err
}

func Millisecond(key string, def int) (time.Duration, error) {
	return Duration(key, def, time.Millisecond)
}

func Second(key string, def int) (time.Duration, error) {
	return Duration(key, def, time.Second)
}

func Minute(key string, def int) (time.Duration, error) {
	return Duration(key, def, time.Minute)
}

func Hour(key string, def int) (time.Duration, error) {
	return Duration(key, def, time.Hour)
}
