// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nibbleshift/gorb/ent/benchresult"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// BenchResultDelete is the builder for deleting a BenchResult entity.
type BenchResultDelete struct {
	config
	hooks    []Hook
	mutation *BenchResultMutation
}

// Where appends a list predicates to the BenchResultDelete builder.
func (brd *BenchResultDelete) Where(ps ...predicate.BenchResult) *BenchResultDelete {
	brd.mutation.Where(ps...)
	return brd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (brd *BenchResultDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, BenchResultMutation](ctx, brd.sqlExec, brd.mutation, brd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (brd *BenchResultDelete) ExecX(ctx context.Context) int {
	n, err := brd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (brd *BenchResultDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(benchresult.Table, sqlgraph.NewFieldSpec(benchresult.FieldID, field.TypeInt))
	if ps := brd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, brd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	brd.mutation.done = true
	return affected, err
}

// BenchResultDeleteOne is the builder for deleting a single BenchResult entity.
type BenchResultDeleteOne struct {
	brd *BenchResultDelete
}

// Where appends a list predicates to the BenchResultDelete builder.
func (brdo *BenchResultDeleteOne) Where(ps ...predicate.BenchResult) *BenchResultDeleteOne {
	brdo.brd.mutation.Where(ps...)
	return brdo
}

// Exec executes the deletion query.
func (brdo *BenchResultDeleteOne) Exec(ctx context.Context) error {
	n, err := brdo.brd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{benchresult.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (brdo *BenchResultDeleteOne) ExecX(ctx context.Context) {
	if err := brdo.Exec(ctx); err != nil {
		panic(err)
	}
}
