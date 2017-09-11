# 关于 http reqeust 的小工具集

- Request 请求远程 url 数据
```go
// 如果有传递的参数, 默认是用 post 方法
// 如果有需要 get 传递的参数, 请自行组装到 url 里面
func Request(url string, https bool, params ...url.Values) ([]byte, error)
```


- Input 获取请求参数包含 url query string, get、post params
```go
// 不能获取 head 及 body 中参数
func Input(r *http.Request) (url.Values, error)
```


- IsMobile 是否手机客户端


- IsWeiXin 是否微信客户端

- Url 完整的URL