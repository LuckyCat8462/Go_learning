package main

import (
	"encoding/json"
	"fmt"
)

// 结构体标签tag介绍
// 1.结构体标签用于为结构体字段添加标识，主要用于JSON编码。
// 2.标签可以引导编码过程中的字段命名和类型转换。

type Teacher struct {
	Name    string `json:"-"`                 // ==> 1.忽略字段：使用反引号`忽略某个字段，使其在JSON编码时不被包含。
	Subject string `json:"Subject_name"`      // ==> 2.定义别名：使用json:定义别名，改变字段在JSON中的名称。在json编码时，这个字段会编码程Subject_name
	Age     int    `json:"age,string"`        // ==> 3.类型转换：使用json:定义别名时，可以指定新的数据类型，如将整数转换为字符串。在json编码时，将age转成程string类型, 一定要两个字段:名字,类型，中间不能有空格
	Address string `json:"address,omitempty"` // ==》4.忽略空字段：使用json:omit empty标签，忽略空字段的编码。

	// 注意，gender是小写的, 小写字母开头的，在json编码时会忽略掉
	gender string
}

type Master struct {
	Name    string
	Subject string
	Age     int
	Address string
	gender  string
}

func main() {

	t1 := Teacher{
		Name:    "neko",
		Subject: "Golang",
		Age:     66,
		gender:  "Man",
		Address: "上海",
	}

	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal(&t1)

	fmt.Println("encodeInfo:", string(encodeInfo))

	// 解码
	t2 := Teacher{}
	_ = json.Unmarshal(encodeInfo, &t2)
	fmt.Println("t2:", t2.Subject)

	m1 := Master{}
	_ = json.Unmarshal(encodeInfo, &m1)
	fmt.Println("m1:", m1)

	// 	结构体标签总结
	// 1.结构体标签在JSON编码时非常有用，可以控制字段的命名和类型转换。
	// 2.标签的使用包括忽略字段、定义别名、类型转换和忽略空字段。

}
