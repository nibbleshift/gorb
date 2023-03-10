// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (b *Bench) BenchResult(ctx context.Context) (result []*BenchResult, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = b.NamedBenchResult(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = b.Edges.BenchResultOrErr()
	}
	if IsNotLoaded(err) {
		result, err = b.QueryBenchResult().All(ctx)
	}
	return result, err
}
