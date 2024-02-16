package handler

import (
	"fmt"
	"goapi/backend"
	"goapi/routes"
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
	server.GET("/", routes.RootCheck)
	server.POST("/upload", routes.Upload)
	server.GET("/hello", routes.Hello)
	server.GET("/user/:id", routes.GetUserInfo)
	server.GET("/long/long/long/path/*test", routes.GetPathTest)

	// handle request
	server.Handle(w, r)
}