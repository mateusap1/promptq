// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/mateusap1/promptq/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/mateusap1/promptq/ent/promptrequest"
	"github.com/mateusap1/promptq/ent/promptresponse"
	"github.com/mateusap1/promptq/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// PromptRequest is the client for interacting with the PromptRequest builders.
	PromptRequest *PromptRequestClient
	// PromptResponse is the client for interacting with the PromptResponse builders.
	PromptResponse *PromptResponseClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.PromptRequest = NewPromptRequestClient(c.config)
	c.PromptResponse = NewPromptResponseClient(c.config)
	c.User = NewUserClient(c.config)
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

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

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

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		PromptRequest:  NewPromptRequestClient(cfg),
		PromptResponse: NewPromptResponseClient(cfg),
		User:           NewUserClient(cfg),
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
		ctx:            ctx,
		config:         cfg,
		PromptRequest:  NewPromptRequestClient(cfg),
		PromptResponse: NewPromptResponseClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		PromptRequest.
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
	c.PromptRequest.Use(hooks...)
	c.PromptResponse.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.PromptRequest.Intercept(interceptors...)
	c.PromptResponse.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PromptRequestMutation:
		return c.PromptRequest.mutate(ctx, m)
	case *PromptResponseMutation:
		return c.PromptResponse.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PromptRequestClient is a client for the PromptRequest schema.
type PromptRequestClient struct {
	config
}

