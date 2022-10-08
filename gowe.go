package gowe

import (
	"net/http"
	"strings"
)

type HandleFunc func(c *Context)

type Engine struct {
	*RouterGroup
	router *router
	groups []*RouterGroup
}

func New() *Engine  {
	engine := &Engine{router: newRouter()}
	engine.RouterGroup = &RouterGroup{engine: engine}
	engine.groups = []*RouterGroup{engine.RouterGroup}
	return engine
}

func (e *Engine) addRoute(method string, pattern string, handler HandleFunc)  {
	e.router.addRoute(method,pattern,handler)
}

func (e *Engine) Get(pattern string, handler HandleFunc)  {
	e.addRoute("GET",pattern,handler)
}

func (e *Engine) Post(pattern string, handler HandleFunc)  {
	e.addRoute("POST",pattern,handler)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	var middlewares []HandleFunc
	for _, group := range e.groups {
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}

	context := NewContext(r, w)
	context.handlers = middlewares
	e.router.handle(context)
}

func (e *Engine) Run(addr string)  {
	http.ListenAndServe(addr,e)
}
