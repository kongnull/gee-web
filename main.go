package main

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", func(context *gee.Context) {
		context.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.Get("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello %s,you're at %s\n", c.Query("name"), c.Path)
	})

	r.Post("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	err := r.Run(":9900")
	if err != nil {
		return
	}
}
