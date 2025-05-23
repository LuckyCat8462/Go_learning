package mian

// 一.new 参数：
//
// - --namespace： 命名空间  == 包名
// - --type ： 微服务类型。
//  - srv： 微服务
//  - web：基于微服务的 web 网站。

// 二.创建服务
// micro new --type srv bj38  —— 创建一个微服务项目!

// - main.go : 项目的入口文件。
// - handler/:  处理 grpc 实现的接口。对应实现接口的子类，都放置在 handler 中。
// - proto/： 预生成的 protobuf 文件。
// - Dockerfile：部署微服务使用的 Dockerfile
// - Makefile：编译文件。—— 快速编译 protobuf 文件。

// 三\查看服务
// - main.go : 项目的入口文件。
// - handler/:  处理 grpc 实现的接口。对应实现接口的子类，都放置在 handler 中。
// - proto/： 预生成的 protobuf 文件。
// - Dockerfile：部署微服务使用的 Dockerfile
// - Makefile：编译文件。—— 快速编译 protobuf 文件。

// func main() {
//	// New Service   -- 初始化服务器对象.
//	service := micro.NewService(
//		micro.Name("go.micro.srv.bj38"),   // 服务器名
//		micro.Version("latest"),		   // 版本
//	)
//	// Initialise service 与newService作用一致,但优先级高.后续代码运行期,初始化才有使用的必要.
//	//service.Init()
//
//	// Register Handler --- 注册服务
//	bj38.RegisterBj38Handler(service.Server(), new(handler.Bj38))
//
//	// Register Struct as Subscriber -- redis 发布订阅.
//	//micro.RegisterSubscriber("go.micro.srv.bj38", service.Server(), new(subscriber.Bj38))
//
//	// Register Function as Subscriber
//	//micro.RegisterSubscriber("go.micro.srv.bj38", service.Server(), subscriber.Handler)
//
//	// Run service  --- 运行服务
//	if err := service.Run(); err != nil {
//		log.Fatal(err)
//	}
// }

// 4.查看 handler/ xxx.go 文件
//
// 包含 与 Interface 严格对应的 3 个函数实现！！
