package main

import "fmt"

// 一、函数概念
// 函数是基本的代码块，用于执行一个任务。可以利用不同的函数来划分不同的功能
// 1.1、函数定义
// 1.1.1格式：
//	func function_name( [parameter list] ) [return_types] {
//	  函数体
//	}
//	function_name：函数名称，参数列表和返回值类型构成了函数签名。
//	parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
//	return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。
//	函数体：函数定义的代码集合。

// 二、实例
// 2.1、函数返回值在参数列表之后
// 2.2、如果有多个返回值，需要使用圆括号包裹，多个参数之间使用,分割
func test2(a int, b int, c string) (int, string, bool) {
	return a + b, c, true
}

// func function_name( [parameter list] ) [return_types]
func test3(a, b int, c string) (res int, str string, bl bool) {

	var i, j, k int
	i = 1
	j = 2
	k = 3
	fmt.Println(i, j, k)

	// 2.3.返回值有变量名
	// 2.3.1.直接使用返回值的变量名字参与运算
	res = a + b
	str = c
	bl = true

	// 2.3.2.当返回值有名字的时候，可以直接简写return
	return

	// return a + b, c, true
}

// 2.4如果返回值只有一个参数，并且没有名字，那么不需要加圆括号
func test4() int {
	return 10
}

// 暂时不确定的参数可以用_下横线代替
func main() {
	v1, s1, _ := test2(10, 20, "hello")
	fmt.Println("v1:", v1, ",s1:", s1)
	// 使用ctrl+shift+i，可以快速的查看函数定义
	v3, s3, _ := test3(20, 30, "world")
	fmt.Println("v3:", v3, ", s3:", s3)
}
