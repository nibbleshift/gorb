// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nibbleshift/gorb/ent/benchresult"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// BenchResultQuery is the builder for querying BenchResult entities.
type BenchResultQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.BenchResult
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*BenchResult) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BenchResultQuery builder.
func (brq *BenchResultQuery) Where(ps ...predicate.BenchResult) *BenchResultQuery {
	brq.predicates = append(brq.predicates, ps...)
	return brq
}

// Limit the number of records to be returned by this query.
func (brq *BenchResultQuery) Limit(limit int) *BenchResultQuery {
	brq.ctx.Limit = &limit
	return brq
}

// Offset to start from.
func (brq *BenchResultQuery) Offset(offset int) *BenchResultQuery {
	brq.ctx.Offset = &offset
	return brq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (brq *BenchResultQuery) Unique(unique bool) *BenchResultQuery {
	brq.ctx.Unique = &unique
	return brq
}

// Order specifies how the records should be ordered.
func (brq *BenchResultQuery) Order(o ...OrderFunc) *BenchResultQuery {
	brq.order = append(brq.order, o...)
	return brq
}

// First returns the first BenchResult entity from the query.
// Returns a *NotFoundError when no BenchResult was found.
func (brq *BenchResultQuery) First(ctx context.Context) (*BenchResult, error) {
	nodes, err := brq.Limit(1).All(setContextOp(ctx, brq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{benchresult.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (brq *BenchResultQuery) FirstX(ctx context.Context) *BenchResult {
	node, err := brq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BenchResult ID from the query.
// Returns a *NotFoundError when no BenchResult ID was found.
func (brq *BenchResultQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = brq.Limit(1).IDs(setContextOp(ctx, brq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{benchresult.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (brq *BenchResultQuery) FirstIDX(ctx context.Context) int {
	id, err := brq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BenchResult entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BenchResult entity is found.
// Returns a *NotFoundError when no BenchResult entities are found.
func (brq *BenchResultQuery) Only(ctx context.Context) (*BenchResult, error) {
	nodes, err := brq.Limit(2).All(setContextOp(ctx, brq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{benchresult.Label}
	default:
		return nil, &NotSingularError{benchresult.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (brq *BenchResultQuery) OnlyX(ctx context.Context) *BenchResult {
	node, err := brq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BenchResult ID in the query.
// Returns a *NotSingularError when more than one BenchResult ID is found.
// Returns a *NotFoundError when no entities are found.
func (brq *BenchResultQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = brq.Limit(2).IDs(setContextOp(ctx, brq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{benchresult.Label}
	default:
		err = &NotSingularError{benchresult.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (brq *BenchResultQuery) OnlyIDX(ctx context.Context) int {
	id, err := brq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BenchResults.
func (brq *BenchResultQuery) All(ctx context.Context) ([]*BenchResult, error) {
	ctx = setContextOp(ctx, brq.ctx, "All")
	if err := brq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BenchResult, *BenchResultQuery]()
	return withInterceptors[[]*BenchResult](ctx, brq, qr, brq.inters)
}

// AllX is like All, but panics if an error occurs.
func (brq *BenchResultQuery) AllX(ctx context.Context) []*BenchResult {
	nodes, err := brq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BenchResult IDs.
func (brq *BenchResultQuery) IDs(ctx context.Context) (ids []int, err error) {
	if brq.ctx.Unique == nil && brq.path != nil {
		brq.Unique(true)
	}
	ctx = setContextOp(ctx, brq.ctx, "IDs")
	if err = brq.Select(benchresult.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (brq *BenchResultQuery) IDsX(ctx context.Context) []int {
	ids, err := brq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (brq *BenchResultQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, brq.ctx, "Count")
	if err := brq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, brq, querierCount[*BenchResultQuery](), brq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (brq *BenchResultQuery) CountX(ctx context.Context) int {
	count, err := brq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (brq *BenchResultQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, brq.ctx, "Exist")
	switch _, err := brq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (brq *BenchResultQuery) ExistX(ctx context.Context) bool {
	exist, err := brq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BenchResultQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (brq *BenchResultQuery) Clone() *BenchResultQuery {
	if brq == nil {
		return nil
	}
	return &BenchResultQuery{
		config:     brq.config,
		ctx:        brq.ctx.Clone(),
		order:      append([]OrderFunc{}, brq.order...),
		inters:     append([]Interceptor{}, brq.inters...),
		predicates: append([]predicate.BenchResult{}, brq.predicates...),
		// clone intermediate query.
		sql:  brq.sql.Clone(),
		path: brq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BenchResult.Query().
//		GroupBy(benchresult.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (brq *BenchResultQuery) GroupBy(field string, fields ...string) *BenchResultGroupBy {
	brq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BenchResultGroupBy{build: brq}
	grbuild.flds = &brq.ctx.Fields
	grbuild.label = benchresult.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"Name,omitempty"`
//	}
//
//	client.BenchResult.Query().
//		Select(benchresult.FieldName).
//		Scan(ctx, &v)
func (brq *BenchResultQuery) Select(fields ...string) *BenchResultSelect {
	brq.ctx.Fields = append(brq.ctx.Fields, fields...)
	sbuild := &BenchResultSelect{BenchResultQuery: brq}
	sbuild.label = benchresult.Label
	sbuild.flds, sbuild.scan = &brq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BenchResultSelect configured with the given aggregations.
func (brq *BenchResultQuery) Aggregate(fns ...AggregateFunc) *BenchResultSelect {
	return brq.Select().Aggregate(fns...)
}

func (brq *BenchResultQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range brq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, brq); err != nil {
				return err
			}
		}
	}
	for _, f := range brq.ctx.Fields {
		if !benchresult.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if brq.path != nil {
		prev, err := brq.path(ctx)
		if err != nil {
			return err
		}
		brq.sql = prev
	}
	return nil
}

func (brq *BenchResultQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BenchResult, error) {
	var (
		nodes   = []*BenchResult{}
		withFKs = brq.withFKs
		_spec   = brq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, benchresult.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BenchResult).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BenchResult{config: brq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(brq.modifiers) > 0 {
		_spec.Modifiers = brq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, brq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range brq.loadTotal {
		if err := brq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (brq *BenchResultQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := brq.querySpec()
	if len(brq.modifiers) > 0 {
		_spec.Modifiers = brq.modifiers
	}
	_spec.Node.Columns = brq.ctx.Fields
	if len(brq.ctx.Fields) > 0 {
		_spec.Unique = brq.ctx.Unique != nil && *brq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, brq.driver, _spec)
}

func (brq *BenchResultQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(benchresult.Table, benchresult.Columns, sqlgraph.NewFieldSpec(benchresult.FieldID, field.TypeInt))
	_spec.From = brq.sql
	if unique := brq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if brq.path != nil {
		_spec.Unique = true
	}
	if fields := brq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, benchresult.FieldID)
		for i := range fields {
			if fields[i] != benchresult.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := brq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := brq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := brq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := brq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (brq *BenchResultQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(brq.driver.Dialect())
	t1 := builder.Table(benchresult.Table)
	columns := brq.ctx.Fields
	if len(columns) == 0 {
		columns = benchresult.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if brq.sql != nil {
		selector = brq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if brq.ctx.Unique != nil && *brq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range brq.predicates {
		p(selector)
	}
	for _, p := range brq.order {
		p(selector)
	}
	if offset := brq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := brq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BenchResultGroupBy is the group-by builder for BenchResult entities.
type BenchResultGroupBy struct {
	selector
	build *BenchResultQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (brgb *BenchResultGroupBy) Aggregate(fns ...AggregateFunc) *BenchResultGroupBy {
	brgb.fns = append(brgb.fns, fns...)
	return brgb
}

// Scan applies the selector query and scans the result into the given value.
func (brgb *BenchResultGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, brgb.build.ctx, "GroupBy")
	if err := brgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BenchResultQuery, *BenchResultGroupBy](ctx, brgb.build, brgb, brgb.build.inters, v)
}

func (brgb *BenchResultGroupBy) sqlScan(ctx context.Context, root *BenchResultQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(brgb.fns))
	for _, fn := range brgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*brgb.flds)+len(brgb.fns))
		for _, f := range *brgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*brgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := brgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BenchResultSelect is the builder for selecting fields of BenchResult entities.
type BenchResultSelect struct {
	*BenchResultQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (brs *BenchResultSelect) Aggregate(fns ...AggregateFunc) *BenchResultSelect {
	brs.fns = append(brs.fns, fns...)
	return brs
}

// Scan applies the selector query and scans the result into the given value.
func (brs *BenchResultSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, brs.ctx, "Select")
	if err := brs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BenchResultQuery, *BenchResultSelect](ctx, brs.BenchResultQuery, brs, brs.inters, v)
}

func (brs *BenchResultSelect) sqlScan(ctx context.Context, root *BenchResultQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(brs.fns))
	for _, fn := range brs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*brs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := brs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
