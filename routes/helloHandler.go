package routes

import (
	"fmt"

	. "github.com/tbxark/g4vercel"
)

func Hello(context *Context) {
	// get the name from the query string /hello?name=xxx
	name := context.Query("name")
	if name == "" {
		context.JSON(400, H{
			"message": "name not found",
		})
	} else {
		context.JSON(200, H{
			"data": fmt.Sprintf("Hello %s!", name),
		})
	}
}

func GetUserInfo(context *Context) {
	// get the user id from the url  /user/:id
	id := context.Param("id")
	context.JSON(200, H{
		"data": H{
			"id": id,
		},
	})
}

func GetPathTest(context *Context) {
	context.JSON(200, H{
		"data": H{
			// get the path from the url /long/long/long/path/*test
			"url": context.Path,
		},
	})
}