package main

import (
	"context"
	"log"
	"net/http"

	"github.com/nibbleshift/gorb/ent"
	"github.com/nibbleshift/gorb/ent/migrate"
	"github.com/nibbleshift/gorb/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	_ "github.com/lib/pq"
)

func main() {
	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	// Create ent.Client and run the schema migration.
	client, err := ent.Open("postgres", "user=postgres password=example dbname=gorb sslmode=disable")

	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(graph.NewSchema(client))
	router.Handle("/", playground.Handler("Starwars", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":8081", router)
	if err != nil {
		panic(err)
	}
}
