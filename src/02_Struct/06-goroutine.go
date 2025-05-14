package main

import (
	"fmt"
	"time"
)

// 并发:电脑同时听歌，看小说，看电影。cpu根据时间片进行划分，交替执行这个三个程序。在人是视角中可能看起来是同时产生的。
// 并行:多个CPU (多核)同时执行
// c语言里面实现并发过程使用的是多线程(C+ +的最小资源单元)，进程
// go语言里面不是线程，而是go程==> goroutine, go程是go语言原生支持的
// 每一个go程占用的系统资源远远小于线程，一个go程大约需要4K-5K的内存资源
// 一个程序可以启动大量的go程:
// ●线程==》几十个
// ●
// go程可以启动成百上千个，===》 对于实现高并发，性能非常好
// ●只需要在目标函数前加上go关键字即可

// 这个将用于子go程使用
func display(num int) {
	count := 1
	for {
		fmt.Println("=============> 这是子go程:", num, "当前count值:", count)
		count++
		// time.Sleep(1 * time.Second)
	}
}

func main() {
	// 启动子go程
	for i := 0; i < 3; i++ {
		go display(i)
	}

	// go func() {
	//	count := 1
	//	for {
	//		fmt.Println("=============> 这是子go程:", count)
	//		count++
	//		time.Sleep(1 * time.Second)
	//	}
	// }()

	// 主go程
	count := 1
	for {
		fmt.Println("这是主go程:", count)
		count++
		time.Sleep(1 * time.Second)
	}
}
