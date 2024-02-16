package routes

import (
	"fmt"

	. "github.com/tbxark/g4vercel"
)

func Hello(context *Context) {
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
			"url": context.Path,
		},
	})
}