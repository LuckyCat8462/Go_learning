package main

import "fmt"

// 18-类操作
// 一、概念
// 	go语言支持类的操作，但是没有class关键字

// 二、案例
// 定义一个类，类名为person，包含了name,age,gender,score四个属性。绑定方法：Eat，Run，Laugh, 成员
type person struct {
	name   string
	age    int
	gender string
	score  float64
}

// 2.1.go语言中并不在内部写，而是在类外面绑定方法
// 2.2.使用指针，修改时会将类的内容也修改

func (this *person) Eat() {
	// 使用this.参数，可以便捷调用
	// fmt.Println(this.name + "正在吃饼")
	this.name = "猫猫"
	fmt.Println("吃1个饼")
}

// 对照组，Eat2，未使用指针
// 未使用指针，修改不会讲类内容修改

func (this person) Eat2() {

	// 类的方法，可以使用自己的成员
	this.name = "猫猫"
	fmt.Println("吃2个")
}

func main() {
	neko := person{
		name:   "neko",
		age:    25,
		gender: "male",
		score:  99.99}

	neko1 := neko
	fmt.Println("Eat，使用p *Person，修改name的值 ...")
	fmt.Println("修改前neko:", neko)
	neko.Eat()
	fmt.Println("修改后neko:", neko) // 改变

	fmt.Println("Eat2，使用p Person，但是不是指针 ...")
	fmt.Println("修改前neko:", neko1)
	neko1.Eat2()
	fmt.Println("修改后neko:", neko1) // 未改变

}
