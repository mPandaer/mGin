package gee

type HandlerFunc func(c *Context)

type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

func genKey(method string, pattern string) string {
	return method + ":" + pattern
}

func getKey(c *Context) string {
	return c.Method + ":" + c.Path
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := genKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := getKey(c)
	handler, ok := r.handlers[key]
	if !ok {
		_, _ = c.Writer.Write([]byte("404 NOT FOUND\n"))
	}
	handler(c)
}
