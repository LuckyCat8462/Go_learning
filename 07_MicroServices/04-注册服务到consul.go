package main

// 1.以server模式启动(记得ip地址要换成自己电脑上的ip地址
// consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=192.168.6.108 -ui -rejoin -config-dir=/etc/consul.d/ -client 0.0.0.0

// 1.首先要cd进入到指定的目录下创建定义服务文件,来注册一个服务
// cd /etc/consul.d/
// 2.创建json文件,例如sudo vim web.json
// 3.按照json的格式,填写服务信息
// 服务定义文件在我们的配置目录下, /etc/consul.d/	,文件都是.json结尾
// {"service":{
// "name":"Faceid",
// "tags":["rails"],
// "port":9000}
// }
// 其中name是服务名,tags是别名,可以有多个,但是要用[]包裹,port是端口号

// 4.重新启动consul
// 5.查询服务
// 5.1.浏览器查看
// 5.2.指令查看
// curl -s 192.168.81.128:8500/v1/catalog/service/nk86

// 可以将生成的json文件复制到json.cn上在线解析成更方便观看的格式
