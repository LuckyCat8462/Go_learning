package main

// Go语言网络编程主要包括两部分：socket编程（TCP/UDP）和HTTP编程。
// 1.socket编程属于传输层和网络层之间的协议。
// 2.HTTP编程是应用层协议，用于Web开发。

// socket编程分为客户端（client）和服务端（server）两部分。

// 1.客户端首先尝试连接服务端（dial），指定IP和端口（port）。
// 2.服务端监听（listen）指定IP和端口，等待客户端连接。
// 3.连接建立后（accept），客户端发送数据，服务端读取并处理数据，然后返回响应数据。
// 4.连接关闭（close）后，本次请求结束。
