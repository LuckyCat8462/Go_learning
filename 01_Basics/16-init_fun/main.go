package main

import (
	_ "1-Basics/16-init函数/sub" // 此时，只会调用sub里面的init函数，编译还不会出错
	"fmt"
	// "fmt"
)

func main() {
	// res := sub.Sub(10, 5)
	// fmt.Println("sub.Sub(20,10) =", res)
	fmt.Println("Hello world")
}
