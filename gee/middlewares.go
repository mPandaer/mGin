package gee

import (
	"fmt"
	"log"
	"time"
)

// 默认的中间件

// Logger 日志中间件
func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

// Recovery 错误处理的中间件
func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				message := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", trace(message))
				c.Fail(500, "服务器内部错误!")
			}
		}()
		c.Next()
	}
}
