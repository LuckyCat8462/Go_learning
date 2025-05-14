package main

// 2-自增语法
import "fmt"

func main() {
	i := 20
	i++
	// 自增不允许和其他语句放在一起，必须单独一行
	// fmt.Println("i:",i++)
	fmt.Println(i)
}
