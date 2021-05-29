package stream_test

import (
	"github.com/yuanyunxu/stream_test/stream"
	"github.com/yuanyunxu/stream_test/stream/types"
)

// - Q1: 输入一个整数 int，字符串string。将这个字符串重复n遍返回
func Question3Sub1(str string, n int) string {
	return stream.RepeatN(str, int64(n)).ReduceWith("", func(acc types.R, e types.T) types.R {
		return acc.(string) + e.(string)
	}).(string)
}
