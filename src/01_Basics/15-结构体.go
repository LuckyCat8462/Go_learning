package main

import "fmt"

// 一、概念
// Go 通过类型别名（alias types）和结构体的形式支持用户自定义类型。
// 结构体是复合类型，当需要定义类型，它由一系列属性组成，每个属性都有自己的类型和值的时候，就应该使用结构体，它把数据聚集在一起。

// c语言里面，我们可以使用typedef  int MyInt
type MyInt int // type相当于typdef

// C:
// struct Person {
//
// }

// go语言结构体使用type + struct来处理
type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

func main() {
	var i, j MyInt
	i, j = 10, 20

	fmt.Println("i+j:", i+j)

	// 创建变量，并赋值
	neko := Student{
		name:   "neko",
		age:    25,
		gender: "男生",
		score:  80} // 最后一个元素后面必须加上逗号，如果不加逗号则必须与}同一行

	// 使用结构体各个字段
	fmt.Println("neko:", neko.name, neko.age, neko.gender, neko.score)

	// 结构体没有-> 操作
	s1 := &neko
	fmt.Println("neko 使用指针s1.name打印:", s1.name, s1.age, s1.gender, s1.score)
	fmt.Println("neko 使用指针(*s1).name打印:", (*s1).name, s1.age, s1.gender, s1.score)

	// 在定义期间对结构体赋值时，如果每个字段都赋值了，那么字段的名字可以省略不写
	// 如果只对局部变量赋值，那么必须明确指定变量名字
	zhi := Student{
		name: "zhi",
		age:  20,
		// "女生",
		// 99,
	}
	zhi.gender = "女生"
	zhi.score = 100

	fmt.Println("zhi:", zhi)
}
