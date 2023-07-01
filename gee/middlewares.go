package gee

import (
	"log"
	"time"
)

// 默认的中间件

func Logger() HandlerFunc {
	return func(c *Context) {
		t := time.Now()
		c.Next()
		log.Printf("[%d] %s in %v\n", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
