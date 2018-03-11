package indexer

import (
	"fmt"
	"strconv"
	"time"

	"log"

	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

type DatastoreIndexer struct {
	serverURL string
	indexName string
}

func (datastore DatastoreIndexer) IndexArticleData(articles []Article) error {

	client, err := retryConnect(15, 5*time.Second, func() (*elastic.Client, error) {
		return elastic.NewClient(elastic.SetURL(datastore.serverURL))
	})
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

	for index, article := range articles {
		_, err = client.Index().
			Index(datastore.indexName).
			Type("article").
			Id(strconv.Itoa(article.ArticleID)).
			BodyJson(article).
			Do(context.TODO())
		if err != nil {
			log.Printf("Failed to index article. %q", article)
		}
		if index%1000 == 0 {
			log.Printf("Indexed %d articles", index)
		}
	}

	return nil
}

func retryConnect(attempts int, sleep time.Duration, callback func() (*elastic.Client, error)) (client *elastic.Client, err error) {
	for i := 0; ; i++ {
		client, err := callback()
		if err == nil {
			return client, nil
		}

		if i >= (attempts - 1) {
			break
		}

		time.Sleep(sleep)

		log.Println("Retry connecting to datastore after error:", err)
	}
	return nil, fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}
