package orm

import (
	"regexp"
	"strings"
)

// hyphen2Hump ab_ef => AbEf
func hyphen2Hump(str string) string {
	re := regexp.MustCompile("(^[[:alnum:]]|_[[:alnum:]])")
	return re.ReplaceAllStringFunc(str, func(m string) string {
		return strings.Trim(strings.ToUpper(m), "_")
	})
}
