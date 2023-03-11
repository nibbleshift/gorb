// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/nibbleshift/gorb/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/nibbleshift/gorb/ent/bench"
	"github.com/nibbleshift/gorb/ent/benchresult"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Bench is the client for interacting with the Bench builders.
	Bench *BenchClient
	// BenchResult is the client for interacting with the BenchResult builders.
	BenchResult *BenchResultClient
	// additional fields for node api
	tables tables
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Bench = NewBenchClient(c.config)
	c.BenchResult = NewBenchResultClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Bench:       NewBenchClient(cfg),
		BenchResult: NewBenchResultClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		Bench:       NewBenchClient(cfg),
		BenchResult: NewBenchResultClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Bench.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Bench.Use(hooks...)
	c.BenchResult.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Bench.Intercept(interceptors...)
	c.BenchResult.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BenchMutation:
		return c.Bench.mutate(ctx, m)
	case *BenchResultMutation:
		return c.BenchResult.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// BenchClient is a client for the Bench schema.
type BenchClient struct {
	config
}

// NewBenchClient returns a client for the Bench from the given config.
func NewBenchClient(c config) *BenchClient {
	return &BenchClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bench.Hooks(f(g(h())))`.
func (c *BenchClient) Use(hooks ...Hook) {
	c.hooks.Bench = append(c.hooks.Bench, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `bench.Intercept(f(g(h())))`.
func (c *BenchClient) Intercept(interceptors ...Interceptor) {
	c.inters.Bench = append(c.inters.Bench, interceptors...)
}

// Create returns a builder for creating a Bench entity.
func (c *BenchClient) Create() *BenchCreate {
	mutation := newBenchMutation(c.config, OpCreate)
	return &BenchCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bench entities.
func (c *BenchClient) CreateBulk(builders ...*BenchCreate) *BenchCreateBulk {
	return &BenchCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bench.
func (c *BenchClient) Update() *BenchUpdate {
	mutation := newBenchMutation(c.config, OpUpdate)
	return &BenchUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BenchClient) UpdateOne(b *Bench) *BenchUpdateOne {
	mutation := newBenchMutation(c.config, OpUpdateOne, withBench(b))
	return &BenchUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BenchClient) UpdateOneID(id int) *BenchUpdateOne {
	mutation := newBenchMutation(c.config, OpUpdateOne, withBenchID(id))
	return &BenchUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bench.
func (c *BenchClient) Delete() *BenchDelete {
	mutation := newBenchMutation(c.config, OpDelete)
	return &BenchDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BenchClient) DeleteOne(b *Bench) *BenchDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BenchClient) DeleteOneID(id int) *BenchDeleteOne {
	builder := c.Delete().Where(bench.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BenchDeleteOne{builder}
}

// Query returns a query builder for Bench.
func (c *BenchClient) Query() *BenchQuery {
	return &BenchQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBench},
		inters: c.Interceptors(),
	}
}

// Get returns a Bench entity by its id.
func (c *BenchClient) Get(ctx context.Context, id int) (*Bench, error) {
	return c.Query().Where(bench.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BenchClient) GetX(ctx context.Context, id int) *Bench {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryResults queries the results edge of a Bench.
func (c *BenchClient) QueryResults(b *Bench) *BenchResultQuery {
	query := (&BenchResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bench.Table, bench.FieldID, id),
			sqlgraph.To(benchresult.Table, benchresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bench.ResultsTable, bench.ResultsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BenchClient) Hooks() []Hook {
	return c.hooks.Bench
}

// Interceptors returns the client interceptors.
func (c *BenchClient) Interceptors() []Interceptor {
	return c.inters.Bench
}

func (c *BenchClient) mutate(ctx context.Context, m *BenchMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BenchCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BenchUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BenchUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BenchDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Bench mutation op: %q", m.Op())
	}
}

// BenchResultClient is a client for the BenchResult schema.
type BenchResultClient struct {
	config
}

// NewBenchResultClient returns a client for the BenchResult from the given config.
func NewBenchResultClient(c config) *BenchResultClient {
	return &BenchResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `benchresult.Hooks(f(g(h())))`.
func (c *BenchResultClient) Use(hooks ...Hook) {
	c.hooks.BenchResult = append(c.hooks.BenchResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `benchresult.Intercept(f(g(h())))`.
func (c *BenchResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.BenchResult = append(c.inters.BenchResult, interceptors...)
}

// Create returns a builder for creating a BenchResult entity.
func (c *BenchResultClient) Create() *BenchResultCreate {
	mutation := newBenchResultMutation(c.config, OpCreate)
	return &BenchResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BenchResult entities.
func (c *BenchResultClient) CreateBulk(builders ...*BenchResultCreate) *BenchResultCreateBulk {
	return &BenchResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BenchResult.
func (c *BenchResultClient) Update() *BenchResultUpdate {
	mutation := newBenchResultMutation(c.config, OpUpdate)
	return &BenchResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BenchResultClient) UpdateOne(br *BenchResult) *BenchResultUpdateOne {
	mutation := newBenchResultMutation(c.config, OpUpdateOne, withBenchResult(br))
	return &BenchResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BenchResultClient) UpdateOneID(id int) *BenchResultUpdateOne {
	mutation := newBenchResultMutation(c.config, OpUpdateOne, withBenchResultID(id))
	return &BenchResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BenchResult.
func (c *BenchResultClient) Delete() *BenchResultDelete {
	mutation := newBenchResultMutation(c.config, OpDelete)
	return &BenchResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BenchResultClient) DeleteOne(br *BenchResult) *BenchResultDeleteOne {
	return c.DeleteOneID(br.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BenchResultClient) DeleteOneID(id int) *BenchResultDeleteOne {
	builder := c.Delete().Where(benchresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BenchResultDeleteOne{builder}
}

// Query returns a query builder for BenchResult.
func (c *BenchResultClient) Query() *BenchResultQuery {
	return &BenchResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBenchResult},
		inters: c.Interceptors(),
	}
}

// Get returns a BenchResult entity by its id.
func (c *BenchResultClient) Get(ctx context.Context, id int) (*BenchResult, error) {
	return c.Query().Where(benchresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BenchResultClient) GetX(ctx context.Context, id int) *BenchResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *BenchResultClient) Hooks() []Hook {
	return c.hooks.BenchResult
}

// Interceptors returns the client interceptors.
func (c *BenchResultClient) Interceptors() []Interceptor {
	return c.inters.BenchResult
}

func (c *BenchResultClient) mutate(ctx context.Context, m *BenchResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BenchResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BenchResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BenchResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BenchResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown BenchResult mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Bench, BenchResult []ent.Hook
	}
	inters struct {
		Bench, BenchResult []ent.Interceptor
	}
)