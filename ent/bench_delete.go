// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nibbleshift/gorb/ent/bench"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// BenchDelete is the builder for deleting a Bench entity.
type BenchDelete struct {
	config
	hooks    []Hook
	mutation *BenchMutation
}

// Where appends a list predicates to the BenchDelete builder.
func (bd *BenchDelete) Where(ps ...predicate.Bench) *BenchDelete {
	bd.mutation.Where(ps...)
	return bd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (bd *BenchDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BenchMutation](ctx, bd.sqlExec, bd.mutation, bd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BenchDelete) ExecX(ctx context.Context) int {
	n, err := bd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (bd *BenchDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(bench.Table, sqlgraph.NewFieldSpec(bench.FieldID, field.TypeInt))
	if ps := bd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, bd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	bd.mutation.done = true
	return affected, err
}

// BenchDeleteOne is the builder for deleting a single Bench entity.
type BenchDeleteOne struct {
	bd *BenchDelete
}

// Where appends a list predicates to the BenchDelete builder.
func (bdo *BenchDeleteOne) Where(ps ...predicate.Bench) *BenchDeleteOne {
	bdo.bd.mutation.Where(ps...)
	return bdo
}

// Exec executes the deletion query.
func (bdo *BenchDeleteOne) Exec(ctx context.Context) error {
	n, err := bdo.bd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{bench.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BenchDeleteOne) ExecX(ctx context.Context) {
	if err := bdo.Exec(ctx); err != nil {
		panic(err)
	}
}
