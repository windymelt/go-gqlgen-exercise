package main

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/handler"
)

func main() {
	http.Handle("/", handler.Playground("GraphQL Playground", "/query"))
	http.Handle("/query", handler.GraphQL(NewExecutableSchema(Config{Resolvers: &Resolver{}})))

	log.Print("listening on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
