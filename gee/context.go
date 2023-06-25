package gee

import (
	"encoding/json"
	"fmt"
)

//用于封装请求和响应的信息 并封装一些用于快速返回响应的方法

//封装一些请求头信息

const (
	ContentType = "Content-Type"
)

const (
	JsonType = "application/json"
	HtmlType = "text/html"
	TextType = "text/plain"
)

type Context struct {
	Req    *Request
	Writer ResponseWriter

	//请求信息
	Path   string
	Method string
	//响应信息
	StatusCode int
}

func newContext(w ResponseWriter, r *Request) *Context {
	return &Context{
		Req:    r,
		Writer: w,
		Path:   r.URL.Path,
		Method: r.Method,
	}
}

//获取请求信息的简单方法

func (c *Context) PostForm(key string) string {
	return c.Req.PostFormValue(key) //todo(不太一样)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

//快速响应的方法

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) Header(key string, value string) {
	c.Writer.Header().Set(key, value)
}

type JSON map[string]interface{} //JSON类型表达

func (c *Context) JSON(code int, data JSON) {
	c.Status(code)
	//c.Writer.Header().Set(CONTENT_TYPE, JSON_TYPE)
	c.Header(ContentType, JsonType)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		panic("JSON解析失败")
	}

}

func (c *Context) HTML(code int, data string) {
	//c.Writer.Header().Set(CONTENT_TYPE, HTML_TYPE)
	c.Header(ContentType, HtmlType)
	c.Status(code)
	_, _ = c.Writer.Write([]byte(data))
}

func (c *Context) String(code int, format string, values ...interface{}) {
	c.Status(code)
	//c.Writer.Header().Set(CONTENT_TYPE, TEXT_TYPE)
	c.Header(ContentType, TextType)
	_, _ = fmt.Fprintf(c.Writer, format, values)

}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, _ = c.Writer.Write(data)
}
