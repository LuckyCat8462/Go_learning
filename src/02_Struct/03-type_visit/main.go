package main

import (
	"2-Struct/03-类的访问权限/src"
	"fmt"
)

func main() {
	// 原本 s1 := Student1
	s1 := src.Student1{
		Hum: src.Human{
			Name:   "zhi",
			Age:    18,
			Gender: "女生",
		},
		School: "猫猫一中",
	}

	fmt.Println("s1.name:", s1.Hum.Name)
	fmt.Println("s1.school:", s1.School)

	t1 := src.Teacher{}

	t1.Subject = "高数"
	t1.Name = "曾老师"
	t1.Age = 46
	t1.Gender = "男生"

	fmt.Println("t1 :", t1)
	t1.Eat()

	fmt.Println("t1.Human.name:", t1.Human.Name)
}
