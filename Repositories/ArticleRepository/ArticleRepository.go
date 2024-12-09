package ArticleRepository

import (
	"context"
	"fmt"
	Repository "gowiki/Repositories"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Article struct {
	Title string
	Url   string
}

func Insert(article Article) {
	result, err := neo4j.ExecuteQuery(context.Background(), Repository.Driver(),
		"CREATE (a:Article {title: $title, url: $url}) RETURN a",
		map[string]any{
			"title": article.Title,
			"url":   article.Url,
		}, neo4j.EagerResultTransformer,
		neo4j.ExecuteQueryWithDatabase("neo4j"))
	if err != nil {
		panic(err)
	}
	summary := result.Summary
	fmt.Printf("Created %v nodes in %+v.\n",
		summary.Counters().NodesCreated(),
		summary.ResultAvailableAfter())
}
