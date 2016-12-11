package main

import (
	"fmt"
)

type DatastoreIndexer struct{}

func (datastore DatastoreIndexer) IndexArticleData(articles []Article) error {

	i := 0

	for _, article := range articles {
		i++
		fmt.Println(article.ArticleID, article.ArticleType, article.Name, article.SecondaryName)
	}

	fmt.Println("Total number or processed articles: ", i)
	return nil
}
