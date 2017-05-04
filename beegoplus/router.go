package beegoplus

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/pborman/uuid"
)

//Router 路由
type Router struct {
	Name       string
	Pattern    string
	Method     []string
	ActionName string
	HeadParam  []Param
	QueryParam []Param
	BodyParam  []Param
	PathParam  []Param
	BodyOutPut []OutParam
	outPut     *OutPut
}

//GetGUID 获取路由的GUID值
func (r Router) GetGUID() string {
	return uuid.NewMD5(spaceuuid, []byte(r.Name)).String()
}

//CheckParam 检查参数
func (r Router) CheckParam(req *http.Request) error {

	//检查head参数
	if err := r.handParam(HEADLOCATION, url.Values(req.Header)); err != nil {
		return err
	}

	//检查body参数
	if len(r.BodyParam) > 0 || len(r.QueryParam) > 0 {

		if err := req.ParseForm(); err != nil {
			return err
		}
		//检查query参数
		if err := r.handParam(QUERYLOCATION, req.Form); err != nil {
			return err
		}
		if err := r.handParam(BODYLOCATION, req.PostForm); err != nil {
			return err
		}
	}

	return nil
}

func (r Router) handParam(location ParamLocation, params url.Values) error {

	var definedparam []Param
	switch location {
	case HEADLOCATION:
		definedparam = r.HeadParam
	case QUERYLOCATION:
		definedparam = r.QueryParam
	case PATHLOCATION:
		definedparam = r.PathParam

	case BODYLOCATION:
		definedparam = r.BodyParam
	}

	for _, param := range definedparam {
		value := params.Get(param.Name)
		//检查必填
		if param.Must && value == "" {
			return fmt.Errorf("缺少必须参数：%s", param.Name)
		}

		values := []string{}
		if param.Muilti {
			values = params[param.Name]
		} else {
			values = append(values, value)
		}

		switch param.Type {
		case STRINGTYPE:

			for _, v := range values {
				//检查最大长度
				if param.MaxLen > 0 && len(v) > param.MaxLen {
					return fmt.Errorf("参数 %s 长度超过最大长度 %d", param.Name, param.MaxLen)
				}

				//检查枚举
				if len(param.Enum) > 0 {

					if !inEnum(v, param.Enum) {
						return fmt.Errorf("请求参数(%s)不在规定的枚举范围内：%s", param.Name, v)
					}

				}

			}

		case NUMBERTYPE:

			for _, v := range values {
				if strings.Contains(v, ".") {
					//带小数点
					_, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return fmt.Errorf("非法的参数：%s", param.Name)
					}
				} else {
					//不带小数点
					_, err := strconv.Atoi(v)
					if err != nil {
						return fmt.Errorf("非法的参数：%s", param.Name)
					}
				}
				//检查枚举
				if len(param.Enum) > 0 {

					if !inEnum(v, param.Enum) {
						return fmt.Errorf("请求参数(%s)不在规定的枚举范围内：%s", param.Name, v)
					}

				}

			}

		case BOOLEANTYPE:

			for _, v := range values {
				_, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("非法的参数：%s", param.Name)
				}
			}

		}
		//正则判断
		if param.Pattern != "" {
			for _, v := range values {
				matched, err := regexp.MatchString(param.Pattern, v)
				if err != nil {
					return fmt.Errorf("对参数(%s)进行正则解析出错：%v", param.Name, err)
				}
				if !matched {
					return fmt.Errorf("参数(%s)不符合正则要求", param.Name)
				}
			}

		}
	}
	return nil
}

// func (r Router) ServerJson(resp *http.ResponseWriter) error {
// 	resp.Write(r.ou)
// }

//NewOutPut 初始化一个输出
func (r Router) NewOutPut(code int, result interface{}) *OutPut {
	op := newOutPut(code, result)
	return op
}

//FilterOut 过滤输出
func (r Router) FilterOut(authmap map[string]bool) map[string]interface{} {

	result := make(map[string]interface{})
	if r.outPut == nil {
		result["code"] = 0
		result["message"] = "没有任何输出"
		return result
	}
	if r.outPut.code != 0 {
		result["message"] = r.outPut.message
		result["code"] = r.outPut.code

		return result
	}

	for _, v := range r.BodyOutPut {

		hasAuth := func(key string) bool {
			has, ok := authmap[key]
			if has {
				return true
			}
			if !ok {
				return true
			}
			return false
		}(v.Name)

		if hasAuth {

			if v.Muilti {
				result[v.Name] = r.outPut.data[v.Name]
			} else {
				if len(r.outPut.data[v.Name]) > 0 {
					result[v.Name] = r.outPut.data[v.Name][0]
				} else {
					result[v.Name] = v.Type.Empty()
				}

			}

		}

	}
	return result
}
