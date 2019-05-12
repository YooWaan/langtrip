package main

import (
	"fmt"
	"reflect"
	"sort"
)

type itr func() (int, bool)

func from(data []int) query {
	len := len(data)
	return query{
		itr: func() itr {
			index := 0
			return func() (int, bool) {
				ok := index < len
				val := 0
				if ok {
					val = data[index]
					index++
				}
				return val, ok
			}
		},
	}
}

type query struct {
	itr func() itr
}

func (q query) apply(v interface{}) {
	res := reflect.ValueOf(v)
	slice := reflect.Indirect(res)

	cap := slice.Cap()
	res.Elem().Set(slice.Slice(0, cap))

	next := q.itr()
	index := 0
	for val, ok := next(); ok; val, ok = next() {
		if index >= cap {
			slice, cap = extendCap(slice)
		}
		slice.Index(index).Set(reflect.ValueOf(val))
		index++
	}
	res.Elem().Set(slice.Slice(0, index))
}

func extendCap(s reflect.Value) (reflect.Value, int) {
	cap := s.Cap()
	if cap == 0 {
		cap = 1
	} else {
		cap *= 2
	}
	newSlice := reflect.MakeSlice(s.Type(), cap, cap)
	reflect.Copy(newSlice, s)
	return newSlice, cap
}

// map,select is reserved word
func (q query) selectBy(mapFn func(int) int) query {
	return query{
		itr: func() itr {
			next := q.itr()
			return func() (int, bool) {
				ret, ok := next()
				if ok {
					ret = mapFn(ret)
				}
				return ret, ok
			}
		},
	}
}

func (q query) where(filter func(i int) bool) query {
	return query{
		itr: func() itr {
			next := q.itr()
			return func() (int, bool) {
				var (
					val int
					ok  bool
				)
				for val, ok = next(); ok; val, ok = next() {
					if filter(val) {
						return val, ok
					}
				}
				return val, ok
			}
		},
	}
}

func (q query) orderBy(less func(i, j int) bool) query {
	var nums []int
	q.apply(&nums)
	sort.Sort(sorter{nums: nums, less: less})
	return from(nums)
}

type nums []int

func (ns nums) Len() int      { return len(ns) }
func (ns nums) Swap(i, j int) { ns[i], ns[j] = ns[j], ns[i] }

type sorter struct {
	nums
	less func(i, j int) bool
}

func (s sorter) Less(i, j int) bool {
	return s.less(s.nums[i], s.nums[j])
}

func main() {
	data := []int{1, 10, 43, 28, 32, 5}
	var values []int
	from(data).apply(&values)
	fmt.Println("data: ", values)

	var selectValues []int
	from(data).selectBy(func(num int) int {
		return num * 2
	}).where(func(num int) bool {
		return num >= 50
	}).orderBy(func(i, j int) bool {
		// DESC
		return i > j
	}).apply(&selectValues)
	fmt.Println("select data: ", selectValues)
}
