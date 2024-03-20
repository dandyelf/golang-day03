package main

import (
	"context"
	_ "embed"
	"encoding/csv"
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
	"os"
	"strconv"
	"strings"
)

type Place struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Location struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	} `json:"location"`
}

var list []Place

func getCsv(csvFileName string) {
	f, err := os.Open(csvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = '\t'
	for {
		record, err := r.Read()
		if err != nil {
			log.Println("CSV err: ", err)
			break
		}
		if record[0] == "" {

			continue
		}
		var place Place
		if place.ID, err = strconv.ParseInt(record[0], 10, 64); err != nil {
			log.Println("getCsv err: ", err)
		}
		place.Name = record[1]
		place.Address = record[2]
		place.Phone = record[3]
		if place.Location.Lat, err = strconv.ParseFloat(record[4], 64); err != nil {
			log.Println("Parse Lat err: ", err)
		}
		if place.Location.Lon, err = strconv.ParseFloat(record[5], 64); err != nil {
			log.Println("Parse Lon err: ", err)
		}
		if checkPlace(place) {
			list = append(list, place)
		}
	}

	log.Println("Data imported successfully")
	log.Println("Number of records: ", len(list))
	log.Println("First record: ", list[0])
}

func checkPlace(place Place) bool {
	return isValidLocation(place.Location.Lat, place.Location.Lon)
}

func isValidLocation(lat float64, lon float64) bool {
	return (lat >= -90 && lat <= 90) && (lon >= -180 && lon <= 180)
}

func putEs(es *elasticsearch.Client) {
	// Prepare the Bulk API request body
	var bulk strings.Builder
	for _, restaurant := range list {

		metaData := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": indexName,
				"_id":    restaurant.ID,
			},
		}
		metaDataJSON, _ := json.Marshal(metaData)

		restaurantJSON, _ := json.Marshal(restaurant)

		bulk.WriteString(string(metaDataJSON) + "\n")
		bulk.WriteString(string(restaurantJSON) + "\n")
	}

	// Perform the Bulk API request
	req := esapi.BulkRequest{
		Body:    strings.NewReader(bulk.String()),
		Refresh: "true",
	}
	response, err := req.Do(context.Background(), es)
	if err != nil {
		log.Fatalf("Error performing Bulk request: %s", err)
	}
	defer response.Body.Close()

	// Check the response status
	if response.IsError() {
		log.Fatalf("Error response: %s", response.String())
	}

	log.Println("Data sent to Elasticsearch successfully")
}
