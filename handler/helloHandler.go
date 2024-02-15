package handler

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