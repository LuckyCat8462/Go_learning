package main

import (
	"fmt"
	"net"
	"strings"
)

// 单连接服务器的问题
// 1.单连接服务器只能接收一个链接，无法接收多个链接。
// 2.数据只能接收一次，读一次后结束。

// 多链接服务器SocketServer
// 1.主体逻辑没问题，但需要进行处理以支持多个连接。
// 2.希望服务器能够接收多个连接，并处理多次数据。
// 3.类似于腾讯,lol等在线游戏服务器，能够支持大量长连接。

// 多连接服务器的实现
// 1.实现多个连接需要主程序监听，多进程处理。
// 2.主进程负责监听，子进程负责数据处理。
// 3.通过for循环在主进程中建立多个连接。

func main() {
	// 创建监听
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)

	// func Listen(network, address string) (Listener, error) {
	// net.Listen("tcp", ":8848") //简写，冒号前面默认是本机: 127.0.0.1
	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}
	fmt.Println("server start ...")

	// 多连接服务器的实现
	// 1.实现多个连接需要主程序监听，多进程处理。
	// 2.主进程负责监听，子进程负责数据处理。
	// 3.通过for循环在主进程中建立多个连接。

	// 需求：
	// server可以接收多个连接， ====> 主go程负责监听，子go程负责数据处理
	// 每个连接可以接收处理多轮数据请求
	// 主go程只负责建立连接，然后调用子go程，也可以将其写为一个函数
	for {
		fmt.Println("监听中...")

		// Accept() (Conn, error)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		fmt.Println("连接建立成功!")
		// 接受多个连接
		go handleFunc(conn)
		// func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
		//    DefaultServeMux.HandleFunc(pattern, handler)
		// }
		// 	handleFunc()这个函数就是注册路由，HandleFunc的第一个参数是路由表达式，也就是请求路径，
		// 	第二个参数是一个函数类型，也就是真正处理请求的函数。没有其他逻辑，直接调用DefaultServeMux.HandleFunc()处理。
	}
}

// 处理函数和数据结构
// 1.将业务处理代码写成单独的函数或构造函数。
// 2.使用connection接口传递连接信息。
// 3.connection是独立的，不同的连接有不同的connection。
// 处理具体业务的逻辑，需要将conn传递进来，每一新连接，conn是不同的
func handleFunc(conn net.Conn) {
	for { // 这个for循环，保证每一个连接可以多次接收处理客户端请求
		// 创建一个容器，用于接收读取到的数据
		buf := make([]byte, 1024) // 使用make来创建字节切片, byte ==> uint8

		fmt.Println("准备读取客户端发送的数据....")

		// Read(b []byte) (n int, err error)
		// cnt：真正读取client发来的数据的长度
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Client =====> Server, 长度:", cnt, "，数据:", string(buf[0:cnt]))

		// 服务器对客户端请求进行响应 ,将数据转成大写 "hello" ==> HELLO
		// func ToUpper(s string) string {
		upperData := strings.ToUpper(string(buf[0:cnt]))

		// Write(b []byte) (n int, err error)
		cnt, err = conn.Write([]byte(upperData))
		fmt.Println("Client  <====== Server, 长度:", cnt, "，数据:", upperData)
	}

	// 关闭连接
	_ = conn.Close()
	// 	_下划线代表忽略它的值
}
