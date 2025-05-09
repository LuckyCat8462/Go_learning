package main

import (
	"fmt"
	"net"
	"time"
)

// socket客户端建立连接
// 1.客户端步骤少于服务器端，首先指定服务器地址和端口。
// 2.使用TCP协议，并指定IP地址和端口号。
// 3.建立连接后，如果出错则返回错误信息。
func main() {
	conn, err := net.Dial("tcp", ":8848")
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return
	}

	fmt.Println("cliet与server连接建立成功!")

	// 客户端发送数据
	// 1.连接建立成功后，客户端可以向服务器发送数据。
	// 2.发送数据时，指定要发送的数据类型和大小。
	// 3.发送数据后，检查发送是否成功，如果出错则退出。
	sendData := []byte("helloworld")

	// 加入for循环，让cilent能够多次发送数据
	for { // 保证多次向server发送数据
		// 向服务器发送数据
		cnt, err := conn.Write(sendData)
		if err != nil {
			fmt.Println("conn.Write err:", err)
			return
		}

		fmt.Println("Client ===> Server cnt:", cnt, ", data:", string(sendData))

		// 客户端接收数据
		// 1.接收服务器返回的数据，需要创建一个缓冲区用于存储数据。
		// 2.读取数据时，检查读取的字节数是否正确。
		// 3.如果读取的字节数不正确，则返回错误信息。
		// 4.读取到数据后，打印出接收到的数据。
		// 接收服务器返回的数据
		// 创建buf，用于接收服务器返回的数据
		buf := make([]byte, 1024)
		cnt, err = conn.Read(buf)
		if err != nil {
			fmt.Println("conn.Read err:", err)
			return
		}

		fmt.Println("Client <==== Server , cnt:", cnt, ", data:", string(buf[0:cnt]))
		time.Sleep(1 * time.Second)
	}
	// 关闭连接
	conn.Close()
}
