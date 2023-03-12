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

	_ "github.com/lib/pq"
)

func main() {
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
	http.Handle("/",
		playground.Handler("gorb", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
