package backend

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"goapi/constants"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
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



// ReadFromES searches documents in the specified index based on the query.
// The query parameter should be a map[string]interface{} representing the Elasticsearch query DSL.
func (backend *ElasticsearchBackend) ReadFromES(query map[string]interface{}, index string) (*esapi.Response, error) {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
		return nil, err
	}

	// Perform the search request.
	res, err := backend.Client.Search(
		backend.Client.Search.WithContext(context.Background()),
		backend.Client.Search.WithIndex(index),
		backend.Client.Search.WithBody(&buf),
		backend.Client.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
		return nil, err
	}

	return res, nil
}