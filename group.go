package gowe

import "log"

type RouterGroup struct {
	prefix      string
	middlewares []HandleFunc // support middleware
	engine      *Engine      // all groups share a Engine instance
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	engine := group.engine
	newGroup := &RouterGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouterGroup) addRoute(method string, comp string, handler HandleFunc) {
	pattern := group.prefix + comp
	log.Printf("Route %4s - %s", method, pattern)
	group.engine.router.addRoute(method, pattern, handler)
}

// Get defines the method to add GET request
func (group *RouterGroup) Get(pattern string, handler HandleFunc) {
	group.addRoute("GET", pattern, handler)
}

// Post  defines the method to add POST request
func (group *RouterGroup) Post(pattern string, handler HandleFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouterGroup) Use(middlewares ...HandleFunc) {
	group.middlewares = append(group.middlewares, middlewares...)
}
