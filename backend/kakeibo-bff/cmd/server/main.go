package main

import (
    "log"
    "net/http"
    "os"

    "github.com/99designs/gqlgen/graphql/handler"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/ZONO33LHD/sircle/backend/kakeibo-bff/graph"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("USERS_PORT")
	if port == "" {
		port = defaultPort
	}

	es := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}})
	srv := handler.NewDefaultServer(es)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
