package main

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
	apiData, apiErr := i.api.GetArticleData(downloadURL)

	if apiErr != nil {
		return apiErr
	}

	parsedData, parseErr := i.parser.ParseArticleData(apiData)

	if parseErr != nil {
		return parseErr
	}

	indexingErr := i.datastore.IndexArticleData(parsedData)

	if indexingErr != nil {
		return indexingErr
	}

	return nil
}

func NewIndexer() Indexer {
	return Indexer{
		api:       sbArticleAPI{},
		parser:    Article{},
		datastore: DatastoreIndexer{},
	}
}
