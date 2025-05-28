package main

// - 路由器：资源分发
// - 路由：请求分发。
//  - service.HandleFunc("/test77/call", handler.Test77Call)
//    - 将 /test77/call 这个请求，通过 回到函数 Test77Call() 处理。
// - URL：
//  - 组成：https://ip+port/资源路径
//    - https://ip+port/   找到 pc机，找到 对应进程
//    - 资源路径：在代码中，称之为路由。
//  - “/ ” : 代表 主机上进程 对应的默认资源。
//    - http协议，自动找当前目录下的 index.html 文件，做默认页面。
