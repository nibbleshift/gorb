package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/nibbleshift/gorb/ent"
	"github.com/nibbleshift/gorb/internal/exec"
)

func main() {
	e := exec.NewExec()
	e.Run("/usr/local/bin/go", "-test", "-bench=.")

	client, err := ent.Open("postgres", "user=postgres password=example dbname=gorb sslmode=disable")

	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
