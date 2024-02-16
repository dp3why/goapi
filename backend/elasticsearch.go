package backend

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"goapi/constants"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

var (
	ESBackend *ElasticsearchBackend
)

type ElasticsearchBackend struct {
	Client *elasticsearch.Client
}

func InitElasticsearchBackend() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			constants.ES_AWS_URL,
		},
		Username: constants.ES_USERNAME,
		Password: constants.ES_AWS_PASSWORD,
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// Check for the existence of the post index
	res, err := es.Indices.Exists([]string{constants.POST_INDEX})
	if err != nil {
		log.Fatalf("Error checking existence of index: %s", err)
	}
	if res.StatusCode == 404 {
		createIndex(es, constants.POST_INDEX, `{
			"mappings": {
				"properties": {
					"id": { "type": "keyword" },
					"user": { "type": "keyword" },
					"message": { "type": "text" },
					"url": { "type": "keyword", "index": false },
					"type": { "type": "keyword", "index": false }
				}
			}
		}`)
	}
	res.Body.Close()

	// Check for the existence of the user index
	res, err = es.Indices.Exists([]string{constants.USER_INDEX})
	if err != nil {
		log.Fatalf("Error checking existence of index: %s", err)
	}
	if res.StatusCode == 404 {
		createIndex(es, constants.USER_INDEX, `{
			"mappings": {
				"properties": {
					"username": { "type": "keyword" },
					"password": { "type": "keyword" },
					"age": { "type": "integer", "index": false },
					"gender": { "type": "keyword", "index": false }
				}
			}
		}`)
	}
	res.Body.Close()

	fmt.Println("====== Indices created successfully ======")
	ESBackend = &ElasticsearchBackend{
		Client: es,
	}
}

func createIndex(es *elasticsearch.Client, indexName string, mapping string) {
	res, err := es.Indices.Create(
		indexName,
		es.Indices.Create.WithContext(context.Background()),
		es.Indices.Create.WithBody(bytes.NewReader([]byte(mapping))),
	)
	if err != nil || res.IsError() {
		log.Fatalf("Error creating index: %s", err)
	}
	fmt.Printf("Index %s created successfully\n", indexName)
	res.Body.Close()
}


