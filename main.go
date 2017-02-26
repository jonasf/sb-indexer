package main

import (
	"flag"
	"log"
	"os"
	"time"
)

func main() {

	//Delayed start to wait for dependencies to start
	delayedStartDuration := os.Getenv("DELAYED-START")
	if delayedStartDuration != "" {
		duration, _ := time.ParseDuration(delayedStartDuration)
		log.Println("Delaying start for ", duration)
		time.Sleep(duration)
	}

	esURL := os.Getenv("ES-URL")

	if esURL == "" {
		esURLPtr := flag.String("es-url", "http://localhost:9200", "ElasticSearch URL")
		flag.Parse()
		esURL = *esURLPtr
	}
	log.Println("Elasticsearch server address: ", esURL)

	esIndexNamePtr := flag.String("es-index-name", "articles", "ElasticSearch index name")
	sbAPIURLPtr := flag.String("sb-api-url", "https://www.systembolaget.se/api/assortment/products/xml", "Systembolaget API URL")

	flag.Parse()

	log.Println("es-url:", esURL)
	log.Println("es-index-name:", *esIndexNamePtr)
	log.Println("sb-api-url:", *sbAPIURLPtr)

	indexer := NewIndexer(esURL, *esIndexNamePtr)
	indexer.Index(*sbAPIURLPtr)
}
