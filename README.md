# kit组件包 #
```
组件包,可以引入kit，也可以单独引入各个需要的子包
```
## safemap ##
加锁的map，直接copy自谢大的beemap
```go
func NewSafeMap() *safemap.SafeMap
```
## structconv ##
struct和其他格式的转换
```go
func StructToURLValues(o interface{}) url.Values
func StructToURLValuesWithTimeFormat(o interface{}, timeformat string) url.Values 
```
