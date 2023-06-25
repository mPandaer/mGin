package main

import (
	"gee"
	"time"
)

// 测试案例
func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.JSON(200, gee.JSON{
			"master": "pandaer",
			"time":   time.Now(),
		})
	})

	engine.GET("/info", func(c *gee.Context) {
		c.HTML(200, "<h1>暂无信息</h1>\n")
	})

	engine.Run(":8080")
}
