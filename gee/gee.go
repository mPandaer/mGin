package gee

import (
	"log"
	"net/http"
)

//框架内部核心代码

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

// HandlerFunc 定义一个请求处理器

type ResponseWriter = http.ResponseWriter
type Request = http.Request

// Engine 作为服务器的入口
type Engine struct {
	r *router
	//context *Context
}

//func getKey(r *Request) string {
//	method := r.Method
//	method = strings.ToUpper(method)
//	return method + ":" + r.URL.Path
//}

// ServeHTTP(ResponseWriter, *Request)
func (e *Engine) ServeHTTP(w ResponseWriter, r *Request) {
	context := newContext(w, r)
	e.r.handle(context)
}

func (e *Engine) AddRouter(method string, pattern string, handler HandlerFunc) {
	e.r.addRoute(method, pattern, handler)
}

func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.AddRouter(GET, pattern, handler)
}

func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.AddRouter(POST, pattern, handler)
}

func (e *Engine) Run(addr string) {
	err := http.ListenAndServe(addr, e)
	if err != nil {
		log.Println("ListenAndServe: ", err)
		return
	}
}

func New() *Engine {
	return &Engine{
		r: newRouter(),
	}
}
