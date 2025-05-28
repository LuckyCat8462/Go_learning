package main

import "fmt"

// 一、内存逃逸 概念
// 编译时，由于某种原因，编译器决定将本应该分配到 栈 的变量分配到了 堆 中。

func main() {
	p1 := testPtr1()
	fmt.Println("*p1:", *p1)
}

// 要想看到是否编译，可以使用指令
// go build -o test.exe --gcflags "-m -m -l" ./1-Basics/10-内存逃逸.go

// 定义一个函数，返回一个string类型的指针, go语言返回写在参数列表后面
func testPtr1() *string {
	name := "neko"
	p0 := &name
	fmt.Println("*p0:", *p0)
	// 可以发现name不需要逃逸
	city := "杭州"
	ptr := &city
	return ptr
	// 而city逃逸了
}
