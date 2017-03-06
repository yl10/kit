//Package beegoplus beego的插件
package beegoplus

const (
	//HEADLOCATION 头部
	HEADLOCATION ParamLocation = "HEAD"
	//QUERYLOCATION QUERY
	QUERYLOCATION = "QUERY"
	//BODYLOCATION BODY
	BODYLOCATION = "BODY"
	//PATHLOCATION PATH
	PATHLOCATION = "PATH"
)

const (
	//STRINGTYPE 字符
	STRINGTYPE ParamType = "String"
	//NUMBERTYPE 数字
	NUMBERTYPE = "Number"
	//BOOLEANTYPE 布尔
	BOOLEANTYPE = "boolean"
	//OBJECTTYPE 对象
	OBJECTTYPE = "object"
)

//ParamLocation type:参数位置
type ParamLocation string

//ParamType 参数类型
type ParamType string

//Param 参数
type Param struct {
	Location    ParamLocation //
	Name        string
	Type        ParamType
	MaxLen      int
	Enum        []interface{}
	DefautValue interface{}
	Muilti      bool //是否数组
	Must        bool
	Pattern     string
	Description string
}

//ParamValues 参数值
type ParamValues map[string][]interface{}

//Empty 空值
func (p ParamType) Empty() interface{} {
	switch p {
	case BOOLEANTYPE:
		return false
	case STRINGTYPE:
		return ""
	case NUMBERTYPE:
		return 0
	}
	return nil
}
