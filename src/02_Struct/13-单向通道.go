package main

//单向通道:为了明确语义，一般用于函数参数
//●单向读通道
//●单向写通道

import (
	"fmt"
	"time"
)

func main() {
	//- 单向读通道:
	//var numChanReadOnly <-chan int
	//- 单向写通道:
	//var numChanWriteOnly chan<- int

	//生产者消费者模型
	//C: 数组+锁   thread1 : 写， thread2：读
	//Go: goroutine + channel

	//1. 在主函数中创建一个双向通道 numChan
	numChan1 := make(chan int, 5)

	//2. 将numChan 传递给producer， 负责生产
	go producer(numChan1) //双向通道可以赋值给同类型的单向通道, 单向不能转双向

	//2. 将numChan 传递给consumer 负责消费
	go consumer(numChan1)

	time.Sleep(2 * time.Second)
	fmt.Println("OVER!")
}

// producer 生产者   ===> 提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		//data := <-out  //写通道不允许有读取操作
		fmt.Println("====> 向管道中写入数据:", i)
	}

	close(out)
}

// consumer 消费者   ===> 提供一个只读的通道
func consumer(in <-chan int) {
	//in <- 10  //读通道不允许有写入操作
	for v := range in {
		fmt.Println("从管道读取数据：", v)
	}

	fmt.Println("consumer end 111111!")
}
