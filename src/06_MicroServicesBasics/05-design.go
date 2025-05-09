package main

// 8-2-12的视频内容，不理解的话再看看

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求, 服务端在注册rpc对象时, 能让编译期检测出 注册对象是否合法.

// MyInterface 创建接口, 在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

// RegisterService 需要理解
// 调用该方法时, 需要给 i 传参, 参数应该是 实现了 HelloWorld 方法的类对象!
func RegisterService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

// -----------------客户端用

// Myclient 向调用本地函数一样,调用远程函数.
// 定义类
type Myclient struct {
	c *rpc.Client
	// 	使用*指针就是引用传递了。而不是值传递
}

// InitClient 由于使用了 c 调用 Call, 因此需要初始化 c
func InitClient(addr string) Myclient {
	conn, _ := jsonrpc.Dial("tcp", addr)

	return Myclient{c: conn}
}

// HelloWorld 实现函数, 原型参照上面的 Interface来实现.
func (this *Myclient) HelloWorld(a string, b *string) error {

	// 参数1, 参照上面的 Interface , RegisterName 而来.  a :传入参数  b:传出参数.
	return this.c.Call("hello.HelloWorld", a, b)
}
