# kit组件包 #
组件包,可以引入kit，也可以单独引入各个需要的子包

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
[详情](./structconv/readme.md)


## dateutil ##
日期工具集 [详情](./dateutil/README.md)
- DateRange 获取指定间隔的日期 slice
- DayRange 获取当前的起止时间 00:00:00 - 23:59:59


## cube ##
复杂的转换函数 [详情](./cube/README.md)
- InputToStruct url.Values 映射为 struct


## encrypt ##
加解密函数 [详情](./encrypt/README.md)
- HmacSha1 HmacSha1加密
- Base64Encode Base64 加密
- Md5 MD5加密


## random ##
随机串函数 [详情](./random/README.md)
- GenRandInt 获取 [0, r) 之间的一个随机数字
- GenRandAphla 获取有字母(大小写区分)和数字组成的随机数


## stringutil ##
字串工具集 [详情](./stringutil/README.md)
- ReplaceSubString 替换指定内容
- GetDateFormat 通用日期格式化字串
- Hump2Hyphen 驼峰转连字符
- HasChinese 字串中是否含有汉字


## websocket ##
简单 websocket [详情](./websocket/README.md)


## 2017-09-28 增加 orm 包