package slice

import (
	"reflect"
)

// Find 查找符合函数 f 定义的 一个元素
func Find(s interface{}, f func(interface{}) bool) interface{} {
	switch reflect.TypeOf(s).Kind() {
	case reflect.Slice:
		values := reflect.Indirect(reflect.ValueOf(s))
		for i := 0; i < values.Len(); i++ {
			it := values.Index(i).Interface()
			if f(it) {
				return it
			}
		}
	}
	return nil
}
