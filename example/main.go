package main

import (
	"fmt"
	"gee"
	"html/template"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsData(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

// 测试案例
func main() {
	engine := gee.New()
	engine.Use(gee.Recovery(), gee.Logger())
	engine.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsData,
	})
	engine.LoadHTMLGlob("templates/*")

	engine.Static("/assets", "./static")

	engine.GET("/", func(c *gee.Context) {
		c.HTML(200, "index.html", nil)
	})

	engine.GET("/panic", func(c *gee.Context) {
		panic("测试错误处理")
	})

	stu1 := &student{Name: "pandaer", Age: 19}
	stu2 := &student{Name: "bobo", Age: 20}
	engine.GET("/stu", func(c *gee.Context) {
		c.HTML(200, "info.html", gee.JSON{
			"title":  "Gin模板渲染",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	engine.GET("/date", func(c *gee.Context) {
		c.HTML(200, "date.html", gee.JSON{
			"title": "模板渲染",
			"now":   time.Now(),
		})
	})

	engine.Run(":8080")
}
