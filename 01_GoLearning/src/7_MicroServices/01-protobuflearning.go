package main

// 1.portobuf编写的注意事项
// 		1. message 成员编号， 可以不从1开始, 但是不能重复. -- 不能使用 19000 - 19999
// 		2. 可以使用 message 嵌套。
// 		3. 定义数组、切片 使用 repeated 关键字
// 		4. 可以使用枚举 enum
// 		5. 可以使用联合体。 oneof 关键字。成员编号，不能重复。

// 2.编译 protobuf
// > 回顾：C++ 编译 命令：
// > protoc --cpp_out=./  *.proto		---> xxx.pb.cc   和  xxx.pb.h   文件
// - go 语言中 编译命令：
// 在终端输入`protoc --go_out=./ *proto`      会得到一个--->  xxx.pb.go 文件。

// 3.序列化的其他形式
// 		1.除了protobuf，还有XML、JSON等序列化形式。
// 		2.XML包含标签和键值对，JSON只包含键值对，而protobuf只包含值。
// 		3.protobuf因其简洁性而被广泛使用，特别适用于RPC服务。

// 4.添加 rpc 服务
// 语法：
// 4.1.service是关键字，用于指定服务名。
// 4.2.rpc是关键字，用于定义RPC函数。
// 4.3.函数名后括号内是参数，必须是消息体。
// 4.4.returns关键字用于指定返回值，也必须是消息体。

// 5.RPC服务的例子
// service 服务名 {
// rpc 函数名(参数：消息体) returns (返回值：消息)
// }
// message People {
// string name = 1;
// }
// message Student {
// int32 age = 2;
// }
// 例：
// service hello {
// rpc HelloWorld(People) returns (Student);
// }、
// 		1.例子中定义了一个名为hello的服务，包含一个rpc函数hello world。
// 		2.函数参数和返回值都是消息体。
// 		3.函数参数是一个名为people的消息体，包含一个string类型的name。
// 		4.函数返回值是一个名为student的消息体，包含string类型的name和int类型的age。

// 6.1补充知识点：
// - 默认，protobuf，编译期间，不编译服务。 要想使之编译。 需要使用 gRPC。
// - 使用的编译指令为：
//  - `protoc --go_out=plugins=grpc:./ *.proto`

// 客户端：
//
// type bj38Client struct {} ----- type MyClient struct {} 类
//
// func NewBj38Client()  ----- InitCient() 函数
//
// func (c *bj38Client) Say() ---- HelloWorld() 方法
//
// 服务端：
//
// type Bj38Server interface {}  ---- type MyInterface interface{} 接口。
//
// func RegisterBj38Server() ---- func RegisterService() 函数。

// 二、grpc
// 2.1.grpc官方网站介绍
// 		1.grpc有一个官方网站，提供中文版。
// 		2.官方网站地址为3w3w.grpc.io。

// 2.2.grpc介绍
// 		1.grpc是一个高性能的开源和通用的rpc框架。
// 		2.grpc面向移动设备和http/2设计。
// 		3.http/2是超文本传输协议的第二版。
// 		4.grpc支持的语言版本：grpc目前支持c、java和go语言版本。

// 2.3.grpc与protobuf的关系
// 		1.grpc基于protobuf实现。
// 		2.使用protobuf定义服务，并生成服务器和客户端代码。
// 		3.grpc使用go API实现服务。

// 2.4.grpc环境安装
// 		1.grpc需要安装才能使用，与rpc不同。
// 		2.官方推荐使用go get方法安装。
// 		3.go get命令用于联网在线获取安装包或源码包。
// 		4.如果无法直接访问google官网，可以前往github下载。
// 		5.github上的下载需要手动一个一个包地导入。
// 		6.安装过程中可以选择使用离线安装包。
