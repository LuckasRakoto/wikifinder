package main

import (
	"context"
	Repository "gowiki/Repositories"
	"gowiki/Services/Parser"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	start := "/wiki/Prime_Minister_of_Lithuania"

	ctx := context.Background()

	Repository.Connect(ctx)
	defer Repository.Disconnect(ctx)

	Parser.Start(start)
}
