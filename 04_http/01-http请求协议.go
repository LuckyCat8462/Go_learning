package main

// 一、HTTP概念
// HTTP协议介绍
// 1.HTTP协议是应用层协议，依赖于传输层的TCP协议和网络层的IP协议。
// 2.HTTP协议主要用于Web开发，常用的编程语言包括Java、PHP和Go。C++主要用于底层开发，对HTTP协议的了解较少。

// HTTP协议的基本概念
// 1.HTTP协议是无状态的，每次请求都是独立的。
// 2.HTTPS不是标准协议，而是在HTTP基础上添加了SSL/TLS层，用于提供安全加密。HTTP是标准协议。
// 3.HTTPS通过数字证书和公钥私钥进行非对称加密，确保数据传输的安全。HTTP是明文传输
// 目前大部分网站都要求使用https协议，因为它更加安全。

// 二、HTTP请求报文格式
// HTTP请求报文包含四个部分：请求行、请求头、空行、请求体。

// 2.1.请求行
// 2.1.1格式：method URL protocol_version。（方法+URL+协议版本号）
// 2.1.2实例：POST + /chapter17/user.html HTTP/1.1
// 2.1.3请求方法：
//		1.GET:获取数据
// 		2.POST:上传数据
// 		3.PUT:修改数据
// 		4.DELETE:删除数据

// 2.2.请求头 由多个key:value对组成，包含协议自带和用户自定义字段。
// 2.2.1.格式：key:value
// 2.2.2.常见重要头
// 		1.Accept:接收数据的格式
// 		2.User-Agent:描述用户浏览器的信息
// 		3.Connection: Keep-Alive(长链接)，Close(短连接)
// 		4.Accept-Encoding:gzip,xxx,描述可以接受的编码
// 		5.Cookie:由服务器设置的key=value数据，客户端下次请求的时候可以携带过来
// 		6. Content-Type:	用于指定请求体的数据格式。
// 			1.application/-form(表示上传的数据是表单格式),
// 			2.application/json(表示body的数据是json格式)
// 		6.用户可以自定义的:
// 			1. name : neko
// 			2. age : 18

// 2.3.空行 用于分隔请求头和请求体。

// 2.4.请求包体 可选，用于携带数据，常用于POST方法上传表单或JSON数据。
// 	一般在POST方法时，会配套提供BODY ;在GET的时候也可以提供BODY，但是这样会容易混淆，建议不这样使用。
// 2.请求体数据格式可以是表单格式或JSON格式。
// 3.表单格式常用于上传表单数据。例如：姓名、电话、年龄；JSON格式常用于上传JSON数据。

// 报文案例
/*1.请求行	POST /chapter17/user.html HTTP/1.1	*/
/*2.请求头	Accept:image/jpeg,application/x-ms-application,**。。。)
Referer: http://localhost:8088/chapter17/user/register.html?
code=100&time=123123
Accept-Language:zh-CN
User-Agent:Mozilla/4.0(compatible, MSIE 8.0: Windows NT 6.1.
Content-type:application/x-www-form-urlencoded
host:localhost:8088
Content-Length:112
Connection: Keep-Alive
Cache-Control:no-cache
Cookie:JSESSIONID=24DF2688E37EE4F66D9669D2542AC17B
*/

/* 3.空行*/

/*4.请求包体	name=tom&password=l234&realName=tomson	*/
