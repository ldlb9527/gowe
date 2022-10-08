package gowe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Req *http.Request
	Writer http.ResponseWriter
	Path string
	Method string
	Params map[string]string
	// middleware
	handlers []HandleFunc
	index int
}

func NewContext(r *http.Request, w http.ResponseWriter) *Context  {
	return &Context{
		Req: r,
		Writer: w,
		Path: r.URL.Path,
		Method: r.Method,
		index: -1,
	}
}

func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

func (c *Context) PostForm(key string) string  {
	return c.Req.FormValue(key)
}

func (c *Context) Query(key string) string  {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int)  {
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string)  {
	c.Writer.Header().Set(key,value)
}

func (c *Context) String(code int, format string, values ...interface{})  {
	c.Writer.Header().Set("Context-type","text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,values...)))
}

func (c *Context) Json(code int, obj interface{})  {
	c.Writer.Header().Set("Context-type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		fmt.Println("json编码异常,"+err.Error())
		http.Error(c.Writer,err.Error(),500)
	}
}

func (c *Context) Html(code int, html string)  {
	c.Status(code)
	c.SetHeader("Context-type","text/html")
	c.Writer.Write([]byte(html))
}

func (c *Context) Data(code int, data []byte)  {
	c.Status(code)
	c.SetHeader("Context-type","text/html")
	c.Writer.Write(data)
}

func (c *Context) Next() {
	c.index++
	s := len(c.handlers)
	for ; c.index < s; c.index++ {
		c.handlers[c.index](c)
	}
}


