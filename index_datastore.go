package main

import (
	"fmt"
	"strconv"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

type DatastoreIndexer struct{}

func (datastore DatastoreIndexer) IndexArticleData(articles []Article) error {

	client, err := elastic.NewClient()
	if err != nil {
		return err
	}

	indexName := "articles"
	exists, err := client.IndexExists(indexName).Do(context.TODO())
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex(indexName).Do(context.TODO())
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	for _, article := range articles {
		_, err = client.Index().
			Index(indexName).
			Type("article").
			Id(strconv.Itoa(article.ArticleID)).
			BodyJson(article).
			Do(context.TODO())
		if err != nil {
			fmt.Println(err)
		}
	}

	return nil
}
