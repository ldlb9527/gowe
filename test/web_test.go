package test

import (
	"fmt"
	"gowe"
	"testing"
)

func TestWeb(t *testing.T) {
	engine := gowe.New()
	group := engine.Group("/hello")
	group.Get("/:name", func(c *gowe.Context) {
		fmt.Println("/hello/"+c.Params["name"])
	})
	group.Get("/b/c", func(c *gowe.Context) {
		fmt.Println("/hello/b/c")
	})

	engine.Get("/hi/:name", func(c *gowe.Context) {
		fmt.Println("/hi/"+c.Params["name"])
	})

	engine.Post("/static/*path", func(c *gowe.Context) {
		fmt.Println(c.Params["path"])
	})

	engine.Run(":8081")

}
