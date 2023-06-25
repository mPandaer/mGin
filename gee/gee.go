package gee

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

//框架内部核心代码

const (
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"
	PUT    = "PUT"
)

// HandlerFunc 定义一个请求处理器
type HandlerFunc http.HandlerFunc
type ResponseWriter = http.ResponseWriter
type Request = http.Request

// Engine 作为服务器的入口
type Engine struct {
	router map[string]HandlerFunc
}

func genKey(r *Request) string {
	method := r.Method
	method = strings.ToUpper(method)
	return method + ":" + r.URL.Path
}

// ServeHTTP(ResponseWriter, *Request)
func (e *Engine) ServeHTTP(w ResponseWriter, r *Request) {
	handler, ok := e.router[genKey(r)]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte("404 NOT FOUND\n"))
		if err != nil {
			log.Println("err occurs!")
		}
		return
	}
	handler(w, r)
}

func (e *Engine) AddRouter(method string, pattern string, handler HandlerFunc) {
	key := method + ":" + pattern
	e.router[key] = handler
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
		router: make(map[string]HandlerFunc),
	}
}

type T int

func TT(arg T) {
	fmt.Println(arg)
}
