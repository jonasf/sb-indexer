package main

func main() {
	indexer := NewIndexer("http://localhost:9200", "articles")
	indexer.Index("https://www.systembolaget.se/api/assortment/products/xml")
}
