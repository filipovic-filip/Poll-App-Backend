// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"filip.filipovic/polling-app/model/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"filip.filipovic/polling-app/model/ent/poll"
	"filip.filipovic/polling-app/model/ent/polloption"
	"filip.filipovic/polling-app/model/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Poll is the client for interacting with the Poll builders.
	Poll *PollClient
	// PollOption is the client for interacting with the PollOption builders.
	PollOption *PollOptionClient
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
	c.Poll = NewPollClient(c.config)
	c.PollOption = NewPollOptionClient(c.config)
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
		ctx:        ctx,
		config:     cfg,
		Poll:       NewPollClient(cfg),
		PollOption: NewPollOptionClient(cfg),
		User:       NewUserClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Poll:       NewPollClient(cfg),
		PollOption: NewPollOptionClient(cfg),
		User:       NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Poll.
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
	c.Poll.Use(hooks...)
	c.PollOption.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Poll.Intercept(interceptors...)
	c.PollOption.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PollMutation:
		return c.Poll.mutate(ctx, m)
	case *PollOptionMutation:
		return c.PollOption.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PollClient is a client for the Poll schema.
type PollClient struct {
	config
}

// NewPollClient returns a client for the Poll from the given config.
func NewPollClient(c config) *PollClient {
	return &PollClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `poll.Hooks(f(g(h())))`.
func (c *PollClient) Use(hooks ...Hook) {
	c.hooks.Poll = append(c.hooks.Poll, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `poll.Intercept(f(g(h())))`.
func (c *PollClient) Intercept(interceptors ...Interceptor) {
	c.inters.Poll = append(c.inters.Poll, interceptors...)
}

// Create returns a builder for creating a Poll entity.
func (c *PollClient) Create() *PollCreate {
	mutation := newPollMutation(c.config, OpCreate)
	return &PollCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Poll entities.
func (c *PollClient) CreateBulk(builders ...*PollCreate) *PollCreateBulk {
	return &PollCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PollClient) MapCreateBulk(slice any, setFunc func(*PollCreate, int)) *PollCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PollCreateBulk{err: fmt.Errorf("calling to PollClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PollCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PollCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Poll.
func (c *PollClient) Update() *PollUpdate {
	mutation := newPollMutation(c.config, OpUpdate)
	return &PollUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PollClient) UpdateOne(po *Poll) *PollUpdateOne {
	mutation := newPollMutation(c.config, OpUpdateOne, withPoll(po))
	return &PollUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PollClient) UpdateOneID(id int) *PollUpdateOne {
	mutation := newPollMutation(c.config, OpUpdateOne, withPollID(id))
	return &PollUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Poll.
func (c *PollClient) Delete() *PollDelete {
	mutation := newPollMutation(c.config, OpDelete)
	return &PollDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PollClient) DeleteOne(po *Poll) *PollDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PollClient) DeleteOneID(id int) *PollDeleteOne {
	builder := c.Delete().Where(poll.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PollDeleteOne{builder}
}

// Query returns a query builder for Poll.
func (c *PollClient) Query() *PollQuery {
	return &PollQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePoll},
		inters: c.Interceptors(),
	}
}

// Get returns a Poll entity by its id.
func (c *PollClient) Get(ctx context.Context, id int) (*Poll, error) {
	return c.Query().Where(poll.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PollClient) GetX(ctx context.Context, id int) *Poll {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPollOptions queries the poll_options edge of a Poll.
func (c *PollClient) QueryPollOptions(po *Poll) *PollOptionQuery {
	query := (&PollOptionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, id),
			sqlgraph.To(polloption.Table, polloption.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, poll.PollOptionsTable, poll.PollOptionsColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCreatedBy queries the created_by edge of a Poll.
func (c *PollClient) QueryCreatedBy(po *Poll) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(poll.Table, poll.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, poll.CreatedByTable, poll.CreatedByColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PollClient) Hooks() []Hook {
	return c.hooks.Poll
}

// Interceptors returns the client interceptors.
func (c *PollClient) Interceptors() []Interceptor {
	return c.inters.Poll
}

func (c *PollClient) mutate(ctx context.Context, m *PollMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PollCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PollUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PollUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PollDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Poll mutation op: %q", m.Op())
	}
}

// PollOptionClient is a client for the PollOption schema.
type PollOptionClient struct {
	config
}

// NewPollOptionClient returns a client for the PollOption from the given config.
func NewPollOptionClient(c config) *PollOptionClient {
	return &PollOptionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `polloption.Hooks(f(g(h())))`.
func (c *PollOptionClient) Use(hooks ...Hook) {
	c.hooks.PollOption = append(c.hooks.PollOption, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `polloption.Intercept(f(g(h())))`.
func (c *PollOptionClient) Intercept(interceptors ...Interceptor) {
	c.inters.PollOption = append(c.inters.PollOption, interceptors...)
}

// Create returns a builder for creating a PollOption entity.
func (c *PollOptionClient) Create() *PollOptionCreate {
	mutation := newPollOptionMutation(c.config, OpCreate)
	return &PollOptionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of PollOption entities.
func (c *PollOptionClient) CreateBulk(builders ...*PollOptionCreate) *PollOptionCreateBulk {
	return &PollOptionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *PollOptionClient) MapCreateBulk(slice any, setFunc func(*PollOptionCreate, int)) *PollOptionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &PollOptionCreateBulk{err: fmt.Errorf("calling to PollOptionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*PollOptionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &PollOptionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for PollOption.
func (c *PollOptionClient) Update() *PollOptionUpdate {
	mutation := newPollOptionMutation(c.config, OpUpdate)
	return &PollOptionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PollOptionClient) UpdateOne(po *PollOption) *PollOptionUpdateOne {
	mutation := newPollOptionMutation(c.config, OpUpdateOne, withPollOption(po))
	return &PollOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PollOptionClient) UpdateOneID(id int) *PollOptionUpdateOne {
	mutation := newPollOptionMutation(c.config, OpUpdateOne, withPollOptionID(id))
	return &PollOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for PollOption.
func (c *PollOptionClient) Delete() *PollOptionDelete {
	mutation := newPollOptionMutation(c.config, OpDelete)
	return &PollOptionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PollOptionClient) DeleteOne(po *PollOption) *PollOptionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PollOptionClient) DeleteOneID(id int) *PollOptionDeleteOne {
	builder := c.Delete().Where(polloption.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PollOptionDeleteOne{builder}
}

// Query returns a query builder for PollOption.
func (c *PollOptionClient) Query() *PollOptionQuery {
	return &PollOptionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePollOption},
		inters: c.Interceptors(),
	}
}

// Get returns a PollOption entity by its id.
func (c *PollOptionClient) Get(ctx context.Context, id int) (*PollOption, error) {
	return c.Query().Where(polloption.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PollOptionClient) GetX(ctx context.Context, id int) *PollOption {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUsersVoted queries the users_voted edge of a PollOption.
func (c *PollOptionClient) QueryUsersVoted(po *PollOption) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(polloption.Table, polloption.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, polloption.UsersVotedTable, polloption.UsersVotedPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPoll queries the poll edge of a PollOption.
func (c *PollOptionClient) QueryPoll(po *PollOption) *PollQuery {
	query := (&PollClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(polloption.Table, polloption.FieldID, id),
			sqlgraph.To(poll.Table, poll.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, polloption.PollTable, polloption.PollColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PollOptionClient) Hooks() []Hook {
	return c.hooks.PollOption
}

// Interceptors returns the client interceptors.
func (c *PollOptionClient) Interceptors() []Interceptor {
	return c.inters.PollOption
}

func (c *PollOptionClient) mutate(ctx context.Context, m *PollOptionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PollOptionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PollOptionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PollOptionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PollOptionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown PollOption mutation op: %q", m.Op())
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

// QueryCreatedPolls queries the created_polls edge of a User.
func (c *UserClient) QueryCreatedPolls(u *User) *PollQuery {
	query := (&PollClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(poll.Table, poll.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.CreatedPollsTable, user.CreatedPollsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryVotes queries the votes edge of a User.
func (c *UserClient) QueryVotes(u *User) *PollOptionQuery {
	query := (&PollOptionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(polloption.Table, polloption.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, user.VotesTable, user.VotesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
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
		Poll, PollOption, User []ent.Hook
	}
	inters struct {
		Poll, PollOption, User []ent.Interceptor
	}
)
