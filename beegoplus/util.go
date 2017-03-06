package beegoplus

import (
	"fmt"
	"strings"

	"github.com/pborman/uuid"
)

var spaceuuid = uuid.NameSpace_DNS

//SetSpaceUUID 设置空UUID的值，会影响到包中所有UUID的生成
func SetSpaceUUID(s string) {
	spaceuuid = uuid.NewMD5(uuid.NameSpace_DNS, []byte(s))
}

//inEnum 判断值是否在枚举类型里，主要用于字符和数字，其他不做处理
func inEnum(v interface{}, enum []interface{}) bool {
	list := []string{}
	for _, e := range enum {
		list = append(list, fmt.Sprintf("%v", e))
	}

	sep := ":::::::::::::::::::::::"
	return strings.Contains(strings.Join(list, sep)+sep, fmt.Sprintf("%v%s", v, sep))

}
