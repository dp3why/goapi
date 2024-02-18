package constants

import (
	"os"

	"github.com/alexsasharegan/dotenv"
)
 
var _ = dotenv.Load()
 
var ES_AWS_PASSWORD string = os.Getenv("ES_AWS_PASSWORD")
var ES_AWS_URL string = os.Getenv("ES_AWS_URL")
var ES_USERNAME string= os.Getenv("ES_USERNAME")
var GCS_BUCKET string = os.Getenv("GCS_BUCKET")
var PRIVATE_KEY string = os.Getenv("GCS_PRIVATE_KEY")
var CLIENT_EMAIL string = os.Getenv("GCS_CLIENT_EMAIL")

const (
	USER_INDEX = "user"
	POST_INDEX = "post"
)
