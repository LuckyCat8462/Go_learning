package main

// consul常用指令
// - consul agent
//    - -bind=0.0.0.0 		指定 consul所在机器的 IP地址。 默认值：0.0.0.0
//    - -http-port=8500    consul 自带一个web访问的默认端口：8500
//    - -client=127.0.0.1   表明哪些机器可以访问consul 。 默认本机。0.0.0.0 所有机器均可访问。
//    - -config-dir=foo      所有主动注册服务的 描述信息
//    - -data-dir=path       储存所有注册过来的srv机器的详细信息。
//    - -dev                         开发者模式，直接以默认配置启动 consul
//    - -node=hostname  服务发现的名字。
//    - -rejoin                     consul 启动的时候，加入到的 consul集群
//    - -server                    以服务方式开启consul， 允许其他的consul 连接到开启的 consul上 （形成集群）。如果不加 -server， 表示以 “客户端” 的方式开启。不能被连接。
//    - -ui 		                  可以使用 web 页面 来查看服务发现的详情

// 小案例
// 以server模式启动一个consul
// 记得ip地址要换成自己的ip地址
// consul agent -server -bootstrap-expect 1 -data-dir /tmp/consul -node=n1 -bind=192.168.6.108 -ui -rejoin -config-dir=/etc/consul.d/ -client 0.0.0.0
// 可以看到提示consul agent running
// 进入浏览器,输入192.168.6.108:8500,可以看到consul实例的页面

// consul members	查看集群中有多少成员
// consul info	查看当前consul的ip信息
// consul leave	优雅的关闭consul
