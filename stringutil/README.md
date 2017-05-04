# 字串类小工具


- ReplaceSubString
将 s 字串，从 start (包含) 开始的位置, 替换为 c , 替换长度 length
如果 c 是单字符, 那么 s 被替换后的长度不变
一般场景类似将手机号的中间4位转换为 *
```go
func ReplaceSubString(s, c string, start, length int) string
```


- GetDateFormat
获取 go 的日期格式化字串
golang 的日期格式化必须要用  Mon Jan 2 15:04:05 -0700 MST 2006 这个日期
这个让我很不爽, 增加一个通用类型的转换


转换说明如下
| 通用格式       | Golang 格式   | 说明          |
| ------------- |:--------------|:--------------|
| Y             | 2006          | 4 位年        |
| y             | 06            | 2 位年        |
| M             | Jan           | 英文 月       |
| m             | 01            | 2 位月        |
| D             | Mon           | 英文 周几      |
| d             | 02            | 2 位天        |
| j             | 2             | 1 位天        |
| H             | 15            | 24小时制 小时  |
| h             | 03            | 12小时制 小时  |
| i             | 04            | 2 位分         |
| s             | 05            | 2 位秒         |

Demo
```go
// 返回 2006-06-01 15:04:05
GetDateFormat("Y-m-d H:i:s")
```


- Hump2Hyphen 驼峰写法转连字符写法, 并转化首字符小写
```go
func Hump2Hyphen(str string) string
```


- HasChinese 是否包含汉字
```go
func HasChinese(str string) bool
```
