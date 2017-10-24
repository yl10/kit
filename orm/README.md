# XormPlus 是 [Xorm](https://github.com/go-xorm/xorm) 的简单扩展

## 说明
把目前业务上经常用的方法，进行了简单封装，也兼容 Xorm 的方法。

## 介绍
### Fetch 获取单行记录
```go
// inst 为 Table 对象指针, filter 为过滤条件, conds 为过滤条件中用到的参数
func Fetch(inst interface{}, filter interface{}, conds ...interface{}) error
```

### List 获取多行记录
```go
// inst 为 Table 对象指针, filter 为过滤条件, conds 为过滤条件中用到的参数
func List(inst interface{}, filter interface{}, conds ...interface{}) error
```

### Update 更新记录
```go
// inst 为 Table 对象指针, cols 指定强制更新的字段, filter 为过滤条件, conds 为过滤条件中用到的参数
func Update(inst interface{}, cols []string, filter interface{}, conds ...interface{}) (int64, error)
```

### SetPageNavi 设置分页参数
```go
// setting 为设置的参数内容, 目前只有 pageNum
func SetPageNavi(setting map[string]int) *XormPlus
```

### Page 设置分页
```go
// n 为获取的分页值, perNum 为每页数据量
// 这个方法，实际上只是加了一个 limit
func Page(n int, perNum ...int) *XormPlus
```

### SetFieldMap 设置字段映射方式
```go
// Collention[s] 等方法，是使用的原生 Sql 来查询的
// 这就涉及到原生结果字段映射的问题
// SetFieldMap 方法需要接受一个函数, 该函数的入参是一个字符串, 出参也是一个字符串
// 入参字符串就是原生结果的字段名称, 出参就是映射的名称
// 默认的是蛇形转驼峰, 比如 字段名为 user_name, 那么结果集就会变成 UserName
func SetFieldMap(fieldMap func(string) string) *XormPlus
```

### Collention 获取单行记录, 结果集为 Map
```go
// 依据原生sql 查询结果
// sqlStr 是原生sql, args 是传给 sql 的参数
func Collention(sqlStr string, args ...interface{}) (map[string]string, error)
```

### Collention 获取多记录, 结果集为 []Map
```go
// 同 Collention 方法一致
// 只不过返回的是多条记录
func Collention(sqlStr string, args ...interface{}) (map[string]string, error)
```

### RawCollention 获取单行记录, 结果集为 Map
```go
// 依据原生sql 查询结果
// sqlStr 是原生sql, args 是传给 sql 的参数
func RawCollention(sqlStr string, args ...interface{}) (map[string]interface{}, error)
```

### RawCollention 获取多记录, 结果集为 []Map
```go
// 同 RawCollention 方法一致
// 只不过返回的是多条记录
func RawCollention(sqlStr string, args ...interface{}) (map[string]interface{}, error)
```
