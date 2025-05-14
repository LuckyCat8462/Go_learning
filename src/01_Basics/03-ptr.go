package main

// 3-指针

import "fmt"

func main() {
	// 一、概念
	// 一个指针变量指向了一个值的 内存地址
	// 1.1、指针声明
	// 格式：var var_name *var-type
	// var str1 *string
	// 1.1.1
	name := "neko"
	ptr := &name

	// p是一个指针变量的名字，表示此指针变量指向的内存地址，如果使用%p来输出的话，它将是一个16进制数。
	// *p表示此指针指向的内存地址中存放的内容，一般是一个和指针类型一致的变量或者常量。
	// &是取地址运算符，&p就是取指针p的地址，&p就表示编译器为变量p分配的内存地址，而因为p是一个指针变量，这种特殊的身份注定了它要指向另外一个内存地址，程序员按照程序的需要让它指向一个内存地址，这个它指向的内存地址就用p表示
	fmt.Println("&name: name变量的地址是:", &name)
	fmt.Println("ptr: ptr变量存储的指针地址是:", ptr)
	fmt.Println("*ptr: 使用ptr指针访问值是:", *ptr)

	// 1.1.2.使用new关键字定义
	name2Ptr := new(string)
	*name2Ptr = "zhi"

	fmt.Println("name2访问得到的值:", *name2Ptr)
	fmt.Println("name2 ptr存储的指针地址:", name2Ptr)

	// 1.2.结构体调用
	// 结构体成员调用时：   c语言:   ptr->name  go语言:  ptr.name
	// go语言在使用指针时，会使用内部的垃圾回收机制(gc : garbage collector)，开发人员不需要手动释放内存
	// c语言不允许返回栈上的指针，go语言可以返回栈上的指针，程序会在编译的时候就确定了变量的分配位置：
	// 编译的时候，如果发现有必要的话，就将变量分配到堆上

	// 1.3.空指针
	// 在c语言： null， c++: nullptr，go： nil
	// 可以返回栈上的指针，编译器在编译程序时，会自动判断这段代码，将city变量分配在堆上, 内存逃逸
	res := testPtr()
	fmt.Println("res city :", *res, ", address:", res)

	_ = testPtr()

	// (a) if两端不用加()
	// (b) 即使有一行代码，也必须使用{}
	if res == nil {
		fmt.Println("res 是空，nil")
	} else {
		fmt.Println("res 是非空")
	}
	// 二、注意点
	// 2.1.地址不允许加减
	// fmt.Println("ptr +1:", *(ptr + 1))

}

// 定义一个函数，返回一个string类型的指针, go语言返回写在参数列表后面
func testPtr() *string {
	city := "杭州"
	ptr := &city
	return ptr

}
