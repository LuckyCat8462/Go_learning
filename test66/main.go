package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//   1 初始化路由	官网为初始化web引擎
	route := gin.Default()
	//   2 路由匹配
	// 常用4种
	// GET	查数据
	// POST	增加数据
	// DELETE	删除数据
	// PUT	修改数据
	route.GET("/", func(c *gin.Context) {})
	//   3 启动运行
}

// TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
