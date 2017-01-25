package kit

import (
	"net/url"

	"github.com/yl10/kit/structconv"
)

//StructToURLValues 转换为url.Values
//如果传入参数不是struct类型，返回空的values
func StructToURLValues(o interface{}) url.Values {

	return structconv.ToURLValues(o)

}

//StructToURLValuesWithTimeFormat 带时间格式把struct转为为urlvalues
//输入:o struct
//输入:timeformat 日期格式
func StructToURLValuesWithTimeFormat(o interface{}, timeformat string) url.Values {
	return structconv.ToURLValuesWithTimeFormart(o, timeformat)
}
