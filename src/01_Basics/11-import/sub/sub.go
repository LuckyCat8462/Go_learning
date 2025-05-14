package sub

//  在go语言中，同一层级目录，不允许出现多个包名

func Sub(a, b int) int {
	test4() // 由于test4与sub.go在同一个包下面，所以可以使用，并且不需要sub.形式
	return a - b
}
func sub1(a, b int) int {
	return a - b + 1
}
