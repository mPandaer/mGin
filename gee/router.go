package gee

import "strings"

type HandlerFunc func(c *Context)

const routerSep = ":"

type router struct {
	roots    map[string]*node //根据 GET POST等HTTP方法进行索引
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func parsePattern(pattern string) []string {
	vs := strings.Split(pattern, pathSep)
	parts := make([]string, 0)
	for _, item := range vs {
		if item == "" {
			continue
		}
		parts = append(parts, item)
		if string(item[0]) == "*" {
			break
		}
	}
	return parts
}

func genKey(method string, pattern string) string {
	return method + routerSep + pattern
}

func getKey(c *Context) string {
	return c.Method + routerSep + c.Path
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {

	parts := parsePattern(pattern)
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{part: method}
	}
	r.roots[method].insert(parts, 0)
	key := genKey(method, pattern)
	r.handlers[key] = handler
}

func (r *router) getRouter(method string, pattern string) (*node, map[string]string) {
	searchParts := parsePattern(pattern)
	root, ok := r.roots[method]
	params := make(map[string]string)
	if !ok {
		return nil, nil
	}
	n := root.search(searchParts, 0)

	if n == nil {
		return nil, nil
	}
	for index, part := range parsePattern(n.pattern) {
		if string(part[0]) == ":" {
			params[part[1:]] = searchParts[index]
		}
		if string(part[0]) == "*" && len(part) > 1 {
			params[part[1:]] = strings.Join(searchParts[index:], pathSep)
		}
	}
	return n, params

}

func (r *router) handle(c *Context) {
	n, params := r.getRouter(c.Method, c.Path)
	if n == nil {
		_, _ = c.Writer.Write([]byte("404 NOT FOUND\n"))
		return
	}
	c.Params = params
	key := c.Method + ":" + n.pattern
	r.handlers[key](c)
}
