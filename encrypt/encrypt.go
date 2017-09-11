package encrypt

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strings"
)

const (
	//SortByASCIISmallToBig 按照ASCII从小到大排
	SortByASCIISmallToBig SortType = iota
	//SortByASCIIBigToSmall 按照ASCII从大到小排
	SortByASCIIBigToSmall
)

//SortType 排序方式，用于一些加密
type SortType int

const (
	//NoChangeCase 不改变大小写
	NoChangeCase CaseType = iota
	//LowerCase 小写
	LowerCase
	//UpperCase 大写
	UpperCase
)

//CaseType 大小写处理方式
type CaseType int

const (
	//NotEncoding 不进行编码
	NotEncoding EncodingType = iota
	//URLEncoding URLencding
	URLEncoding
	//Base64Encoding Base64Encoding
	Base64Encoding
)

//EncodingType 编码方式
type EncodingType int

const (
	//EncryptMD5 md5加密
	EncryptMD5 EncryptTyp = iota
	//EncryptSHA1 sha1
	EncryptSHA1
	//EncryptSHA256 SHA256
	EncryptSHA256

	EncryptBase64
)

//EncryptTyp 加密方式
type EncryptTyp int

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

//MapEncode 对map进行加密处理
//encryptType 加密方式
//params 待签名的map
//keySort key的排序方式。SortByASCIISmallToBig|SortByASCIIBigToSmall
//keyCaseType key是否要做大小写处理 NoChangeCase|LowerCase|UpperCase
//encodingvalue 对value进行编码的方式 。NotEncoding|URLEncoding|Base64Encoding
//joinwith 每对KV之间的连接符，比如 &
//k2vwith k v之间的间隔符，比如= ：
//saltwith 另外加盐处理
func MapEncode(encryptType EncryptTyp, params map[string]string, keySort SortType, keyCaseType CaseType, encodingvalue EncodingType, joinwith string, k2vwith string, saltwith ...string) (string, error) {

	//先把key放到slice里

	type advMap struct {
		ID       int
		OldKey   string
		NewKey   string
		Value    string
		NewValue string
	}

	maps := make([]advMap, len(params))

	i := 0
	for k, v := range params {
		maps = append(maps, advMap{
			ID:     i,
			OldKey: k,
			NewKey: caseString(v, keyCaseType),
			Value:  v,
		})
		i++
	}

	//对key进行排序处理
	switch keySort {
	case SortByASCIISmallToBig:

		sort.SliceStable(maps, func(i, j int) bool {
			return maps[i].NewKey < maps[j].NewKey
		})

	case SortByASCIIBigToSmall:

		sort.SliceStable(maps, func(i, j int) bool {
			return maps[i].NewKey > maps[j].NewKey
		})

	}

	//对value进行编码
	switch encodingvalue {
	case URLEncoding:
		for k, v := range maps {
			v.NewValue = url.QueryEscape(v.Value)
			maps[k] = v
		}
	case Base64Encoding:
		for k, v := range maps {
			v.NewValue = Base64Encode(v.Value)
			maps[k] = v
		}
	default:
		return "", fmt.Errorf("暂未支持的编码方式:%v", encodingvalue)
	}

	//拼字符串
	str := ""
	for _, v := range maps {
		str = str + joinwith + v.NewKey + k2vwith + v.NewValue
	}
	//去掉第一个连接符
	str = strings.TrimLeft(str, joinwith)

	//加上加密盐

	salt := strings.Join(saltwith, "")

	byteStr := []byte(str)
	byteSalt := []byte(salt)
	//加密处理
	switch encryptType {
	case EncryptMD5:
		md5Hash := md5.New()
		md5Hash.Write(byteStr)
		return fmt.Sprintf("%x", md5Hash.Sum(byteSalt)), nil
	case EncryptSHA1:
		sha1Hash := sha1.New()
		sha1Hash.Write(byteStr)
		return fmt.Sprintf("%x", sha1Hash.Sum(byteSalt)), nil
	case EncryptSHA256:
		sha256Hash := sha256.New()
		sha256Hash.Write(byteStr)
		return fmt.Sprintf("%x", sha256Hash.Sum(byteSalt)), nil

	default:
		return "", fmt.Errorf("未支持的加密类型:%v", encryptType)
	}

}

func caseString(s string, ctype CaseType) string {
	switch ctype {
	case UpperCase:
		return strings.ToUpper(s)
	case LowerCase:
		return strings.ToLower(s)
	default:
		return s
	}
}
