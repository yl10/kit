# struct转其他类型 #
## 通用说明 ##
1. tag的使用方法参考json，默认使用yl，如
```go
type foo struct{
	A string `yl:"a"`// 导出名为 a
	B string `yl:"-"` //B,不导出
}
```
## 功能说明 ##
### url.Values ###
```
仅对字符、数字、布尔、日期类型进行转化，其他类型忽略;闭包字段会展开进行解析，非闭包会忽略
```

1. ToURLValues(o interface{}) url.Values
2. ToURLValuesWithTimeFormart(o interface{}, timeformt string) url.Values 

