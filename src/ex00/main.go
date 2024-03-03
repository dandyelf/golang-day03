package main

//curl -XDELETE "http://localhost:9200/places"

import (
	// "fmt"
	// "context"
	_ "embed"
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

const indexName = "places"

//go:embed schema.json
var jsonSchema string

func main() {
	esClient, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// Check connection
	code, err := esClient.Ping()
	if err != nil {
		log.Fatalf("Error pinging the Elasticsearch server: %s", err)
	}
	log.Println("Elasticsearch returned with code___", code.Status())
	// Create index
	res, err := esClient.Indices.Create(indexName, esClient.Indices.Create.WithBody(strings.NewReader(jsonSchema)))

	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Response error: Cannot create index: %s", res)
	}
	res.Body.Close()

	query := `{ "query": { "match_all": {} } }`
	result, err := esClient.Search(
		esClient.Search.WithIndex("places"),
		esClient.Search.WithBody(strings.NewReader(query)),
	)
	fmt.Println("Search executed successfully", err, result.IsError(), result.String())
	esClient.Indices.Delete([]string{"places"})
}
