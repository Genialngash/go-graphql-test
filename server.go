package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/Genialngash/graphql-go-test/customemiddleware"
	"github.com/Genialngash/graphql-go-test/graph"
	"github.com/Genialngash/graphql-go-test/postgress"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-pg/pg/v10"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	DB := postgress.New(&pg.Options{
		Addr:     "postgres:5432",
		User:     "ngash",
		Password: "login",
		Database: "meetup_dev",
	})

	defer DB.Close()

	DB.AddQueryHook(postgress.DbLogger{})
	userrepo := postgress.UsersRepo{DB: DB}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		Debug:            true,
		AllowedOrigins:   []string{"http://localhost:8080"},
		AllowCredentials: true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(customemiddleware.AuthMiddleware(userrepo))

	c := graph.Config{Resolvers: &graph.Resolver{
		MeetupRepo: postgress.MeetupsRepo{DB: DB},
		UsersRepo:  userrepo,
	}}
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", graph.DataLoaderMiddleWare(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
