package service

import (
	"encoding/json"
	"goapi/model"

	. "github.com/tbxark/g4vercel"
)

 
func Upload(context *Context) {
	body := context.Req.Body
	decoder := json.NewDecoder(body)
	var p  model.Post
	if err := decoder.Decode(&p); err != nil {
		context.JSON(400, H{
			"message": "bad requrest",
		})
		return
	} 
	context.JSON(200, H{
		"message": "Post received: " + p.Message,
	})
}
