package graph

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/nibbleshift/gorb/ent"
	"github.com/nibbleshift/gorb/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client *ent.Client
}

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
