package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// HmacSha1 使用密钥 k 对 消息 s 进行加密
func HmacSha1(s string, k string) string {
	key := []byte(k)
	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(s))
	res := mac.Sum(nil)
	return string(res[:])
}

// Base64Encode 对字串 s 进行 Base64 加密
func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// Md5 对字串 key 进行 MD5 加盐 with 处理
func Md5(key string, with ...string) string {

	var md5Builder = md5.New()

	md5Builder.Write([]byte(key + strings.Join(with, "")))

	return hex.EncodeToString(md5Builder.Sum(nil))
}
