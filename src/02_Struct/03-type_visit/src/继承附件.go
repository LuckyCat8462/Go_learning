package src

import "fmt"

// 一、类权限概念
// 1.1.在Go语言中，权限都是通过首字母大小来控制
// 大写字母开头的成员或方法是public的，可以在其他包中访问。
// 小写字母开头的成员或方法是private的，只能在当前包内访问。

// 1.2.import ---> 如果包名不同，那么只有大写字母开头的才是public的
// 对于类里面的成员、方法--->只有大写开头的才能在其他包中使用

type Human struct {
	// 成员属性:
	Name   string
	Age    int
	Gender string
}

func (this *Human) Eat() {
	fmt.Println("this is :", this.Name)
}

type Student1 struct {
	Hum    Human // 包含Human类型的变量, 此时是类的嵌套
	Score  float64
	School string
}

type Teacher struct {
	Human          // 在一个类中，直接包含另外一个类的类型Human，并且没有字段名字，默认是对它的继承
	Subject string // 学科
}
