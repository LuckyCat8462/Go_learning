package main

import "fmt"

// c语言的多态需要父子继承关系
// go语言的多态不需要继承，只要实现相同的接口即可。

// 1.定义一个接口，里面设计好需要的接口，可以有多个
// 2.任何实现了这个接口的类型，都可以赋值给这个接口，从而实现多态
// 3.多个类之间不需要有继承关系
// 4.如果interface中定义了多个接口，那么实际的类必须全部实现接口函数，才可以赋值

// 实现go多态，需要实现定义接口
// 人类的武器发起攻击，不同等级子弹效果不同

// IAttack 定义一个接口, 注意类型是interface
type IAttack interface {
	// Attack 接口函数可以有多个，但是只能有函数原型，不可以有实现
	Attack()
}

// HumanLowLevel 低等级人员
type HumanLowLevel struct {
	name  string
	level int
}

func (a *HumanLowLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ", 造成1000点伤害")
}

// 高等级
type HumanHighLevel struct {
	name  string
	level int
}

func (a *HumanHighLevel) Attack() {
	fmt.Println("我是:", a.name, ",等级为:", a.level, ", 造成5000点伤害")
}

// 定义一个多态的通用接口,传入不同的对象，调用同样的方法，实现不同的效果 ==》 多态
func DoAttack(a IAttack) {
	a.Attack()
}

func main() {
	// var player interface{}
	var player IAttack // 定义一个包含Attack的接口变量

	lowLevel := HumanLowLevel{
		name:  "zhang",
		level: 1,
	}

	highLevel := HumanHighLevel{
		name:  "li",
		level: 10,
	}

	lowLevel.Attack()
	highLevel.Attack()

	// 对player赋值为lowLevel，接口需要使用指针类型来赋值
	player = &lowLevel
	player.Attack()
	// 同样一个player，赋值不同的对象，传入的是低等级和高等级时，造出的伤害不一样，多态
	player = &highLevel
	player.Attack()

	fmt.Println("多态......")
	DoAttack(&lowLevel)
	DoAttack(&highLevel)
}
