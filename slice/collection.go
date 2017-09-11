package slice

import (
	"reflect"
)

// Collection 通用集合类型
type Collection []interface{}

// Predicate 通用过滤方法
type Predicate func(it, inx interface{}) bool

// From 转换 slice 为 Collection
func From(s interface{}) (rtn Collection) {
	if !IsSlice(s) {
		return
	}

	v := reflect.Indirect(reflect.ValueOf(s))
	len := v.Len()

	for i := 0; i < len; i++ {
		rtn = append(rtn, v.Index(i).Interface())
	}

	return
}

// Filter 过滤符合条件的集合
func (c Collection) Filter(f Predicate) (rtn Collection) {
	for inx, it := range c {
		if f(it, inx) {
			rtn = append(rtn, it)
		}
	}
	return
}

// Find 查找集合中的元素
func (c Collection) Find(f Predicate) (res interface{}) {
	for inx, it := range c {
		if f(it, inx) {
			res = it
		}
	}
	return
}
