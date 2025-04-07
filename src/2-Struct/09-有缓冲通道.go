package main

import (
	"fmt"
	"time"
)

func main() {
	//numsChan := make(chan int, 10)
	//有缓冲的通道，当缓冲区满时，写入操作会阻塞，当缓冲区空时，读取操作会阻塞。
	//使用make函数创建带缓冲的通道时，需要指定缓冲区的大小。如果管道没有使用make分配空间，那么管道默认是nil的，读取、写入都会阻塞
	//对于一个管道，读与写的次数，必须对等

	var names chan string //默认是nil的
	names = make(chan string, 10)

	go func() {
		fmt.Println("names:", <-names)
	}()

	names <- "hello" //由于names是nil的，写操作会阻塞在这里
	time.Sleep(1 * time.Second)

	numsChan1 := make(chan int, 10)

	//子go写
	go func() {
		for i := 0; i < 50; i++ {
			numsChan1 <- i
			fmt.Println("写入数据:", i)
		}
	}()

	//主go程：读
	go func() {
		for i := 0; i < 50; i++ {
			fmt.Println("主程序准备读取数据.....")
			data := <-numsChan1
			fmt.Println("读取数据:", data)
		}
	}()

	for {
		fmt.Println("这是主go程，正在死循环")
		time.Sleep(1 * time.Second)
	}
}
