# 加密小工具
大部分是对字串的处理

- HmacSha1 使用密钥 k 对 消息 s 进行加密
```go
func HmacSha1(s string, k string) string
```

- Base64Encode 对字串 s 进行 Base64 加密
```go
func Base64Encode(s string) string
```

- Md5 对字串 key 进行 MD5 加盐 with 处理
```go
func Md5(key string, with ...string) string
```
