package api

import (
	. "github.com/tbxark/g4vercel"
)

func rootCheck(context *Context) {
	context.JSON(200, H{
		"message": "hello go from vercel !!!!",
	})
}
 