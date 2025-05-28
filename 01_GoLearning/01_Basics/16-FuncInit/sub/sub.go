package sub

import "fmt"

// 1.init函数没有参数，没有返回值，原型固定如下
// 2.一个包中包含多个init时，调用顺序是不确定的(同一个包的多个文件中都可以有init)
// 3. init函数时不允许用户显示调用的
// 4. 有的时候引用一个包，可能只想使用这个包里面的init函数（mysql的init对驱动进行初始化）
// 但是不想使用这个包里面的其他函数，为了防止编译器报错，可以使用_形式来处理
// import _ "xxx/xx/sub"
func init() {
	fmt.Println("this is first init() in package sub ==> sub.go")
}

func init() {
	fmt.Println("this is second init() in package sub ==> sub.go ")
}

// 在go语言中，同一层级目录，不允许出现多个包名

func Sub(a, b int) int {
	// init() ==> 不允许显示调用
	test4() // 由于test4与sub.go在同一个包下面，所以可以使用，并且不需要sub.形式
	return a - b
}
