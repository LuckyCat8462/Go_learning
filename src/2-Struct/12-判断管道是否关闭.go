package main

//需要知道一个管道的状态，如果已经关闭了，读不怕，会返回零，如果再写入的话，有崩溃的风险。
//map:==> v, ok:= m1[0]
//channel: ==> v, ok:= <- numChan

import "fmt"

func main() {
	numChan := make(chan int, 10)

	//写
	go func() {
		for i := 0; i < 10; i++ {
			numChan <- i
			fmt.Println("写入数据:", i)
		}

		close(numChan)
	}()

	//读数据
	//for v := range numChan {
	//	fmt.Println("v:", v)
	//}

	for {
		//在读之前进行判断
		v, ok := <-numChan // ok-idom模式判断
		if !ok {
			fmt.Println("管道已经关闭了，准备退出!")
			break
		}

		fmt.Println("v:", v)
	}

	fmt.Println("OVER!")
}
