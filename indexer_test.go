package main

import (
	"errors"
	"testing"
)

type APIFetchDataFailStub struct{}

func (api *APIFetchDataFailStub) GetArticleData(downloadURL string) ([]byte, error) {
	return nil, errors.New("Data fetch failed")
}

type APIFetchDataStub struct{}

func (api *APIFetchDataStub) GetArticleData(downloadURL string) ([]byte, error) {
	data := make([]byte, 5, 5)
	return data, nil
}

type APIDataParseFailStub struct{}

func (parse APIDataParseFailStub) ParseArticleData(data []byte) ([]Article, error) {
	return nil, errors.New("Data parse failed")
}

func TestFailedApiDataFetch(t *testing.T) {
	api := &APIFetchDataFailStub{}
	i := &Indexer{
		api: api,
	}

	err := i.Index("https://www.systembolaget.se/api/assortment/products/xml")

	if err.Error() != "Data fetch failed" {
		t.Errorf("Expected error message to be \"Data fetch failed\" but got %q", err)
	}
}

func TestFailedApiDataParse(t *testing.T) {
	api := &APIFetchDataStub{}
	parser := APIDataParseFailStub{}
	i := &Indexer{
		api:    api,
		parser: parser,
	}

	err := i.Index("https://www.systembolaget.se/api/assortment/products/xml")

	if err.Error() != "Data parse failed" {
		t.Errorf("Expected error message to be \"Data parse failed\" but got %q", err)
	}
}
