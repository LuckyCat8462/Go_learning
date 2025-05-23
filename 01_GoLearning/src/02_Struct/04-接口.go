package main

import "fmt"

// 一、接口的概念
// 1.1.接口是编程中一种重要的数据类型，用于实现多态。
// 1.2.C语言中没有接口的概念，但可以通过其他方式实现多态，C++和JAVA使用纯虚函数来实现接口。
// 1.3.Go语言中使用了专门的关键字interface来表示接口。interface不仅仅是用于处理多态的，它可以接受任意的数据类型，有点类似void

func main() {
	// 可变参数+接口的形式
	// func Println(a ...interface{}) (n int, err error) {
	fmt.Println("")

	// var i,j,k int
	// 定义三个接口类型
	var i, j, k interface{} // 空接口
	names := []string{"neko", "zhi"}
	i = names
	fmt.Println("i代表切片数组:", i)

	age := 20
	j = age
	fmt.Println("j代表数字：", j)

	str := "hello"
	k = str
	fmt.Println("k代表字符串:", k)

	// 我们现在只知道k是interface，但是不能够明确知道它代表的数据的类型
	// 所以有时候要加上一个判断
	kvalue, ok := k.(int)
	if !ok {
		fmt.Println("k不是int")
	} else {
		fmt.Println("k是int， 值为:", kvalue)
	}

	// 最常用的场景： 把interface当成一个函数的参数，（类似于print），使用switch来判断用户输入的不同类型
	// 根据不同类型，做相应逻辑处理

	// 创建一个具有三个接口类型的切片
	array := make([]interface{}, 3)
	array[0] = 1
	array[1] = "Hello world"
	array[2] = 3.14

	for _, value := range array {
		switch v := value.(type) { // 可以获取当前接口的真正数据类型
		case int:
			fmt.Printf("当前类型为int, 内容为：%d\n", v)
		case string:
			fmt.Printf("当前类型为string, 内容为: %s\n", v)
		case bool:
			fmt.Printf("当前类型为bool, 内容为: %v\n", v) // %v可以自动推到输出类型
		default:
			fmt.Println("不是合理的数据类型")
		}
	}
}
