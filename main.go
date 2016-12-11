package main

func main() {
	indexer := NewIndexer()
	indexer.Index("https://www.systembolaget.se/api/assortment/products/xml")
}
