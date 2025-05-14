package main

import (
	"fmt"
	"runtime"
	"time"
)

// return  ===> 返回当前函数
// exit ===> 退出当前进程
// GOEXIT ===> 提前退出当前go程

func main() {
	go func() {
		go func() {
			func() {
				fmt.Println("这是子go程内部的函数!")
				// return //这是返回当前函数
				// os.Exit(-1) //退出进程
				runtime.Goexit() // 退出子当前go程
			}()

			fmt.Println("子go程结束!")
			fmt.Println("go 2222222222 ")

		}()
		time.Sleep(2 * time.Second)
		fmt.Println("go 111111111111111")
	}()

	fmt.Println("这是主go程!")
	time.Sleep(3 * time.Second)
	fmt.Println("OVER!")
}
