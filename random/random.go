package random

import (
	"math/rand"
	"strings"
	"time"
)

// GenRandInt 获取 [0, r) 之间的一个随机数字
func GenRandInt(r int) int {
	inst := rand.New(rand.NewSource(time.Now().UnixNano()))
	return inst.Intn(r)
}

// GenRandAphla 获取有字母(大小写区分)和数字组成的 l 长度的随机数
func GenRandAphla(l int) string {

	inst := rand.New(rand.NewSource(time.Now().UnixNano()))
	dict := strings.Split("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", "")
	le := len(dict)

	var str = ""
	for i := 0; i < l; i++ {
		p := inst.Intn(le)
		str += dict[p]
	}

	return str
}
