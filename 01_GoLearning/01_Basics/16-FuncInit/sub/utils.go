package sub

// package utils //不允许出现多个包名

import "fmt"

func init() {
	fmt.Println("this is init in sub utils.go")
}

func test4() {
	fmt.Println("this is test4() in sub/utils!")
}
