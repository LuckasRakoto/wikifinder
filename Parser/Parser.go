package Parser

import (
	"io"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const TITLE_CLASS = ".mw-page-title-main"

var logger = log.Default()

func fetch(url string) *goquery.Document {
	logger.Printf("Fetching %s", url)
	resp, err := http.Get(url)
	if err != nil {
		logger.Fatalf("Encountered error when fetching data %s: %s", url, err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logger.Fatalf("Encountered error when closing the response body: %s", err)
		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		logger.Fatalf("Encountered error when parsing body: %s", err)
		return nil
	}
	return doc
}

func getTitle(node *goquery.Document) string {
	var title string
	node.Find(TITLE_CLASS).Each(func(_ int, node *goquery.Selection) {
		title = node.Text()
	})
	logger.Printf("Found title %s", title)
	return title
}

func ParseArticle(url string) {
	logger.Printf("Parsing article %s", url)
	document := fetch(url)
	getTitle(document)
}
