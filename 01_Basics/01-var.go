package main

//	1-变量

//	goland会帮我们自动导入程序中使用的包
import "fmt"

func main() {
	//	一、概念
	//	变量可以通过变量名访问。
	//	Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。
	//	变量定义：var，常量定义： const

	//	1.1.变量声明
	//	格式：var var_name var_type
	var name string
	//	1.1.1.先定义变量，再赋值
	name = "neko"
	//	(a) 行分隔符
	//	在Go程序中，一行代表一个语句结束。不需要像其他语言一样以分号为结尾

	// ctrl+alt+l 可以快速格式化代码

	var age int
	age = 25
	fmt.Println("name", name)
	fmt.Printf("name is :%s,%d\n", name, age)

	// 1.1.2.定义时直接赋值
	var gender = "man"
	fmt.Println("gender", gender)

	// 1.1.3.定义直接赋值，使用自动推导(常用)
	address := "杭州"
	fmt.Println("address", address)
	// 灰色的内容是形参，自动出现
	test(10, "str")

	// 1.1.4.平行赋值
	i, j := 10, 20
	fmt.Println("变换前——》", "i值:", i, "j值:", j)

	i, j = j, i
	fmt.Println("变换后i值：", i, "变换后,j值:", j)

}
func test(a int, b string) {
	fmt.Println(a)
	fmt.Println(b)
}
