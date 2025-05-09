package main

// 该demo只能接受一次连接，且只能发送一次数据。仅用于理解思路

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 一、创建监听
	ip := "127.0.0.1"
	port := 8848 // 8848为默认端口
	address := fmt.Sprintf("%s:%d", ip, port)
	// net.Listen("tcp", ":8848") //简写，冒号前面默认是本机: 127.0.0.1

	// fmt.Sprintf 格式化字符串
	// 格式化样式：字符串形式，格式化符号以 % 开头， %s 字符串格式，%d 十进制的整数格式。
	// 参数列表：多个参数以逗号分隔，个数必须与格式化样式中的个数一一对应，否则运行时会报错。

	// func Listen(network, address string) (Listener, error) {

	listener, err := net.Listen("tcp", address)

	if err != nil {
		fmt.Println("net.Listen err:", err)
		return
	}

	fmt.Println("监听中...")

	// 二、接受连接
	// 2.1.Accept() (Conn, error)
	// 1.accept函数用于接受连接，返回一个connection和一个error。
	// 2.accept函数会阻塞，直到有连接建立。
	// 3.连接建立后，可以进行读写操作。
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
		return
	}

	fmt.Println("连接建立成功!")

	// 2.2.创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) // 使用make来创建字节切片, byte ==> uint8

	// 2.3.read数据
	// 1.read函数用于从连接中读取数据，需要一个字节切片作为缓冲区。
	// 2.read函数返回实际读取的字节数和error。
	// 3.读取到的数据存储在缓冲区中，可以根据实际需求处理数据。

	// Read(b []byte) (n int, err error)
	// cnt：真正读取client发来的数据的长度
	cnt, err := conn.Read(buf) // ctrl+左键，查看函数原型
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}

	fmt.Println("Client =====> Server, 长度:", cnt, "，数据:", string(buf[0:cnt]))

	// 服务器对客户端请求处理并进行响应 ,例如将数据转成大写 "hello" ==> HELLO
	// func ToUpper(s string) string {
	upperData := strings.ToUpper(string(buf[0:cnt]))

	// 通过connection建立的连接，所以要发送回客户端也需要找到这个cnt，使用connection的Write方法将响应数据发送回去。
	// Write(b []byte) (n int, err error)
	cnt, err = conn.Write([]byte(upperData))
	fmt.Println("Client  <====== Server, 长度:", cnt, "，数据:", upperData)

	// 关闭连接
	conn.Close()
}
