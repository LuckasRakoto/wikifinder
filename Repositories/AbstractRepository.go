package Repository

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

var driver neo4j.DriverWithContext

func Connect(ctx context.Context) {

	dbUri := os.Getenv("NEO4J_URI")
	dbUser := os.Getenv("NEO4J_USER")
	dbPassword := os.Getenv("NEO4J_PASSWORD")

	var err error
	driver, err = neo4j.NewDriverWithContext(
		dbUri,
		neo4j.BasicAuth(dbUser, dbPassword, ""))
	if err != nil {
		fmt.Println("Error creating driver")
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		panic(err)
	}
}

func Driver() neo4j.DriverWithContext {
	return driver
}

func Disconnect(ctx context.Context) {
	err := driver.Close(ctx)
	if err != nil {
		fmt.Println("Error when disconnecting")
	}
}
