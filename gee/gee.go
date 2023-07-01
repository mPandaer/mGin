package gee

import (
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

type (
	ResponseWriter = http.ResponseWriter
	Request        = http.Request
)

// Engine 作为服务器的入口 并且作为RouterGroup的顶层分组,换句话说Engine自己也是一个RouterGroup
type Engine struct {
	*RouterGroup
	r      *router
	groups []*RouterGroup //保存所有的路由分组
}

//增加分组逻辑 目的给一类URL进行相同的操作

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc
	parent      *RouterGroup
	engine      *Engine
}

func New() *Engine {
	engine := &Engine{r: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup} //将自己也保存进去
	return engine
}

// 创建一个新的分组

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
		parent: group,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s\n", method, pattern)
	group.engine.r.addRoute(method, pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute(GET, pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute(POST, pattern, handler)
}

// ServeHTTP(ResponseWriter, *Request)
func (e *Engine) ServeHTTP(w ResponseWriter, r *Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	context := newContext(w, r)
	context.handlers = middlewares
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
