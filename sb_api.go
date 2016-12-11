package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

type sbArticleAPI struct{}

func (api sbArticleAPI) GetArticleData(downloadURL string) ([]byte, error) {
	filePath := "articles.xml"
	err := downloadArticleDataFile(downloadURL, filePath)

	if err != nil {
		return nil, err
	}

	file, err := os.Open(filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	rawData, err := ioutil.ReadAll(file)

	if err != nil {
		return nil, err
	}

	return rawData, nil
}

func downloadArticleDataFile(downloadURL string, destination string) error {
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(downloadURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
