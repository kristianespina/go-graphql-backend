package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/jackc/pgx/v4"
)

func main() {
	s := `
	schema {
		query: Query
	}

	type Query {
		GetThread(): Thread
	}

	type Thread {
		id: String
		author: String
		title: String
		date_posted: String
	}
	`
	fmt.Println(" Connecting to Postgres Database...")
	conn, err := pgx.Connect(context.Background(), os.Getenv("postgres://postgres@localhost:5432"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	schema := graphql.MustParseSchema(s, &Resolver{})
	fmt.Println(" 🚀  - Server is now listening on port 8080")
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
