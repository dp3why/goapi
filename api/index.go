package handler

import (
	"goapi/service"
	"net/http"

	. "github.com/tbxark/g4vercel"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	server := New()

	server.GET("/", service.RootCheck)
	server.POST("/upload", service.Upload)
	server.GET("/hello", service.Hello)
	server.GET("/user/:id", func(context *Context) {
		context.JSON(400, H{
			"data": H{
				"id": context.Param("id"),
			},
		})
	})
	server.GET("/long/long/long/path/*test", func(context *Context) {
		context.JSON(200, H{
			"data": H{
				"url": context.Path,
			},
		})
	})
	server.Handle(w, r)
}