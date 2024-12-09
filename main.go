package main

import (
	"context"
	"fmt"
	Repository "gowiki/Repositories"
	"gowiki/Services/Parser"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Couldn't load .env")
	}

	start := "/wiki/Prime_Minister_of_Lithuania"

	ctx := context.Background()

	Repository.Connect(ctx)
	defer Repository.Disconnect(ctx)

	Parser.Start(start)
}
