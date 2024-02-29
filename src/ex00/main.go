package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	err, esCfg := userInterface()
	if err != nil {
		fmt.Println("Crit error")
		return
	}
	esClient, err := elasticsearch.NewClient(esCfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	// Check connection
	code, err := esClient.Ping()
	if err != nil {
		log.Fatalf("Error pinging the Elasticsearch server: %s", err)
	}
	log.Println("Elasticsearch returned with code___", code.Status())
	log.Println(elasticsearch.Version)
	log.Println(esClient.Info())

	query := `{ "query": { "match_all": {} } }`
	result, err := esClient.Search(
		esClient.Search.WithIndex("my_index"),
		esClient.Search.WithBody(strings.NewReader(query)),
	)
	fmt.Println("Search executed successfully", err, result.IsError(), result.String())
}

func userInterface() (err error, cfg elasticsearch.Config) {
	cfg = elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "changeme",
	}
	return
}
