package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// 1.导入 http包
	client := http.Client{}

	// 2.使用http.Client进行GET请求。访问百度作为示例。
	// func (c *Client) Get(url string) (resp *Response, err error) {
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("client.Get err:", err)
		return
	}

	// 5.响应体是一个read closer接口，需要读取才能获取数据。

	// 放在上面，内容很多
	body := resp.Body
	fmt.Println("body ---:", body)
	// // 使用ioutil这个工具读取响应体并将其转换为string。（readall就是全部读取）
	// func ReadAll(r io.Reader) ([]byte, error) {
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}

	fmt.Println("body string:", string(readBodyStr))

	// beego, gin  ==> web框架
	// 3.通过response.Header获取响应头。
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("header : ", resp.Header)

	fmt.Println("content-type:", ct)
	fmt.Println("date:", date)
	// BWS是Baidu Web Server,是百度开发的一个web服务器 大部分百度的web应用程序使用的是BWS
	fmt.Println("server:", server)

	// 4.通过response.StatusCode和response.Status获取响应状态码和描述。
	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("url:", url)       // https://www.baidu.com
	fmt.Println("code:", code)     // 200
	fmt.Println("status:", status) // OK

}
