package Parser

import (
	"fmt"
	"gowiki/Repositories/ArticleRepository"
	"io"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const TITLE_CLASS = ".mw-page-title-main"
const BASE_URL = "https://en.wikipedia.org"

var seen []string
var queue []string

func fetch(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil
	}
	return doc
}

func getTitle(node *goquery.Document) string {
	title := node.Find(TITLE_CLASS).First().Text()
	return title
}

func findNeighbors(node *goquery.Document) {
	node.Find("a").Each(func(_ int, node *goquery.Selection) {
		link, exists := node.Attr("href")
		if exists && isRedirect(link) {
			queue = append(queue, buildLink(link))
		}
	})
}

func isRedirect(s string) bool {
	return strings.Contains(s, "/wiki/") && !strings.Contains(s, "https://")
}

func removeAd(node *goquery.Document) {
	node.Find("table.box-More_citations_needed").Remove()
}

func removeHead(node *goquery.Document) {
	node.Find("div.vector-header-container").Remove()
}

func removeFooter(node *goquery.Document) {
	node.Find("div.mw-footer-container").Remove()
}

func removeBadLinks(node *goquery.Document) {
	removeAd(node)
	removeHead(node)
	removeFooter(node)
}

func parseArticle(url string) {
	seen = append(seen, url)
	document := fetch(url)
	removeBadLinks(document)
	findNeighbors(document)
	title := getTitle(document)
	ArticleRepository.Insert(ArticleRepository.Article{
		Title: title,
		Url:   url,
	})
}

func buildLink(redirect string) string {
	return fmt.Sprintf("%s%s", BASE_URL, redirect)
}

func Start(startUrl string) {
	parseArticle(buildLink(startUrl))
	// QUEUE is full of neighbors
}
