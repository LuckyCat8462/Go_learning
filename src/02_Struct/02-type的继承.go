package main

import "fmt"

// 一、继承的概念
// 在编程语言中，继承表示一个类（子类）继承另一个类（父类）的特性和方法。

// 二、实例

type Human struct {
	// 成员属性:
	Name   string
	Age    int
	Gender string
}

func (this *Human) Eat() {
	fmt.Println("this is :", this.Name)
}

// 三、继承与嵌套的区别

// Student1 3.1.类的嵌套。
// 定义一个学生类，去嵌套一个Hum
type Student1 struct {
	Hum    Human // 包含Human类型的变量, 此时是类的嵌套
	Score  float64
	School string
}

// Teacher 3.2.类的继承
// 定义一个老师，去继承Human
type Teacher struct {
	// 3.2.1.在一个类中，直接包含另外一个类的类型Human，并且没有字段名字，默认是对它的继承
	Human
	Subject string // 学科
}

func main() {
	s1 := Student1{
		Hum: Human{
			Name:   "zhi",
			Age:    18,
			Gender: "女生",
		},
		School: "猫猫一中",
	}

	fmt.Println("s1.name:", s1.Hum.Name)
	fmt.Println("s1.school:", s1.School)

	t1 := Teacher{}
	// 3.2.2.子类可以继承父类的字段和方法，包括私有方法和属性。
	// 下面这几个字段name,age等都是继承自Human
	t1.Subject = "高数"
	t1.Name = "曾老师"
	t1.Age = 46
	t1.Gender = "男生"

	fmt.Println("t1 :", t1)
	t1.Eat()

	// 3.2.3.子类可以覆盖父类的同名字段和方法，实现方法的重写。
	// 继承的时候，虽然我们没有定义字段名字，但是会自动创建一个默认的同名字段
	// 即t1.Name可以写为t1.Human.Name
	// 这是为了在子类中依然可以操作父类（的字段），因为：子类父类可能出现同名的字段
	fmt.Println("t1.Human.name:", t1.Human.Name)
}
