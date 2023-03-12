package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/nibbleshift/gorb/ent"
	"github.com/nibbleshift/gorb/graph/generated"
)

// AllocedBytesPerOp is the resolver for the allocedbytesperop field.
func (r *benchResultResolver) AllocedBytesPerOp(ctx context.Context, obj *ent.BenchResult) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// AllocsPerOp is the resolver for the allocsperop field.
func (r *benchResultResolver) AllocsPerOp(ctx context.Context, obj *ent.BenchResult) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Benches is the resolver for the benches field.
func (r *queryResolver) Benches(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int) (*ent.BenchConnection, error) {
	return r.client.Bench.Query().
		Paginate(ctx, after, first, before, last)
}

// BenchResults is the resolver for the benchResults field.
func (r *queryResolver) BenchResults(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int) (*ent.BenchResultConnection, error) {
	panic(fmt.Errorf("not implemented: BenchResults - benchResults"))
}

// Allocedbytesperop is the resolver for the allocedbytesperop field.
func (r *createBenchResultInputResolver) Allocedbytesperop(ctx context.Context, obj *ent.CreateBenchResultInput, data int) error {
	panic(fmt.Errorf("not implemented: Allocedbytesperop - allocedbytesperop"))
}

// Allocsperop is the resolver for the allocsperop field.
func (r *createBenchResultInputResolver) Allocsperop(ctx context.Context, obj *ent.CreateBenchResultInput, data int) error {
	panic(fmt.Errorf("not implemented: Allocsperop - allocsperop"))
}

// BenchResult returns generated.BenchResultResolver implementation.
func (r *Resolver) BenchResult() generated.BenchResultResolver { return &benchResultResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// CreateBenchResultInput returns generated.CreateBenchResultInputResolver implementation.
func (r *Resolver) CreateBenchResultInput() generated.CreateBenchResultInputResolver {
	return &createBenchResultInputResolver{r}
}

type benchResultResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type createBenchResultInputResolver struct{ *Resolver }
