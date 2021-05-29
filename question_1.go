package stream_test

import (
	"github.com/Pallinder/go-randomdata"
	"github.com/yuanyunxu/stream_test/stream"
	"github.com/yuanyunxu/stream_test/stream/types"
)

// - Q1: 输入 employees，返回 年龄 >22岁 的所有员工，年龄总和
func Question1Sub1(employees []*Employee) int64 {
	return int64(stream.OfSlice(employees).Filter(func(e types.T) bool {
		return *e.(*Employee).Age > 22
	}).ReduceWith(0, func(acc types.R, e types.T) types.R {
		return acc.(int) + *e.(*Employee).Age
	}).(int))
}

// - Q2: - 输入 employees，返回 id 最小的十个员工，按 id 升序排序
func Question1Sub2(employees []*Employee) []*Employee {
	return stream.OfSlice(employees).Sorted(func(left types.T, right types.T) int {
		return types.Int64Comparator(left.(*Employee).Id, right.(*Employee).Id)
	}).Limit(10).ToElementSlice(&Employee{}).([]*Employee)
}

// - Q3: - 输入 employees，对于没有手机号为0的数据，随机填写一个
func Question1Sub3(employees []*Employee) []*Employee {
	stream.OfSlice(employees).ForEach(func(e types.T) {
		if e.(*Employee).Phone == nil {
			temPhoneNums := randomdata.PhoneNumber()
			e.(*Employee).Phone = &temPhoneNums
		}
	})
	return employees
}

// - Q4: - 输入 employees ，返回一个map[int][]int，其中 key 为 员工年龄 Age，value 为该年龄段员工ID
func Question1Sub4(employees []*Employee) map[int][]int64 {
	stream.OfSlice(employees).ReduceWith(make(map[int][]int64, 0), func(acc types.R, e types.T) types.R {
		acc.(map[int][]int64)[*e.(*Employee).Age] = append(acc.(map[int][]int64)[*e.(*Employee).Age], e.(*Employee).Id)
		return acc
	})
	return nil
}
