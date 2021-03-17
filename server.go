package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/paul/go-server/dataloader"
	"github.com/paul/go-server/graph"
	database "github.com/paul/go-server/graph/db"
	"github.com/paul/go-server/graph/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	database.InitDB()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloader.Middleware(database.Db, srv))

	log.Printf("connect to http://localhost:%s/query for GraphQL queries", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
