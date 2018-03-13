package main

import (
	"log"

	indexer "github.com/jonasf/systembolaget-article-indexer/internal/systembolaget-article-indexer"
)

func main() {
	configuration := LoadConfig()

	log.Println("Using Elasticsearch Server:", configuration.ElasticsearchURL)
	log.Println("Using Elasticsearch Index Name:", configuration.ElasticsearchIndex)
	log.Println("Using Systembolaget API Url:", configuration.SystembolagetAPIURL)

	indexer := indexer.NewIndexer(configuration.ElasticsearchURL, configuration.ElasticsearchIndex)
	indexer.Index(configuration.SystembolagetAPIURL)
}
