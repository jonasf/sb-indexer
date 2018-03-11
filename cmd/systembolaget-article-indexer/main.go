package main

import (
	"flag"
	"log"
	"os"

	indexer "github.com/jonasf/sb-indexer/internal/systembolaget-article-indexer"
)

func main() {

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

	indexer := indexer.NewIndexer(esURL, *esIndexNamePtr)
	indexer.Index(*sbAPIURLPtr)
}
