package structconv

import (
	"net/url"
	"reflect"
	"strconv"
	"time"
)

type structconv struct {
	tagName    string
	timeFormat string
	vurlValues url.Values
	mapValue   map[string]interface{}
}

func newDefaultconv() *structconv {
	return &structconv{
		tagName:    "yl",
		timeFormat: "2006-01-02 15:04:05",
		vurlValues: url.Values{}}
}

//ToURLValues 转换为url.Values
//如果传入参数不是struct类型，返回空的values
func ToURLValues(o interface{}) url.Values {
	val := newDefaultconv()
	val.toValues(o)
	return val.vurlValues
}

//ToURLValuesWithTimeFormart 指定时间格式转换
//输入:o struct
//输入:timeformat 日期格式
func ToURLValuesWithTimeFormart(o interface{}, timeformt string) url.Values {
	val := newDefaultconv()
	val.timeFormat = timeformt
	val.toValues(o)
	return val.vurlValues
}
func (conv *structconv) toValues(o interface{}) {

	v := reflect.ValueOf(o)
	t := reflect.TypeOf(o)
	if t.Kind() != reflect.Struct {
		panic("传入参数不是struct")
	}

	for i := 0; i < t.NumField(); i++ {

		innerv := reflect.ValueOf(v.Field(i).Interface())
		innerStructField := t.Field(i)

		name, _ := parseTag(innerStructField.Tag.Get(conv.tagName))

		if name == "" {
			name = innerStructField.Name
		}

		if name != "-" {
			switch v.Field(i).Kind() {
			case reflect.Struct:
				switch {
				//时间格式的导出
				case t.Field(i).Type == reflect.TypeOf(*new(time.Time)):
					tv, _ := innerv.Interface().(time.Time)
					conv.vurlValues.Set(name, tv.Format(conv.timeFormat))
				//闭包的时候导出，不是闭包的不导出
				case innerStructField.Anonymous:
					conv.toValues(v.Field(i).Interface())
				}

			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				conv.vurlValues.Set(name, strconv.FormatInt(innerv.Int(), 10))
			case reflect.Bool:
				conv.vurlValues.Set(name, strconv.FormatBool(innerv.Bool()))
			case reflect.String:
				conv.vurlValues.Set(name, innerv.String())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				conv.vurlValues.Set(name, strconv.FormatUint(innerv.Uint(), 10))
			case reflect.Float32, reflect.Float64:
				conv.vurlValues.Set(name, strconv.FormatFloat(innerv.Float(), 'E', 10, 64))

			}
		}

	}
}
