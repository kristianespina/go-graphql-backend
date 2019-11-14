package main

import (
	"fmt"
	"log"
	"net/http"

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
	
	schema := graphql.MustParseSchema(s, &Resolver{})
	fmt.Println(" 🚀  - Server is now listening on port 8080")
	http.Handle("/query", &relay.Handler{Schema: schema})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
