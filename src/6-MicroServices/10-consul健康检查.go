package main

// 健康检查是服务发现的关键组件.预防使用到不健康的服务.
// 和服务注册类似,一个检查可以通过检查定义或HTTP API请求来注册,我们将使用和检查定义来注册检查,和服务类似,
// 因为这是建立检査最常用的方式.

// 检查步骤
// 1.sudo vim /etc/consul.d/web.json 打开配置文件
// 2.写入 服务的配置 信息。
// 3.执行命令，consul reload。或者，关闭consul 再重启
// 4. 使用 浏览器 键入 192.168.6.108:8500 查看“bi38”这个服务 的 健康状况。
// 不健康!没有服务bi38 给 consu 实时回复!
// 5.除了 http 实现健康检査外，还可以使用“脚本”、“tcp”、“tt!”方式进行健康检查。

// {"service":{
//	"name": "web2",
//	"tags":["neko01","zhi01"],
//	"address":"192.168.81.128",
//	"port": 8800,
//	"check":{
//		"id": "api",
//		"name": "HTTP API on port 8800",
//		"http": "http://192.168.81.128:8800",	#此处的192...这部分是localhost本地ip
//		"interval":"10s",
//		"timeout":"ls"
//	}
// }
// }
// 健康检查需要完善服务配置信息，包括创建web.json文件。
// 重新加载配置文件
// 1.重新加载配置文件可以使用console reload命令
// 2.reload命令需要输入正确的参数，如serviceid或address。
// 3.确保服务已启动后再执行reload命令。
