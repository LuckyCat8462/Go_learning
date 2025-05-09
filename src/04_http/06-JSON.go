package main

// JSON介绍
// 1.JSON（JavaScript Object Notation）是一种轻量级的数据交换格式，广泛应用于网络传输和数据处理。
// 2.JSON数据由键值对组成，键是字符串，值可以是字符串、数字、数组、对象等。
// 3.JSON语法要求最后一个元素后面不能加逗号。
import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	gender string // 注意，gender是小写的, 小写字母开头的，在json编码时会忽略掉
}

func main() {
	// JSON编解码概念
	// 1.编码：将结构体转换为JSON字符串，以便在网络中传输。传输  ===》 结构体 ==》 字符串  ==》 编码
	// 2.解码：将JSON字符串转换为结构体，以便进行进一步操作。操作 ==》 字符串 ==》 结构体  ==》解密
	// 3.编码和解码是可逆的过程，用于在结构和JSON字符串之间进行转换。

	neko := Student{
		Id:     1,
		Name:   "neko",
		Age:    25,
		gender: "man",
	}

	// Go语言中的json包
	// 1.Go语言中有一个json包，提供了marshal和unmarshal两个方法，用于编码和解码JSON数据。
	// 2.marshal方法将结构体编码为JSON字符串，unmarshal方法将JSON字符串解码为结构体。
	// 编码（序列化）,结构=》字符串
	// func Marshal(v interface{}) ([]byte, error) {
	encodeInfo, err := json.Marshal(&neko)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}

	fmt.Println("encodeInfo:", string(encodeInfo))

	// 对端接收到数据
	// 反序列化（解码）： 字符串=》结构体

	var zhi Student

	// func Unmarshal(data []byte, v interface{}) error {
	if err := json.Unmarshal([]byte(encodeInfo), &zhi); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}

	fmt.Println("name:", zhi.Name)
	fmt.Println("gender:", zhi.gender)
	fmt.Println("age:", zhi.Age)
	fmt.Println("id:", zhi.Id)

}
