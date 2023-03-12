package gorb

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nibbleshift/gorb/ent"
)

func (r *mutationResolver) CreateBenchmark(ctx context.Context, input ent.CreateBenchInput) (*ent.Bench, error) {
	return r.client.Bench.Create().SetInput(input).Save(ctx)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
