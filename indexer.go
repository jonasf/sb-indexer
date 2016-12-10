package main

type sbAPI interface {
	GetArticleData(downloadURL string) ([]byte, error)
}

type apiDataParser interface {
	ParseArticleData(data []byte) ([]Article, error)
}

type Indexer struct {
	api    sbAPI
	parser apiDataParser
}

func (i Indexer) Index(downloadURL string) error {
	apiData, apiErr := i.api.GetArticleData(downloadURL)

	if apiErr != nil {
		return apiErr
	}

	_, parseErr := i.parser.ParseArticleData(apiData)

	if parseErr != nil {
		return parseErr
	}

	return nil
}

/*func NewIndexer() (*Indexer, error) {
	api := &sbAPI{}
	i := &Indexer{
		api: api,
	}
	return i, nil
}*/
