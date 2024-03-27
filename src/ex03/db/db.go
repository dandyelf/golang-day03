package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"myGeoserv/types"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

var (
	response *esapi.Response
	err      error
	esClient *elasticsearch.Client
)

func init() {
	startDbCli()
}

type PlaceStore struct {
	// Дополнительные поля, если необходимо
}

func startDbCli() {
	esClient, err = elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	response, err = esClient.Ping()
	if err != nil {
		log.Fatalf("Error pinging the Elasticsearch server: %s", err)
	}
	log.Println("Elasticsearch returned with code___", response.Status())
}

func (ps *PlaceStore) GetPlaces(limit int, offset int) ([]types.Place, int, error) {
	from := offset*limit - limit + 1
	var b bytes.Buffer
	fmt.Fprintf(&b, `{
		  "from": %v,
		  "size": %v,
		  "track_total_hits": true,
		  "query": {
			"match_all": {}
		  }
		}`, from, limit)
	response, _ = esClient.Search(
		esClient.Search.WithBody(&b),
		esClient.Search.WithPretty(),
	)
	defer response.Body.Close()
	if response.IsError() {
		return nil, 0, fmt.Errorf("error response: %s", response.String())
	}
	var result types.SearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, 0, fmt.Errorf("error parsing the response body: %s", err)
	}
	var list []types.Place
	for _, hit := range result.Hits.Hits {
		list = append(list, hit.Source)
	}
	var maxVall int
	if result.Hits.Total.Value < 20000 {
		maxVall = result.Hits.Total.Value
	} else {
		maxVall = 20000
	}
	return list, maxVall, nil
}

func (ps *PlaceStore) GetGeoPlaces(location types.Location) ([]types.Place, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, `{
		"size": 3,
		"sort": [
    		{
      			"_geo_distance": {
        		"location": {
          			"lat": %v,
          			"lon": %v
        	},
        "order": "asc",
        "unit": "km",
        "mode": "min",
        "distance_type": "arc",
        "ignore_unmapped": true
      }
    }
]

		}`, location.Lat, location.Lon)
	response, _ = esClient.Search(
		esClient.Search.WithBody(&b),
		esClient.Search.WithPretty(),
	)
	defer response.Body.Close()
	if response.IsError() {
		return nil, fmt.Errorf("error response: %s", response.String())
	}
	var result types.SearchResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("error parsing the response body: %s", err)
	}
	var list []types.Place
	for _, hit := range result.Hits.Hits {
		list = append(list, hit.Source)
	}
	return list, nil
}
