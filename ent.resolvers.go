package gorb

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nibbleshift/gorb/ent"
)

func (r *benchResultResolver) AllocedBytesPerOp(ctx context.Context, obj *ent.BenchResult) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *benchResultResolver) AllocsPerOp(ctx context.Context, obj *ent.BenchResult) (int, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Benches(ctx context.Context) ([]*ent.Bench, error) {
	panic(fmt.Errorf("not implemented"))
}

// BenchResult returns BenchResultResolver implementation.
func (r *Resolver) BenchResult() BenchResultResolver { return &benchResultResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type benchResultResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
