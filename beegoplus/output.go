package beegoplus

import (
	"fmt"
	"reflect"
)

//OutParam 输出参数
type OutParam struct {
	Name        string
	Type        ParamType
	Muilti      bool //是否数组
	Description string
}

//OutPut 输出
type OutPut struct {
	code    int
	message string
	data    map[string][]interface{}
}

//NewOutPut 实例化一个OutPut
func newOutPut(code int, result interface{}) *OutPut {
	op := OutPut{code: code}
	if code == 0 {
		op.Set("data", result)
		return &op
	}
	op.Set("message", fmt.Sprintf("%v", result))
	return &op
}

//Get get
func (o OutPut) Get(key string) interface{} {
	if o.data == nil {
		return nil
	}

	return o.data[key]
}

// //Gets 获取
// func (o OutPut) Gets(key string) []interface{} {
// 	if o == nil {
// 		return nil
// 	}
// 	vs := o[key]
// 	if len(vs) == 0 {
// 		return nil
// 	}
// 	return vs
// }

// Set sets the key to value. It replaces any existing
// values.
func (o OutPut) Set(key string, value interface{}) {
	o.data[key] = []interface{}{value}
}

// Add adds
func (o OutPut) Add(key string, value ...interface{}) {
	o.data[key] = append(o.data[key], value...)
}

// Del deletes the values associated with key.
func (o OutPut) Del(key string) {
	delete(o.data, key)
}

//AddStructObject 增加一个struct
//仅支持struct结构，其他结构不做任何处理
func (o OutPut) AddStructObject(structobject interface{}) {

	v := reflect.Indirect(reflect.ValueOf(structobject))
	t := v.Type()
	if t.Kind() == reflect.Struct {
		for i := 0; i < t.NumField(); i++ {
			o.Add(t.Field(i).Name, v.Field(i).Interface())
		}
	}
}

// //SetAuth 对输出设置权限控制
// func (o OutPut) SetAuth(aumap map[string]bool) {

// }

// //MarshalJSON 实现json接口
// func (o OutPut) MarshalJSON() ([]byte, error) {

// 	return json.Marshal(o.MarshalMap())

// }

// //MarshalMap 转为map
// func (o OutPut) MarshalMap() map[string]interface{} {
// 	r := make(map[string]interface{})
// 	r["code"] = o.code

// 	if o.code != 0 {
// 		r["message"] = o.message
// 		return r
// 	}
// 	for k, v := range o.data {
// 		isout := o.auth[k]
// 		if isout {
// 			//根据路由里的设置，判断是否输出数组
// 			if len(v) <= 1 {
// 				//对象赋值，暂未实现
// 			} else {
// 				r[k] = v
// 			}

// 		}
// 	}
// 	return r
// }
