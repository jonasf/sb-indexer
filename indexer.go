package main

import (
	"log"
)

type sbAPI interface {
	GetArticleData(downloadURL string) ([]byte, error)
}

type apiDataParser interface {
	ParseArticleData(data []byte) ([]Article, error)
}

type indexDataStore interface {
	IndexArticleData(articles []Article) error
}

type Indexer struct {
	api       sbAPI
	parser    apiDataParser
	datastore indexDataStore
}

func (i Indexer) Index(downloadURL string) error {

	log.Print("Data indexing start")

	apiData, apiErr := i.api.GetArticleData(downloadURL)

	if apiErr != nil {
		log.Fatalf("Failed to download article data. %s", apiErr)
		return apiErr
	}

	parsedData, parseErr := i.parser.ParseArticleData(apiData)

	if parseErr != nil {
		log.Fatalf("Failed to parse article data. %s", parseErr)
		return parseErr
	}

	indexingErr := i.datastore.IndexArticleData(parsedData)

	if indexingErr != nil {
		log.Fatalf("Failed to index article data. %s", indexingErr)
		return indexingErr
	}

	log.Print("Data indexing completed")

	return nil
}

func NewIndexer(datastoreURL, indexName string) Indexer {
	return Indexer{
		api:    sbArticleAPI{},
		parser: Article{},
		datastore: DatastoreIndexer{
			serverURL: datastoreURL,
			indexName: indexName,
		},
	}
}
