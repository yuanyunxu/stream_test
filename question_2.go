package stream_test

import (
	"strconv"
	"unicode"

	"github.com/yuanyunxu/stream_test/stream"
	"github.com/yuanyunxu/stream_test/stream/types"
)

// - Q1: 计算一个 string 中小写字母的个数
func Question2Sub1(str string) int64 {
	return stream.OfStrings(str).ReduceWith(int64(0), func(acc types.R, e types.T) types.R {
		intE, _ := strconv.Atoi(e.(string))
		if unicode.IsLower(rune(intE)) {
			acc = (types.R)(acc.(int64) + 1)
		}
		return acc
	}).(int64)
}

// - Q2: 找出 []string 中，包含小写字母最多的字符串
func Question2Sub2(list []string) string {
	type initType struct {
		res   string
		count int64
	}
	return stream.OfSlice(list).ReduceWith(&initType{}, func(acc types.R, e types.T) types.R {
		count := acc.(*initType).count
		newCount := Question2Sub1(e.(string))
		if newCount > count {
			acc.(*initType).count = newCount
			acc.(*initType).res = e.(string)
		}
		return acc
	}).(*initType).res
}
