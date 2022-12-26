package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Genialngash/graphql-go-test/graph"
	"github.com/Genialngash/graphql-go-test/postgress"
	"github.com/go-pg/pg/v10"
)

const defaultPort = "8080"

func main() {
	DB := postgress.New(&pg.Options{
		User:     "ngash",
		Password: "login",
		Database: "meetup_dev",
	})
	defer DB.Close()

	DB.AddQueryHook(postgress.DbLogger{})

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		MeetupRepo: postgress.MeetupsRepo{DB: DB},
		UsersRepo:  postgress.UsersRepo{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
