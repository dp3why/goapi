package constants

import (
	"os"

	"github.com/alexsasharegan/dotenv"
)
 
var _ = dotenv.Load()

var ES_AWS_PASSWORD string = os.Getenv("ES_AWS_PASSWORD")
var ES_AWS_URL string = os.Getenv("ES_AWS_URL")
var ES_USERNAME string= os.Getenv("ES_USERNAME")

const (
	USER_INDEX = "user"
	POST_INDEX = "post"

)
