package main

import (
	"1-Basics/11-import/add"
	SUB "1-Basics/11-import/sub" // SUB是重命名的包名
	// "11-import/sub"
	// . "11-import/sub" //不推荐使用，其中所有的函数都可以直接调用，多个包同样函数名时会冲突
	"fmt"
)

// sub是文件夹名，也是包名

func main() {
	// 导入包的格式为：包名.函数
	res := SUB.Sub(25, 20)
	fmt.Println("sub的值为", res)

	res1 := add.Caps(10, 10)
	fmt.Println(res1)
	// add.add(30,20),这种形式无法调用，因为函数的首写字母是小写的
	// 要想包里的函数对外提供访问权限，那么首字母一定要大写
}
