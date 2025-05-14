package main

// 5-string 字符串类型

import "fmt"

func main() {
	// 一、概念
	// 字符串就是一串固定长度的字符连接起来的字符序列。
	// Go 的字符串是由单个字节连接起来的。
	// Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。
	name := "neko"

	// 二、使用
	// 2.1需要换行，原生输出字符串时，使用反引号``
	usage := `./a.out <option> 
         -h   help 
		 -a  xxxx`
	// c语言使用单引号 + \来解决
	fmt.Println("name :", name)
	fmt.Println("usage :", usage)

	// 2.2.字符串长度len()
	// C++: name.length
	// GO: string没有.length方法，可以使用自由函数len()进行处理
	l1 := len(name)
	fmt.Println("name字符串的长度为:", l1)

	// for循环不需要加()
	for i := 0; i < len(name); i++ {
		fmt.Printf("i: %d, v: %c\n", i, name[i])
	}

	// 2.3.字符串拼接
	i, j := "hello", " world"
	fmt.Println("i+j=", i+j)

	// (a) 使用const修饰为常量，不能修改
	const address = "杭州" // 在编译期就确定了类型，是预处理，所以不需要推导
	const test = 100
	fmt.Println("address:", address)

	if i > j {

	}
}
