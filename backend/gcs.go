package backend

import (
	"context"
	"fmt"
	"io"
	"log"

	"goapi/constants"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
	"google.golang.org/api/option"
)

var (
	GCSBackend *GoogleCloudStorageBackend
)

type GoogleCloudStorageBackend struct {
	client *storage.Client
	bucket string
}

func InitGCSBackend() {
	ctx := context.Background()

    config := &jwt.Config{
        Email:      constants.CLIENT_EMAIL,
        PrivateKey: []byte(constants.PRIVATE_KEY),
        Scopes: []string{
            storage.ScopeFullControl,
        },
        TokenURL: google.JWTTokenURL,
    }

	// Create an HTTP client using the JWT config
	httpClient := config.Client(ctx)
	
	client, err := storage.NewClient(ctx,  option.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	GCSBackend = &GoogleCloudStorageBackend{
		client: client,
		bucket: constants.GCS_BUCKET,
	}
}

func (backend *GoogleCloudStorageBackend) SaveTOGCS(r io.Reader, objectName string) (string, error) {
	ctx := context.Background()
	object := backend.client.Bucket(backend.bucket).Object(objectName)
	fmt.Println("===== saving to GCS ==============")
	wc := object.NewWriter(ctx)
	if _, err := io.Copy(wc, r); err != nil {
		return "", fmt.Errorf("io.Copy: %w", err)
	}
	if err := wc.Close(); err != nil {
		return "", fmt.Errorf("Writer.Close: %w", err)
	}
	if err := object.ACL().Set(ctx, storage.AllUsers, storage.RoleReader); err != nil {
		return "", fmt.Errorf("ACL().Set: %w", err)
	}
	
    fmt.Println("======= Get link============")

	attrs, err := object.Attrs(ctx)
	if err != nil {
		return "", fmt.Errorf("object.Attrs: %w", err)
	}
	fmt.Printf("File %v saved to GCS: %s.\n",  objectName, attrs.MediaLink)
	return attrs.MediaLink, nil
}
 