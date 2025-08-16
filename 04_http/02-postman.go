package main

// 首先先去网站上搜索postman并进行下载、安装、注册

// 一、新建一个request
// 在中间栏目，选择+号，即可新增一个request

// 二、测试
// 启动3-socket中的04项目，开始监听
// 前往postman，输入URL后，可以选择send将其传输到server，server监听后获取信息

// 2.1.尝试添加几个headers
// eg：添加key：name，value：neko；添加key：age，value：25
// 添加完成后选择send，前往goland查看监听到的信息
// 可以看到报文中出现了name，age字段。NAME: NEKO AGE: 25

// 2.2.添加body
// body栏目中可以选择几种形式none、form-data、raw等
// 此处选择raw(原生)形式,此时右侧下拉列表可选,此次选择JSON形式
// {
// "address":"Shanghai",
// "gender":"man"
// }

// //
// 前端与后端传输数据的方法
// 1、放在请求头中
// 2、放在请求包体内
// 3、放在URL中：使用？分割参数与URL，多个参数之间使用&分割，每一个参数都是一个键值对
