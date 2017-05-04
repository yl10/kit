package stringutil

import (
	"regexp"
	"strings"
	"unicode"
)

// ReplaceSubString 替换定长字串
// 将 s 字串，从 start (包含) 开始的位置, 替换为 c , 替换长度 length
// 如果 c 是单字符, 那么 s 被替换后的长度不变
// 一般场景类似将手机号的中间4位转换为 *
func ReplaceSubString(s, c string, start, length int) string {

	a := strings.Split(s, "")

	for inx, _ := range a {
		if inx >= start && inx <= start+length-1 {
			a[inx] = c
		}
	}

	return strings.Join(a, "")
}

// GetDateFormat 获取 go 的日期格式化字串
// golang 的日期格式化必须要用  Mon Jan 2 15:04:05 -0700 MST 2006 这个日期
// 这个让我很不爽, 增加一个通用类型的转换
func GetDateFormat(commstr string) string {
	if len(commstr) == 0 {
		return ""
	}

	type fmtpt struct {
		fmtstr string
		ptstr  string
	}

	// year
	fullYear := fmtpt{fmtstr: "Y", ptstr: "2006"}
	shortYear := fmtpt{fmtstr: "y", ptstr: "06"}

	// month
	textMonth := fmtpt{fmtstr: "M", ptstr: "Jan"}
	numMonth := fmtpt{fmtstr: "m", ptstr: "01"}

	// day
	textDay := fmtpt{fmtstr: "D", ptstr: "Mon"}
	numDay := fmtpt{fmtstr: "d", ptstr: "02"}
	numNofixDay := fmtpt{fmtstr: "j", ptstr: "2"}

	// hour
	h24Hour := fmtpt{fmtstr: "H", ptstr: "15"}
	h12Hour := fmtpt{fmtstr: "h", ptstr: "03"}

	// minute
	zPreMinute := fmtpt{fmtstr: "i", ptstr: "04"}

	// second
	zPreSecond := fmtpt{fmtstr: "s", ptstr: "05"}

	// GMT
	gwTime := fmtpt{fmtstr: "O", ptstr: "-0700"}

	// timezone
	tz := fmtpt{fmtstr: "T", ptstr: "MST"}

	//
	matchStruct := []fmtpt{
		fullYear,
		shortYear,
		textMonth,
		numMonth,
		textDay,
		numDay,
		numNofixDay,
		h24Hour,
		h12Hour,
		zPreMinute,
		zPreSecond,
		gwTime,
		tz,
	}

	matchChars := strings.Split(commstr, "")

	var parseStr string
	for _, ch := range matchChars {
		inStruct := false
		for _, f := range matchStruct {
			if ch == f.fmtstr {
				inStruct = true
				parseStr = parseStr + f.ptstr
				break
			}
		}
		if !inStruct {
			parseStr = parseStr + ch
		}
	}

	return parseStr
}

// Hump2Hyphen 驼峰写法转连字符写法, 并转化首字符小写
func Hump2Hyphen(str string) string {
	re := regexp.MustCompile("([A-Z])")
	rtn := re.ReplaceAllStringFunc(str, func(m string) string {
		return "_" + strings.ToLower(m)
	})
	return strings.TrimPrefix(rtn, "_")
}

// HasChinese 是否包含汉字
func HasChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}

	return false
}
