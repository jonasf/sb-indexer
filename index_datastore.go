package main

import (
	"strconv"

	"log"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

type DatastoreIndexer struct {
	serverURL string
	indexName string
}

func (datastore DatastoreIndexer) IndexArticleData(articles []Article) error {

	client, err := elastic.NewClient(elastic.SetURL(datastore.serverURL))
	if err != nil {
		return err
	}

	exists, err := client.IndexExists(datastore.indexName).Do(context.TODO())
	if err != nil {
		log.Panicf("Failed to connect to server. %s", err)
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex(datastore.indexName).Do(context.TODO())
		if err != nil {
			log.Panicf("Failed to create index. %s", err)
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	for _, article := range articles {
		_, err = client.Index().
			Index(datastore.indexName).
			Type("article").
			Id(strconv.Itoa(article.ArticleID)).
			BodyJson(article).
			Do(context.TODO())
		if err != nil {
			log.Printf("Failed to index article. %q", article)
		}
	}

	return nil
}
