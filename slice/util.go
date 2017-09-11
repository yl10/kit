package slice

import (
	"reflect"
)

// IsSlice 是否 slice
func IsSlice(target interface{}) bool {
	return reflect.ValueOf(target).Kind() == reflect.Slice
}
