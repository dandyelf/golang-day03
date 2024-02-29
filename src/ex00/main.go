package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	"log"
)

func main() {
	es, _ := elasticsearch.NewDefaultClient()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}
