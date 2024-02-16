package handler

import (
	"fmt"
	"goapi/backend"
	"goapi/service"
	"net/http"

	. "github.com/tbxark/g4vercel"
)
 
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server starts...")
	// init elasticsearch backend
	backend.InitElasticsearchBackend()
	
	// create server
	server := New()
	// define route
	server.GET("/", service.RootCheck)
	server.POST("/upload", service.Upload)
	server.GET("/hello", service.Hello)
	server.GET("/user/:id", service.GetUserInfo)
	server.GET("/long/long/long/path/*test", service.GetPathTest)

	// handle request
	server.Handle(w, r)
}