// NewPromptRequestClient returns a client for the PromptRequest from the given config.
func NewPromptRequestClient(c config) *PromptRequestClient {
	return &PromptRequestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `promptrequest.Hooks(f(g(h())))`.
func (c *PromptRequestClient) Use(hooks ...Hook) {
	c.hooks.PromptRequest = append(c.hooks.PromptRequest, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `promptrequest.Intercept(f(g(h())))`.
func (c *PromptRequestClient) Intercept(interceptors ...Interceptor) {
	c.inters.PromptRequest = append(c.inters.PromptRequest, interceptors...)
}

// Create returns a builder for creating a PromptRequest entity.
func (c *PromptRequestClient) Create() *PromptRequestCreate {
	mutation := newPromptRequestMutation(c.config, OpCreate)
	return &PromptRequestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PromptRequest entities.
func (c *PromptRequestClient) CreateBulk(builders ...*PromptRequestCreate) *PromptRequestCreateBulk {
	return &PromptRequestCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PromptRequestClient) MapCreateBulk(slice any, setFunc func(*PromptRequestCreate, int)) *PromptRequestCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PromptRequestCreateBulk{err: fmt.Errorf("calling to PromptRequestClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PromptRequestCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PromptRequestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PromptRequest.
func (c *PromptRequestClient) Update() *PromptRequestUpdate {
	mutation := newPromptRequestMutation(c.config, OpUpdate)
	return &PromptRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PromptRequestClient) UpdateOne(pr *PromptRequest) *PromptRequestUpdateOne {
	mutation := newPromptRequestMutation(c.config, OpUpdateOne, withPromptRequest(pr))
	return &PromptRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PromptRequestClient) UpdateOneID(id int) *PromptRequestUpdateOne {
	mutation := newPromptRequestMutation(c.config, OpUpdateOne, withPromptRequestID(id))
	return &PromptRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PromptRequest.
func (c *PromptRequestClient) Delete() *PromptRequestDelete {
	mutation := newPromptRequestMutation(c.config, OpDelete)
	return &PromptRequestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PromptRequestClient) DeleteOne(pr *PromptRequest) *PromptRequestDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PromptRequestClient) DeleteOneID(id int) *PromptRequestDeleteOne {
	builder := c.Delete().Where(promptrequest.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PromptRequestDeleteOne{builder}
}

// Query returns a query builder for PromptRequest.
func (c *PromptRequestClient) Query() *PromptRequestQuery {
	return &PromptRequestQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePromptRequest},
		inters: c.Interceptors(),
	}
}

// Get returns a PromptRequest entity by its id.
func (c *PromptRequestClient) Get(ctx context.Context, id int) (*PromptRequest, error) {
	return c.Query().Where(promptrequest.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PromptRequestClient) GetX(ctx context.Context, id int) *PromptRequest {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPromptResponse queries the prompt_response edge of a PromptRequest.
func (c *PromptRequestClient) QueryPromptResponse(pr *PromptRequest) *PromptResponseQuery {
	query := (&PromptResponseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(promptrequest.Table, promptrequest.FieldID, id),
			sqlgraph.To(promptresponse.Table, promptresponse.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, promptrequest.PromptResponseTable, promptrequest.PromptResponseColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PromptRequestClient) Hooks() []Hook {
	return c.hooks.PromptRequest
}

// Interceptors returns the client interceptors.
func (c *PromptRequestClient) Interceptors() []Interceptor {
	return c.inters.PromptRequest
}

func (c *PromptRequestClient) mutate(ctx context.Context, m *PromptRequestMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PromptRequestCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PromptRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PromptRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PromptRequestDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PromptRequest mutation op: %q", m.Op())
	}
}

// PromptResponseClient is a client for the PromptResponse schema.
type PromptResponseClient struct {
	config
}

// NewPromptResponseClient returns a client for the PromptResponse from the given config.
func NewPromptResponseClient(c config) *PromptResponseClient {
	return &PromptResponseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `promptresponse.Hooks(f(g(h())))`.
func (c *PromptResponseClient) Use(hooks ...Hook) {
	c.hooks.PromptResponse = append(c.hooks.PromptResponse, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `promptresponse.Intercept(f(g(h())))`.
func (c *PromptResponseClient) Intercept(interceptors ...Interceptor) {
	c.inters.PromptResponse = append(c.inters.PromptResponse, interceptors...)
}

// Create returns a builder for creating a PromptResponse entity.
func (c *PromptResponseClient) Create() *PromptResponseCreate {
	mutation := newPromptResponseMutation(c.config, OpCreate)
	return &PromptResponseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PromptResponse entities.
func (c *PromptResponseClient) CreateBulk(builders ...*PromptResponseCreate) *PromptResponseCreateBulk {
	return &PromptResponseCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PromptResponseClient) MapCreateBulk(slice any, setFunc func(*PromptResponseCreate, int)) *PromptResponseCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PromptResponseCreateBulk{err: fmt.Errorf("calling to PromptResponseClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PromptResponseCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PromptResponseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PromptResponse.
func (c *PromptResponseClient) Update() *PromptResponseUpdate {
	mutation := newPromptResponseMutation(c.config, OpUpdate)
	return &PromptResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PromptResponseClient) UpdateOne(pr *PromptResponse) *PromptResponseUpdateOne {
	mutation := newPromptResponseMutation(c.config, OpUpdateOne, withPromptResponse(pr))
	return &PromptResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PromptResponseClient) UpdateOneID(id int) *PromptResponseUpdateOne {
	mutation := newPromptResponseMutation(c.config, OpUpdateOne, withPromptResponseID(id))
	return &PromptResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PromptResponse.
func (c *PromptResponseClient) Delete() *PromptResponseDelete {
	mutation := newPromptResponseMutation(c.config, OpDelete)
	return &PromptResponseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PromptResponseClient) DeleteOne(pr *PromptResponse) *PromptResponseDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PromptResponseClient) DeleteOneID(id int) *PromptResponseDeleteOne {
	builder := c.Delete().Where(promptresponse.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PromptResponseDeleteOne{builder}
}

// Query returns a query builder for PromptResponse.
func (c *PromptResponseClient) Query() *PromptResponseQuery {
	return &PromptResponseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePromptResponse},
		inters: c.Interceptors(),
	}
}

// Get returns a PromptResponse entity by its id.
func (c *PromptResponseClient) Get(ctx context.Context, id int) (*PromptResponse, error) {
	return c.Query().Where(promptresponse.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PromptResponseClient) GetX(ctx context.Context, id int) *PromptResponse {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPromptRequest queries the prompt_request edge of a PromptResponse.
func (c *PromptResponseClient) QueryPromptRequest(pr *PromptResponse) *PromptRequestQuery {
	query := (&PromptRequestClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(promptresponse.Table, promptresponse.FieldID, id),
			sqlgraph.To(promptrequest.Table, promptrequest.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, promptresponse.PromptRequestTable, promptresponse.PromptRequestColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PromptResponseClient) Hooks() []Hook {
	return c.hooks.PromptResponse
}

// Interceptors returns the client interceptors.
func (c *PromptResponseClient) Interceptors() []Interceptor {
	return c.inters.PromptResponse
}

func (c *PromptResponseClient) mutate(ctx context.Context, m *PromptResponseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PromptResponseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PromptResponseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PromptResponseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PromptResponseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PromptResponse mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		PromptRequest, PromptResponse, User []ent.Hook
	}
	inters struct {
		PromptRequest, PromptResponse, User []ent.Interceptor
	}
)
