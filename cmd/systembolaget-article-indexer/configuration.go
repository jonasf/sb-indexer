package main

import (
	"flag"
	"os"
)

type Configuration struct {
	ElasticsearchURL    string
	ElasticsearchIndex  string
	SystembolagetAPIURL string
}

func LoadConfig() Configuration {
	elasticsearchURL := flag.String("es-url", "http://localhost:9200", "ElasticSearch URL")
	elasticsearchIndexName := flag.String("es-index-name", "articles", "ElasticSearch index name")
	systembolagetAPIURL := flag.String("sb-api-url", "https://www.systembolaget.se/api/assortment/products/xml", "Systembolaget API URL")
	flag.Parse()

	configuration := Configuration{
		ElasticsearchURL:    os.Getenv("ELASTICSEARCH_URL"),
		ElasticsearchIndex:  os.Getenv("ELASTICSEARCH_INDEX"),
		SystembolagetAPIURL: os.Getenv("SYSTEMBOLAGET_API_URL"),
	}

	if configuration.ElasticsearchURL == "" {
		configuration.ElasticsearchURL = *elasticsearchURL
	}
	if configuration.ElasticsearchIndex == "" {
		configuration.ElasticsearchIndex = *elasticsearchIndexName
	}
	if configuration.SystembolagetAPIURL == "" {
		configuration.SystembolagetAPIURL = *systembolagetAPIURL
	}

	return configuration
}
