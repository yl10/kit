package slice

// Filter 过滤 sliece 中符合 f 函数定义的元素
func Filter(s interface{}, f func(interface{}, interface{}) bool) interface{} {
	return From(s).Filter(f)
}
