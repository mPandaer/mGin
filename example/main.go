package main

import (
	"gee"
)

// 测试案例
func main() {
	engine := gee.New()
	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>hello,gee</h1>\n")
	})

	engine.GET("/*info", func(c *gee.Context) {
		c.JSON(200, gee.JSON{
			"infoPath": c.Param("info"),
			"info":     "yyds",
		})
	})

	engine.Run(":8080")
}
