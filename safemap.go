//safemap
//从谢大的beemap copy过来的

package kit

import (
	"github.com/yl10/kit/safemap"
)

//NewSafeMap 返回一个新的安全map指针
func NewSafeMap() *safemap.SafeMap {
	return safemap.NewSafeMap()
}
