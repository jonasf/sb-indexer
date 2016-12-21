package main

import (
	"flag"
	"log"
)

func main() {

	esURLPtr := flag.String("es-url", "http://localhost:9200", "ElasticSearch URL")
	esIndexNamePtr := flag.String("es-index-name", "articles", "ElasticSearch index name")
	sbAPIURLPtr := flag.String("sb-api-url", "https://www.systembolaget.se/api/assortment/products/xml", "Systembolaget API URL")

	flag.Parse()

	log.Println("es-url:", *esURLPtr)
	log.Println("es-index-name:", *esIndexNamePtr)
	log.Println("sb-api-url:", *sbAPIURLPtr)

	indexer := NewIndexer(*esURLPtr, *esIndexNamePtr)
	indexer.Index(*sbAPIURLPtr)
}
