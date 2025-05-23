package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "src/06_MicroServicesBasics/pb06"
)

func main() {
	// 1. 连接 grpc 服务
	grpcConn, err := grpc.Dial("127.0.0.1:8800", grpc.WithInsecure())
	if err != nil {
		fmt.Println("grpc.Dial err:", err)
		return
	}
	defer grpcConn.Close()

	// 2. 初始化 grpc 客户端
	grpcClient := pb.NewSayHelloClient(grpcConn)

	// 创建并初始化Teacher 对象
	var teacher pb.Person
	teacher.Name = "itcast"
	teacher.Age = 18

	// 3. 调用远程服务。
	t, err := grpcClient.SayHello(context.TODO(), &teacher)

	fmt.Println(t, err)
}
