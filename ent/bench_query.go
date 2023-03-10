// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nibbleshift/gorb/ent/bench"
	"github.com/nibbleshift/gorb/ent/benchresult"
	"github.com/nibbleshift/gorb/ent/predicate"
)

// BenchQuery is the builder for querying Bench entities.
type BenchQuery struct {
	config
	ctx                  *QueryContext
	order                []OrderFunc
	inters               []Interceptor
	predicates           []predicate.Bench
	withBenchResult      *BenchResultQuery
	modifiers            []func(*sql.Selector)
	loadTotal            []func(context.Context, []*Bench) error
	withNamedBenchResult map[string]*BenchResultQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BenchQuery builder.
func (bq *BenchQuery) Where(ps ...predicate.Bench) *BenchQuery {
	bq.predicates = append(bq.predicates, ps...)
	return bq
}

// Limit the number of records to be returned by this query.
func (bq *BenchQuery) Limit(limit int) *BenchQuery {
	bq.ctx.Limit = &limit
	return bq
}

// Offset to start from.
func (bq *BenchQuery) Offset(offset int) *BenchQuery {
	bq.ctx.Offset = &offset
	return bq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bq *BenchQuery) Unique(unique bool) *BenchQuery {
	bq.ctx.Unique = &unique
	return bq
}

// Order specifies how the records should be ordered.
func (bq *BenchQuery) Order(o ...OrderFunc) *BenchQuery {
	bq.order = append(bq.order, o...)
	return bq
}

// QueryBenchResult chains the current query on the "bench_result" edge.
func (bq *BenchQuery) QueryBenchResult() *BenchResultQuery {
	query := (&BenchResultClient{config: bq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := bq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := bq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(bench.Table, bench.FieldID, selector),
			sqlgraph.To(benchresult.Table, benchresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bench.BenchResultTable, bench.BenchResultColumn),
		)
		fromU = sqlgraph.SetNeighbors(bq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Bench entity from the query.
// Returns a *NotFoundError when no Bench was found.
func (bq *BenchQuery) First(ctx context.Context) (*Bench, error) {
	nodes, err := bq.Limit(1).All(setContextOp(ctx, bq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{bench.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bq *BenchQuery) FirstX(ctx context.Context) *Bench {
	node, err := bq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Bench ID from the query.
// Returns a *NotFoundError when no Bench ID was found.
func (bq *BenchQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(1).IDs(setContextOp(ctx, bq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{bench.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bq *BenchQuery) FirstIDX(ctx context.Context) int {
	id, err := bq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Bench entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Bench entity is found.
// Returns a *NotFoundError when no Bench entities are found.
func (bq *BenchQuery) Only(ctx context.Context) (*Bench, error) {
	nodes, err := bq.Limit(2).All(setContextOp(ctx, bq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{bench.Label}
	default:
		return nil, &NotSingularError{bench.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bq *BenchQuery) OnlyX(ctx context.Context) *Bench {
	node, err := bq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Bench ID in the query.
// Returns a *NotSingularError when more than one Bench ID is found.
// Returns a *NotFoundError when no entities are found.
func (bq *BenchQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bq.Limit(2).IDs(setContextOp(ctx, bq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{bench.Label}
	default:
		err = &NotSingularError{bench.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bq *BenchQuery) OnlyIDX(ctx context.Context) int {
	id, err := bq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Benches.
func (bq *BenchQuery) All(ctx context.Context) ([]*Bench, error) {
	ctx = setContextOp(ctx, bq.ctx, "All")
	if err := bq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Bench, *BenchQuery]()
	return withInterceptors[[]*Bench](ctx, bq, qr, bq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bq *BenchQuery) AllX(ctx context.Context) []*Bench {
	nodes, err := bq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Bench IDs.
func (bq *BenchQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bq.ctx.Unique == nil && bq.path != nil {
		bq.Unique(true)
	}
	ctx = setContextOp(ctx, bq.ctx, "IDs")
	if err = bq.Select(bench.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bq *BenchQuery) IDsX(ctx context.Context) []int {
	ids, err := bq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bq *BenchQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bq.ctx, "Count")
	if err := bq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bq, querierCount[*BenchQuery](), bq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bq *BenchQuery) CountX(ctx context.Context) int {
	count, err := bq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bq *BenchQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bq.ctx, "Exist")
	switch _, err := bq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bq *BenchQuery) ExistX(ctx context.Context) bool {
	exist, err := bq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BenchQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bq *BenchQuery) Clone() *BenchQuery {
	if bq == nil {
		return nil
	}
	return &BenchQuery{
		config:          bq.config,
		ctx:             bq.ctx.Clone(),
		order:           append([]OrderFunc{}, bq.order...),
		inters:          append([]Interceptor{}, bq.inters...),
		predicates:      append([]predicate.Bench{}, bq.predicates...),
		withBenchResult: bq.withBenchResult.Clone(),
		// clone intermediate query.
		sql:  bq.sql.Clone(),
		path: bq.path,
	}
}

// WithBenchResult tells the query-builder to eager-load the nodes that are connected to
// the "bench_result" edge. The optional arguments are used to configure the query builder of the edge.
func (bq *BenchQuery) WithBenchResult(opts ...func(*BenchResultQuery)) *BenchQuery {
	query := (&BenchResultClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	bq.withBenchResult = query
	return bq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Os string `json:"os,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Bench.Query().
//		GroupBy(bench.FieldOs).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bq *BenchQuery) GroupBy(field string, fields ...string) *BenchGroupBy {
	bq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BenchGroupBy{build: bq}
	grbuild.flds = &bq.ctx.Fields
	grbuild.label = bench.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Os string `json:"os,omitempty"`
//	}
//
//	client.Bench.Query().
//		Select(bench.FieldOs).
//		Scan(ctx, &v)
func (bq *BenchQuery) Select(fields ...string) *BenchSelect {
	bq.ctx.Fields = append(bq.ctx.Fields, fields...)
	sbuild := &BenchSelect{BenchQuery: bq}
	sbuild.label = bench.Label
	sbuild.flds, sbuild.scan = &bq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BenchSelect configured with the given aggregations.
func (bq *BenchQuery) Aggregate(fns ...AggregateFunc) *BenchSelect {
	return bq.Select().Aggregate(fns...)
}

func (bq *BenchQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bq); err != nil {
				return err
			}
		}
	}
	for _, f := range bq.ctx.Fields {
		if !bench.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bq.path != nil {
		prev, err := bq.path(ctx)
		if err != nil {
			return err
		}
		bq.sql = prev
	}
	return nil
}

func (bq *BenchQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Bench, error) {
	var (
		nodes       = []*Bench{}
		_spec       = bq.querySpec()
		loadedTypes = [1]bool{
			bq.withBenchResult != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Bench).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Bench{config: bq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := bq.withBenchResult; query != nil {
		if err := bq.loadBenchResult(ctx, query, nodes,
			func(n *Bench) { n.Edges.BenchResult = []*BenchResult{} },
			func(n *Bench, e *BenchResult) { n.Edges.BenchResult = append(n.Edges.BenchResult, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range bq.withNamedBenchResult {
		if err := bq.loadBenchResult(ctx, query, nodes,
			func(n *Bench) { n.appendNamedBenchResult(name) },
			func(n *Bench, e *BenchResult) { n.appendNamedBenchResult(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range bq.loadTotal {
		if err := bq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (bq *BenchQuery) loadBenchResult(ctx context.Context, query *BenchResultQuery, nodes []*Bench, init func(*Bench), assign func(*Bench, *BenchResult)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Bench)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.BenchResult(func(s *sql.Selector) {
		s.Where(sql.InValues(bench.BenchResultColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.bench_bench_result
		if fk == nil {
			return fmt.Errorf(`foreign-key "bench_bench_result" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "bench_bench_result" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (bq *BenchQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bq.querySpec()
	if len(bq.modifiers) > 0 {
		_spec.Modifiers = bq.modifiers
	}
	_spec.Node.Columns = bq.ctx.Fields
	if len(bq.ctx.Fields) > 0 {
		_spec.Unique = bq.ctx.Unique != nil && *bq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bq.driver, _spec)
}

func (bq *BenchQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(bench.Table, bench.Columns, sqlgraph.NewFieldSpec(bench.FieldID, field.TypeInt))
	_spec.From = bq.sql
	if unique := bq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bq.path != nil {
		_spec.Unique = true
	}
	if fields := bq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, bench.FieldID)
		for i := range fields {
			if fields[i] != bench.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bq *BenchQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bq.driver.Dialect())
	t1 := builder.Table(bench.Table)
	columns := bq.ctx.Fields
	if len(columns) == 0 {
		columns = bench.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bq.sql != nil {
		selector = bq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bq.ctx.Unique != nil && *bq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bq.predicates {
		p(selector)
	}
	for _, p := range bq.order {
		p(selector)
	}
	if offset := bq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedBenchResult tells the query-builder to eager-load the nodes that are connected to the "bench_result"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (bq *BenchQuery) WithNamedBenchResult(name string, opts ...func(*BenchResultQuery)) *BenchQuery {
	query := (&BenchResultClient{config: bq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if bq.withNamedBenchResult == nil {
		bq.withNamedBenchResult = make(map[string]*BenchResultQuery)
	}
	bq.withNamedBenchResult[name] = query
	return bq
}

// BenchGroupBy is the group-by builder for Bench entities.
type BenchGroupBy struct {
	selector
	build *BenchQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bgb *BenchGroupBy) Aggregate(fns ...AggregateFunc) *BenchGroupBy {
	bgb.fns = append(bgb.fns, fns...)
	return bgb
}

// Scan applies the selector query and scans the result into the given value.
func (bgb *BenchGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bgb.build.ctx, "GroupBy")
	if err := bgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BenchQuery, *BenchGroupBy](ctx, bgb.build, bgb, bgb.build.inters, v)
}

func (bgb *BenchGroupBy) sqlScan(ctx context.Context, root *BenchQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bgb.fns))
	for _, fn := range bgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bgb.flds)+len(bgb.fns))
		for _, f := range *bgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BenchSelect is the builder for selecting fields of Bench entities.
type BenchSelect struct {
	*BenchQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bs *BenchSelect) Aggregate(fns ...AggregateFunc) *BenchSelect {
	bs.fns = append(bs.fns, fns...)
	return bs
}

// Scan applies the selector query and scans the result into the given value.
func (bs *BenchSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bs.ctx, "Select")
	if err := bs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BenchQuery, *BenchSelect](ctx, bs.BenchQuery, bs, bs.inters, v)
}

func (bs *BenchSelect) sqlScan(ctx context.Context, root *BenchQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bs.fns))
	for _, fn := range bs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
