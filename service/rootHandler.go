package service

import (
	. "github.com/tbxark/g4vercel"
)

func RootCheck(context *Context) {
	context.JSON(200, H{
		"message": "hello go from vercel !!!!",
	})
}
 