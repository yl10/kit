package guid

import (
	"database/sql/driver"

	"github.com/pborman/uuid"
)

var (
	//SpaceUUID 空GUID，用于生成不变化的GUID
	SpaceUUID = uuid.Parse("f2093908-9293-41f0-97a6-413e94f788ef")
)

//GUID guid
type GUID string

//Value GUID 实现database/driver中的接口，否则部分场景下会抛出异常。
func (g GUID) Value() (driver.Value, error) {
	return string(g), nil
}

//NewGUID 生成一个唯一的GUID
func NewGUID() GUID {
	return GUID(uuid.New())
}

//NewGUIDString 生成一个GUID string
func NewGUIDString() string {
	return uuid.New()
}

//NewMD5GUID 根据输入参数生成一个GUID
func NewMD5GUID(str string) GUID {

	return GUID(uuid.NewMD5(SpaceUUID, []byte(str)).String())
}
