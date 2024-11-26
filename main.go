package main

import "gowiki/Parser"

func main() {
	start := "https://en.wikipedia.org/wiki/Prime_Minister_of_Lithuania"

	Parser.ParseArticle(start)
}
