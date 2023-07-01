package main

import (
	"gee"
	"log"
	"time"
)

//自定义中间件

func OnlyForApi() gee.HandlerFunc {
	return func(c *gee.Context) {
		t := time.Now()
		c.Fail(500, "Internal Server Error\n")
		log.Printf("[%d] %s in %v for api group\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

// 测试案例
func main() {
	engine := gee.New()
	engine.Use(gee.Logger())
	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "<h1>hello,gee</h1>\n")
	})

	apiGroup := engine.Group("/api")
	apiGroup.Use(OnlyForApi())

	{
		apiGroup.GET("/user", func(c *gee.Context) {
			c.JSON(200, gee.JSON{
				"name":   "pandaer",
				"age":    19,
				"height": 1.78,
			})
		})

		apiGroup.GET("/dream", func(c *gee.Context) {
			c.JSON(200, gee.JSON{
				"dream1": "doctor",
				"dream2": "技术大佬",
			})
		})
	}

	pageGroup := engine.Group("/static")

	{
		pageGroup.GET("/good", func(c *gee.Context) {
			c.HTML(200, "<h1>Good Mood!</h1>")
		})

		pageGroup.GET("/baidu", func(c *gee.Context) {
			c.HTML(200, "<a href='http://www.baidu.com'>点击访问百度</a>")
		})
	}

	engine.Run(":8080")
}
