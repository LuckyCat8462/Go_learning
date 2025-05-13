package main

//	0-hello world 测试用例

//	每个go程序，都是.go结尾的
//	每个go程序，都有一个package main

//	这是导入一个标准包fmt，format，一般用于格式化输出
import "fmt"

// 主函数，所有的函数都必须使用func开头
// 一个函数的返回值，不会放在func前，而是放在参数后面
// 函数左括号必须要与函数名同行，不能写在下一行
func main() {
	//	go语言语句不需要使用分号结尾
	fmt.Println("hello world")
}
