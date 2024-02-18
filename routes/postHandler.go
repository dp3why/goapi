package routes

import (
	"encoding/json"
	"goapi/model"
	"goapi/service"
	"log"
	"path/filepath"

	"github.com/google/uuid"
	. "github.com/tbxark/g4vercel"
)

var (
    mediaTypes = map[string]string{
        ".jpeg": "image",
        ".jpg":  "image",
        ".gif":  "image",
        ".png":  "image",
        ".mov":  "video",
        ".mp4":  "video",
        ".avi":  "video",
        ".flv":  "video",
        ".wmv":  "video",
    }
)
 
// 1. upload
func Upload(context *Context) {

	// parse multipart form from body
	err := context.Req.ParseMultipartForm(32 << 20)
	if err != nil {
		context.JSON(400, H{
			"message": "Failed to parse multipart form",
		})
		return
	}
	// Example of retrieving a form value
	message := context.Req.FormValue("message")
	username := context.Req.FormValue("username")
	
	p := model.Post{
        Id: uuid.New().String(),
        User:   username,
        Message: message, 		
	}

	//  Retrieving file
	file, header, err := context.Req.FormFile("media_file")
	if err != nil {
		context.JSON(400, H{
			"message": "Failed to retrieve media file",
		})	
		return
	}
	defer file.Close()
	suffix := filepath.Ext(header.Filename)
    if t, ok := mediaTypes[suffix]; ok {
        p.Type = t
    } else {
        p.Type = "unknown"
    }



	/*
	========= Parse JSON format body ==========

	body := context.Req.Body
	decoder := json.NewDecoder(body)
	var p  model.Post
	if err := decoder.Decode(&p); err != nil {
		context.JSON(400, H{
			"message": "bad requrest",
		})
		return
	} 
	===========================================
	*/

	err = service.SavePost(p, file)
    if err != nil {
        context.JSON(400, H{
			"message": "Failed to save post to GCS or Elasticsearch",
		})

        log.Fatalf("Failed to save post to GCS or Elasticsearch %v\n", err)
        return
    }

	context.JSON(200, H{
		"message": "Post created: " + p.Message,
	})
}


// 2. search
func Search(context *Context) {
	user := context.Query("name")
	keywords := context.Query("keywords")

 	var posts []model.Post
	var err error
	if user != "" {
        posts, err = service.SearchPostsByUser(user)
    } else {
        posts, err = service.SearchPostsByKeywords(keywords)
    }
	if err != nil {
        context.JSON(404, H{
			"message": "Not found",
		})
        log.Fatalf("Failed to read post from backend %v.\n", err)
        return
    }
	js, err := json.Marshal(posts)

	if err != nil {
        context.JSON(400, H{
			"data": "Failed to parse posts into JSON format",
		})
        log.Fatalf("Failed to parse posts into JSON format %v.\n", err)
        return
    }
	context.JSON(200, H{
		"data": string(js),
	})
}


