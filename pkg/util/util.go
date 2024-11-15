package util

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Ptr[T any](val T) *T {
	return &val
}

func Int(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		panic("int conversion failed: " + err.Error())
	}
	return i
}

func Int64(val string) int64 {
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic("int64 conversion failed: " + err.Error())
	}
	return i
}

func Float(val string) float64 {
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic("float conversion failed: " + err.Error())
	}
	return f
}

func Str(val any) string {
	return fmt.Sprintf("%v", val)
}

func ByteStr(b ...byte) string {
	return string(b)
}

func ReadFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		panic("file read failed: " + err.Error())
	}
	return string(content)
}

func ReadLines(fileName string) []string {
	return strings.Split(ReadFile(fileName), "\n")
}

func ForLines(fileName string, processor func(line string)) {
	for _, line := range ReadLines(fileName) {
		processor(line)
	}
}

func ForBytes(str string, processor func(c byte)) {
	bytes := []byte(str)
	for i := range bytes {
		processor(bytes[i])
	}
}

func ForLinesBytes(fileName string, processor func(line []byte)) {
	for _, line := range ReadLines(fileName) {
		processor([]byte(line))
	}
}

func Map[F, T any](in []F, mapper func(item F) T) []T {
	result := make([]T, len(in))
	for i := range in {
		result[i] = mapper(in[i])
	}
	return result
}

func Filter[T any](in []T, filter func(item T) bool) []T {
	return slices.DeleteFunc(in, func(v T) bool {
		return !filter(v)
	})
}

func IsNumeric(c byte) bool {
	return c >= '0' && c <= '9'
}

func Assert(condition bool) {
	if !condition {
		panic("assertion failure")
	}
}
