package request

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	UA_MOBILE = "mobile"
	UA_WECHAT = "MicroMessenger"
)

// Request 请求远程 url 数据
// 如果有传递的参数, 默认是用 post 方法
// 如果有需要 get 传递的参数, 请自行组装到 url 里面
func Request(url string, https bool, params ...url.Values) ([]byte, error) {

	var client *http.Client
	var resp *http.Response
	var err error

	if https {
		// 不对 https 的证书验证
		tr := &http.Transport{
			TLSClientConfig:    &tls.Config{InsecureSkipVerify: true},
			DisableCompression: true,
		}
		client = &http.Client{Transport: tr}
	} else {
		client = &http.Client{}
	}

	// 有参数说明是POST模式
	if len(params) > 0 {
		resp, err = client.PostForm(url, params[0])
	} else {
		resp, err = client.Get(url)
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(resp.Body)
}

// Input 获取请求参数包含 url query string, get、post params
// 限制，不能获取 head 及 body 中参数
func Input(r *http.Request) (url.Values, error) {
	val := url.Values{}

	if r.Form == nil {
		if err := r.ParseForm(); err != nil {
			return val, err
		}
	}

	// 如果是非 multipart/form-data 提交
	// 则直接返回
	ct := r.Header.Get("Content-Type")
	val = r.Form
	if ct != "multipart/form-data" {
		return val, nil
	}

	// 解析 multipart/form-data
	// 如果存在，则替换 ParseForm 的解析结果
	if r.MultipartForm == nil {
		if err := r.ParseMultipartForm(1 << 22); err != nil {
			return val, err
		}
	}

	multiForm := *r.MultipartForm
	for k, v := range multiForm.Value {
		val.Del(k)
		val[k] = v
	}

	return val, nil
}

// IsMobile 是否手机客户端
func IsMobile(ua string) bool {
	return strings.Contains(strings.ToLower(ua), strings.ToLower(UA_MOBILE))
}

// IsWeiXin 是否微信客户端
func IsWeiXin(ua string) bool {
	return strings.Contains(strings.ToLower(ua), strings.ToLower(UA_WECHAT))
}
