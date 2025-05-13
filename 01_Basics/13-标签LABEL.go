package main

import "fmt"

func main() {
	// 一、概念
	// 标签是Go语言中用于标记代码位置的标识符，通常与continue、break或goto结合使用，实现跨层控制流。
	// 标签 LABEL1
	// goto LABEL  ===> 下次进入循环时，i不会保存之前的状态，重新从0开始计算，重新来过
	// continue LABEL1 ===> continue会跳到指定的位置，但是会记录之前的状态，i变成1
	// break LABEL1  ==> 直接跳出指定位置的循环

	// 二、注意点
	// 2.1.label在goto中是必选的，continue,break中可选
	// 2.2.label作用范围:定义label的函数体内
	// 2.3.label可以定义在函数的任何位置,无论是调用点的前面还是后面

	// 三、实例
	// 标签的名字是自定义的，后面加上冒号
LABEL121:

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				// goto LABEL1
				// continue LABEL1
				break LABEL121
				// break
			}

			fmt.Println("i:", i, ",j:", j)
		}
	}

	fmt.Println("over!")
}
