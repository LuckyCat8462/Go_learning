package main

// 一、RPC概念
// 1.1.概念
// 		1.RPC（Remote Procedure Call Protocol）（远程过程调用的简写）是一种协议，用于微服务开发中的进程间通信。
// 		2.RPC的概念包括远程过程调用协议，用于规范远程过程调用的规则。
// 		3.RPC的实现需要网络支持，用于实现远程进程间的通信。

// 1.2、网络分层模型
// 		1.网络分层模型包括OSI七层模型和TCP/IP四层模型。
// 			OSI七层模型包括物理层、数据链路层、网络层、传输层、会话层、表示层和应用层。
// 			TCP/IP四层模型包括链路层、网络层、传输层和应用层。
// 		2.RPC协议位于应用层，基于TCP协议实现。

// 1.3、RPC的工作原理
// 		1.RPC允许像调用本地函数一样调用远程函数。
// 		2.通过RPC协议，可以传递函数名和参数，并在本地调用远程函数，获取返回值。
// 		3.RPC协议实现了远程进程间的通信，使得不同进程间的调用成为可能。

// 1.4、微服务中使用RPC的原因
// 		1.微服务将每个服务封装成进程，进程间独立。
// 		2.进程间或网络间的进程可以使用不同的语言实现。
// 		3.RPC协议允许遵循调用规则，实现跨语言的服务调用。

//
//
// 二、RPC使用步骤
// 2.1.回顾网络通信

// 2.2.1.Server端网络通信的主要步骤
// 		1.调用listen函数，获取listener对象，调用accept函数，获取用于通信的connect。
// 		2.数据通信：客户端先发数据，服务器端读取数据并进行处理，再发送回客户端。
// 		3.结束通信：关闭用于通信的connect和监听器。
// 		大致流程：①net.Listen()--listener	创建监听器
// 				②listener.Accpet()--conn	启动监听，建立连接
// 				③conn.read()	读写数据
// 				④conn.write()
// 				⑤defer conn.Close() /listener.Close()

// 2.2.2.Client端网络通信的主要步骤
// 		1.调用dial函数，获取用于通信的connect。
// 		2.数据通信：客户端发送数据，服务器端读取数据并进行处理，再发送回客户端。
// 		3.结束通信：关闭用于通信的connect。
// 		大致流程：①net.Dial()-- conn
// 				②conn.Write()
// 				③conn.Read()
// 				④defer conn.Close()

// 2.2.RPC使用步骤
// 2.2.1.server端使用步骤
// 		1.注册RPC服务对象，给对象绑定方法（① 定义类，②绑定类方法）
// 			rpc.RegisterName("服务名"，回调对象)
// 		2.创建监听器
// 			listener,err := net.Listen()
// 		3.建立连接。
// 			conn,err := listener.Accept()
// 		4.将连接绑定到RPC服务。
// 			rpc.serveconn(conn)

// 2.2.2.Client端使用步骤
// 		1.用rpc协议-rpc.Dial()连接服务器
//			conn,err := rpc.Dial()
// 		2.调用远程函数。
// 			conn.call("服务名.方法名"，传入参数，)
// 		2.远程函数调用：像调用本地函数一样调用远程函数。
// 		3.结束通信：关闭连接。

// 三、RPC相关函数
// 3.1.注册rpc服务
//
//	func (server *Server)RegisterName(name string, rcvr interface{}) error
//	参1:服务名。字符串类型。
//	参2:对应rpc对象。该对象绑定方法要满足如下条件:
//		1)方法必须是导出的--包外可见。首字母大写。
//		2)方法必须有两个参数，都是导出类型、內建类型
//		3)方法的第二个参数必须是“指针”(传出参数)
//		4)方法只有一个error类型的返回值
//

// 例子
// type World struct {
// }
//
// func (w *World) Hello(name string, resp *string) error {
//
// 	return
// }
// rpc.RegisterName("World", new(World))

// 3.2.绑定rpc服务
// func(server *Server) ServeConn(conn io.ReadwriteCloser
// conn:成功建立好连接的 socket-- conn

// 3.3.调用远程函数
// func (client *client) call(serviceMethod string, args interface{}, reply interface{})error
// 		serviceMethod:“服务名.方法名"
// 		args:传入参数。方法需要的数据。
// 		reply:传出参数。定义 var 变量，&变量名  完成传参。

// 四、json版rpc
// 使用 nc-127.0.0.1 880 充当服务器
// 02-client.go 充当 客户端。发起通信。-- 乱码。
// 		因为:RPC 使用了go语言特有的数据序列化 gob。其他编程语言不能解析。
// 		使用 通用的 序列化、反序列化。 --json、protobuf
// 4.1.修改客户端
//	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
// 使用 nc -|127.0.0.1 880 充当服务器
// 看到结果:
// {"method":"hello.HelloWorld","params":["李白"],"id":0}
// 4.2.修改服务器端
// 修改服务器端，使用 jsonrpc:
// 		jsonrpc.ServeConn(conn)
// 		使用 nc 127.0.0.1 880 充当客户端
// 		看到结果:
// 			echo -e '{"method":"hello.HelloWorld","params":["李白",,"id":0}'nc 127.0.0.1 8800
// 	如果绑定方法返回值的error不为空，无论传出参数是否有值，服务端都不会返回数据

// 五、rpc封装
// 5.1.服务端封装

//	1.封装接口
//		type XXX interface {
//			方法名(传入参数，传出参数) error
//		}
// 		type MyInterface interface {
// 			HelloWorld(string, *string) error
// 		}

// 2.封装注册服务方法
// 		func Registerservice(iMyInterface){
// 			rpc.RegisterName("he1lo",i)
// 		}

// 5.2.客户端封装

// 1.定义类
// 		type Myclient struct {
// 			c *rpc.client
// 		}
// 2.绑定类方法
// 		func (this *Myclient)Helloworld(a string,b *string)error
// 			return this.c.call("hello.Helloworld",a，b)
// 		}
// 3.初始客户端
// 		func Initclient(addr string)error {
// 			conn,_:= jsonrpc.Dial("tcp", adddr)
// 			return Myclientic:conn
// 		}
