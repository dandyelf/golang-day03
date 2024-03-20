package main

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"io"
	"log"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

const (
	csvFileName = "../../materials/data.csv"
	indexName   = "places"
)

//go:embed schema.json
var jsonSchema string

type SearchResult struct {
	Hits struct {
		Hits []struct {
			ID     string `json:"_id"`
			Source struct {
				Name    string `json:"name"`
				Address string `json:"address"`
				Phone   string `json:"phone"`
			} `json:"_source"`
		} `json:"hits"`
	} `json:"hits"`
}

var (
	response *esapi.Response
	err      error
	esClient *elasticsearch.Client
)

func main() {
	esClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	code, err := esClient.Ping()
	if err != nil {
		log.Fatalf("Error pinging the Elasticsearch server: %s", err)
	}
	log.Println("Elasticsearch returned with code___", code.Status())

	if response, err = esClient.Indices.Delete([]string{indexName}, esClient.Indices.Delete.WithIgnoreUnavailable(true)); err != nil || response.IsError() {
		log.Fatalf("Cannot delete index: %s", err)
	}
	response.Body.Close()

	res, err := esClient.Indices.Create(indexName, esClient.Indices.Create.WithBody(strings.NewReader(jsonSchema)))
	if err != nil {
		log.Fatalf("Cannot create index: %s", err)
	}
	if res.IsError() {
		log.Fatalf("Response error: Cannot create index: %s", res)
	}
	if err = res.Body.Close(); err != nil {
		log.Fatalf("Internal err %s", err)
	}
	settingsAdd()
	getCsv(csvFileName)
	putEs(esClient)
	getIndex(esClient)
	searchDB(esClient)
	testGet(esClient)
}

func settingsAdd() {
	settings := `{
		"index" : {
			"max_result_window" : 20000
		}
	}`
	req := esapi.IndicesPutSettingsRequest{
		Index: []string{indexName},
		Body:  strings.NewReader(settings),
	}

	res, err := req.Do(context.Background(), esClient)
	if err != nil {
		log.Fatalf("Error updating index settings: %s", err)
	}
	defer res.Body.Close()
	log.Println("Index settings updated successfully")
}

func searchDB(es *elasticsearch.Client) {
	log.Println("Search DB Req: ")

	var query = map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"name": "Doner Kebab",
			},
		},
	}

	queryJSON, err := json.Marshal(query)
	if err != nil {
		log.Fatalf("Error marshalling query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex(indexName),
		es.Search.WithBody(strings.NewReader(string(queryJSON))),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error performing search request: %s", err)
	}

	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	}

	var result SearchResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}

	for _, hit := range result.Hits.Hits {
		log.Printf("ID: %s, Name: %s, Address: %s, Phone: %s",
			hit.ID, hit.Source.Name, hit.Source.Address, hit.Source.Phone)
	}
}

func getIndex(es *elasticsearch.Client) {
	res, err := es.Indices.Get([]string{"places"})
	if err != nil {
		log.Fatalf("Error getting mapping and settings: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error response: %s", res.String())
	}
	// Parse and print the response
	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		log.Fatalf("Error parsing the response body: %s", err)
	}
	// Print the response
	prettyJSON, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Fatalf("Error formatting JSON: %s", err)
	}
	log.Printf("Get index response:\n%s", string(prettyJSON))
}

func testGet(es *elasticsearch.Client) {
	res, err := es.Get(indexName, "0")
	if err != nil {
		log.Fatalf("Error getting document: %s", err)
	}

	body, _ := io.ReadAll(res.Body)
	res.Body.Close()

	var out bytes.Buffer
	json.Indent(&out, body, "", "  ")

	log.Printf("testPlace Response:\n%s", out.String())
}
