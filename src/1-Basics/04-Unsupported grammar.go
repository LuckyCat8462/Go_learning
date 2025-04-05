package main

// 4-不支持的语法

import "fmt"

// 1.不支持自增自减中的--i,++i
// 2.不支持地址加减
// 3.不支持三目运算符 ?:
// 4.只有false才能代码逻辑假，数字0与nil不能

func main() {
	// if 0 {
	//	fmt.Println("0可以逻辑假")
	// }
	//
	// if nil {
	//	fmt.Println("nil可以逻辑假")
	// }

	if true {
		fmt.Println("false可以逻辑假")
	}

	// a, b := 1, 2
	// x := a > b ? 0:-1
}
