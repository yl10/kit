package cube

import (
	"errors"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// InputToStruct url.Values 映射为 struct
// 将 url.Values 中的数据, 转入到 struct 中
// 一般使用场景是, 将 request 传入的 params 转到方便操作的 struct 中
// 一个缺点是, 数组会转为 逗号连接的字串
func InputToStruct(input url.Values, s interface{}) error {
	modType := reflect.TypeOf(s)
	modVal := reflect.ValueOf(s)

	if modType.Kind() == reflect.Ptr {
		modVal = modVal.Elem()
		modType = modVal.Type()
	}

	for i := 0; i < modType.NumField(); i++ {
		f := modType.Field(i)
		nm := f.Name
		// nm := Hump2Hyphen(f.Name)
		typ := f.Type

		vals, isNotNull := input[nm]
		if !isNotNull {
			vals = []string{""}
		}

		val := vals[0]
		if len(vals) > 1 {
			val = strings.Join(vals, ",")
		}

		switch typ.String() {
		case "string":
			modVal.Field(i).SetString(val)
		case "int":
			if !isNotNull {
				modVal.Field(i).SetInt(0)
			} else {
				if v, e := strconv.ParseInt(val, 10, 64); e != nil {
					return e
				} else {
					modVal.Field(i).SetInt(v)
				}
			}
		case "time.Time":
			if !isNotNull {
				modVal.Field(i).Set(reflect.ValueOf(time.Now()))
			} else {
				tz, _ := time.LoadLocation("Asia/Shanghai")
				if v, e := time.ParseInLocation("2006-01-02 15:04:05", val, tz); e != nil {
					return e
				} else {
					modVal.Field(i).Set(reflect.ValueOf(v))
				}
			}
		default:
			return errors.New("Map映射Model失败, Model 中含有不支持的类型 (" + typ.String() + ")")
		}
	}

	return nil
}
