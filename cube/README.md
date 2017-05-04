# 一些复杂的特定类型转换器

- InputToStruct url.Values 映射为 struct
```go
// 将 url.Values 中的数据, 转入到 struct 中
// 一般使用场景是, 将 request 传入的 params 转到方便操作的 struct 中
// 一个缺点是, 数组会转为 逗号连接的字串
// s 必须为指针
func InputToStruct(input url.Values, s interface{}) error
```

