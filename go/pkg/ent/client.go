// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/k3forx/coffee_memo/pkg/ent/migrate"

	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeeshop"
	"github.com/k3forx/coffee_memo/pkg/ent/driprecipe"
	"github.com/k3forx/coffee_memo/pkg/ent/goosedbversion"
	"github.com/k3forx/coffee_memo/pkg/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// CoffeeBean is the client for interacting with the CoffeeBean builders.
	CoffeeBean *CoffeeBeanClient
	// CoffeeShop is the client for interacting with the CoffeeShop builders.
	CoffeeShop *CoffeeShopClient
	// DripRecipe is the client for interacting with the DripRecipe builders.
	DripRecipe *DripRecipeClient
	// GooseDbVersion is the client for interacting with the GooseDbVersion builders.
	GooseDbVersion *GooseDbVersionClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.CoffeeBean = NewCoffeeBeanClient(c.config)
	c.CoffeeShop = NewCoffeeShopClient(c.config)
	c.DripRecipe = NewDripRecipeClient(c.config)
	c.GooseDbVersion = NewGooseDbVersionClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		CoffeeBean:     NewCoffeeBeanClient(cfg),
		CoffeeShop:     NewCoffeeShopClient(cfg),
		DripRecipe:     NewDripRecipeClient(cfg),
		GooseDbVersion: NewGooseDbVersionClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		CoffeeBean:     NewCoffeeBeanClient(cfg),
		CoffeeShop:     NewCoffeeShopClient(cfg),
		DripRecipe:     NewDripRecipeClient(cfg),
		GooseDbVersion: NewGooseDbVersionClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		CoffeeBean.
//		Query().
//		Count(ctx)
//
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
	c.CoffeeBean.Use(hooks...)
	c.CoffeeShop.Use(hooks...)
	c.DripRecipe.Use(hooks...)
	c.GooseDbVersion.Use(hooks...)
	c.User.Use(hooks...)
}

// CoffeeBeanClient is a client for the CoffeeBean schema.
type CoffeeBeanClient struct {
	config
}

// NewCoffeeBeanClient returns a client for the CoffeeBean from the given config.
func NewCoffeeBeanClient(c config) *CoffeeBeanClient {
	return &CoffeeBeanClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coffeebean.Hooks(f(g(h())))`.
func (c *CoffeeBeanClient) Use(hooks ...Hook) {
	c.hooks.CoffeeBean = append(c.hooks.CoffeeBean, hooks...)
}

// Create returns a create builder for CoffeeBean.
func (c *CoffeeBeanClient) Create() *CoffeeBeanCreate {
	mutation := newCoffeeBeanMutation(c.config, OpCreate)
	return &CoffeeBeanCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CoffeeBean entities.
func (c *CoffeeBeanClient) CreateBulk(builders ...*CoffeeBeanCreate) *CoffeeBeanCreateBulk {
	return &CoffeeBeanCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CoffeeBean.
func (c *CoffeeBeanClient) Update() *CoffeeBeanUpdate {
	mutation := newCoffeeBeanMutation(c.config, OpUpdate)
	return &CoffeeBeanUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoffeeBeanClient) UpdateOne(cb *CoffeeBean) *CoffeeBeanUpdateOne {
	mutation := newCoffeeBeanMutation(c.config, OpUpdateOne, withCoffeeBean(cb))
	return &CoffeeBeanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoffeeBeanClient) UpdateOneID(id int32) *CoffeeBeanUpdateOne {
	mutation := newCoffeeBeanMutation(c.config, OpUpdateOne, withCoffeeBeanID(id))
	return &CoffeeBeanUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoffeeBean.
func (c *CoffeeBeanClient) Delete() *CoffeeBeanDelete {
	mutation := newCoffeeBeanMutation(c.config, OpDelete)
	return &CoffeeBeanDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoffeeBeanClient) DeleteOne(cb *CoffeeBean) *CoffeeBeanDeleteOne {
	return c.DeleteOneID(cb.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoffeeBeanClient) DeleteOneID(id int32) *CoffeeBeanDeleteOne {
	builder := c.Delete().Where(coffeebean.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoffeeBeanDeleteOne{builder}
}

// Query returns a query builder for CoffeeBean.
func (c *CoffeeBeanClient) Query() *CoffeeBeanQuery {
	return &CoffeeBeanQuery{
		config: c.config,
	}
}

// Get returns a CoffeeBean entity by its id.
func (c *CoffeeBeanClient) Get(ctx context.Context, id int32) (*CoffeeBean, error) {
	return c.Query().Where(coffeebean.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoffeeBeanClient) GetX(ctx context.Context, id int32) *CoffeeBean {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CoffeeBeanClient) Hooks() []Hook {
	return c.hooks.CoffeeBean
}

// CoffeeShopClient is a client for the CoffeeShop schema.
type CoffeeShopClient struct {
	config
}

// NewCoffeeShopClient returns a client for the CoffeeShop from the given config.
func NewCoffeeShopClient(c config) *CoffeeShopClient {
	return &CoffeeShopClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `coffeeshop.Hooks(f(g(h())))`.
func (c *CoffeeShopClient) Use(hooks ...Hook) {
	c.hooks.CoffeeShop = append(c.hooks.CoffeeShop, hooks...)
}

// Create returns a create builder for CoffeeShop.
func (c *CoffeeShopClient) Create() *CoffeeShopCreate {
	mutation := newCoffeeShopMutation(c.config, OpCreate)
	return &CoffeeShopCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CoffeeShop entities.
func (c *CoffeeShopClient) CreateBulk(builders ...*CoffeeShopCreate) *CoffeeShopCreateBulk {
	return &CoffeeShopCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CoffeeShop.
func (c *CoffeeShopClient) Update() *CoffeeShopUpdate {
	mutation := newCoffeeShopMutation(c.config, OpUpdate)
	return &CoffeeShopUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CoffeeShopClient) UpdateOne(cs *CoffeeShop) *CoffeeShopUpdateOne {
	mutation := newCoffeeShopMutation(c.config, OpUpdateOne, withCoffeeShop(cs))
	return &CoffeeShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CoffeeShopClient) UpdateOneID(id int32) *CoffeeShopUpdateOne {
	mutation := newCoffeeShopMutation(c.config, OpUpdateOne, withCoffeeShopID(id))
	return &CoffeeShopUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CoffeeShop.
func (c *CoffeeShopClient) Delete() *CoffeeShopDelete {
	mutation := newCoffeeShopMutation(c.config, OpDelete)
	return &CoffeeShopDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CoffeeShopClient) DeleteOne(cs *CoffeeShop) *CoffeeShopDeleteOne {
	return c.DeleteOneID(cs.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CoffeeShopClient) DeleteOneID(id int32) *CoffeeShopDeleteOne {
	builder := c.Delete().Where(coffeeshop.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CoffeeShopDeleteOne{builder}
}

// Query returns a query builder for CoffeeShop.
func (c *CoffeeShopClient) Query() *CoffeeShopQuery {
	return &CoffeeShopQuery{
		config: c.config,
	}
}

// Get returns a CoffeeShop entity by its id.
func (c *CoffeeShopClient) Get(ctx context.Context, id int32) (*CoffeeShop, error) {
	return c.Query().Where(coffeeshop.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CoffeeShopClient) GetX(ctx context.Context, id int32) *CoffeeShop {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CoffeeShopClient) Hooks() []Hook {
	return c.hooks.CoffeeShop
}

// DripRecipeClient is a client for the DripRecipe schema.
type DripRecipeClient struct {
	config
}

// NewDripRecipeClient returns a client for the DripRecipe from the given config.
func NewDripRecipeClient(c config) *DripRecipeClient {
	return &DripRecipeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `driprecipe.Hooks(f(g(h())))`.
func (c *DripRecipeClient) Use(hooks ...Hook) {
	c.hooks.DripRecipe = append(c.hooks.DripRecipe, hooks...)
}

// Create returns a create builder for DripRecipe.
func (c *DripRecipeClient) Create() *DripRecipeCreate {
	mutation := newDripRecipeMutation(c.config, OpCreate)
	return &DripRecipeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of DripRecipe entities.
func (c *DripRecipeClient) CreateBulk(builders ...*DripRecipeCreate) *DripRecipeCreateBulk {
	return &DripRecipeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for DripRecipe.
func (c *DripRecipeClient) Update() *DripRecipeUpdate {
	mutation := newDripRecipeMutation(c.config, OpUpdate)
	return &DripRecipeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *DripRecipeClient) UpdateOne(dr *DripRecipe) *DripRecipeUpdateOne {
	mutation := newDripRecipeMutation(c.config, OpUpdateOne, withDripRecipe(dr))
	return &DripRecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *DripRecipeClient) UpdateOneID(id int32) *DripRecipeUpdateOne {
	mutation := newDripRecipeMutation(c.config, OpUpdateOne, withDripRecipeID(id))
	return &DripRecipeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for DripRecipe.
func (c *DripRecipeClient) Delete() *DripRecipeDelete {
	mutation := newDripRecipeMutation(c.config, OpDelete)
	return &DripRecipeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *DripRecipeClient) DeleteOne(dr *DripRecipe) *DripRecipeDeleteOne {
	return c.DeleteOneID(dr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *DripRecipeClient) DeleteOneID(id int32) *DripRecipeDeleteOne {
	builder := c.Delete().Where(driprecipe.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &DripRecipeDeleteOne{builder}
}

// Query returns a query builder for DripRecipe.
func (c *DripRecipeClient) Query() *DripRecipeQuery {
	return &DripRecipeQuery{
		config: c.config,
	}
}

// Get returns a DripRecipe entity by its id.
func (c *DripRecipeClient) Get(ctx context.Context, id int32) (*DripRecipe, error) {
	return c.Query().Where(driprecipe.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *DripRecipeClient) GetX(ctx context.Context, id int32) *DripRecipe {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *DripRecipeClient) Hooks() []Hook {
	return c.hooks.DripRecipe
}

// GooseDbVersionClient is a client for the GooseDbVersion schema.
type GooseDbVersionClient struct {
	config
}

// NewGooseDbVersionClient returns a client for the GooseDbVersion from the given config.
func NewGooseDbVersionClient(c config) *GooseDbVersionClient {
	return &GooseDbVersionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `goosedbversion.Hooks(f(g(h())))`.
func (c *GooseDbVersionClient) Use(hooks ...Hook) {
	c.hooks.GooseDbVersion = append(c.hooks.GooseDbVersion, hooks...)
}

// Create returns a create builder for GooseDbVersion.
func (c *GooseDbVersionClient) Create() *GooseDbVersionCreate {
	mutation := newGooseDbVersionMutation(c.config, OpCreate)
	return &GooseDbVersionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of GooseDbVersion entities.
func (c *GooseDbVersionClient) CreateBulk(builders ...*GooseDbVersionCreate) *GooseDbVersionCreateBulk {
	return &GooseDbVersionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for GooseDbVersion.
func (c *GooseDbVersionClient) Update() *GooseDbVersionUpdate {
	mutation := newGooseDbVersionMutation(c.config, OpUpdate)
	return &GooseDbVersionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *GooseDbVersionClient) UpdateOne(gdv *GooseDbVersion) *GooseDbVersionUpdateOne {
	mutation := newGooseDbVersionMutation(c.config, OpUpdateOne, withGooseDbVersion(gdv))
	return &GooseDbVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *GooseDbVersionClient) UpdateOneID(id uint64) *GooseDbVersionUpdateOne {
	mutation := newGooseDbVersionMutation(c.config, OpUpdateOne, withGooseDbVersionID(id))
	return &GooseDbVersionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for GooseDbVersion.
func (c *GooseDbVersionClient) Delete() *GooseDbVersionDelete {
	mutation := newGooseDbVersionMutation(c.config, OpDelete)
	return &GooseDbVersionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GooseDbVersionClient) DeleteOne(gdv *GooseDbVersion) *GooseDbVersionDeleteOne {
	return c.DeleteOneID(gdv.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GooseDbVersionClient) DeleteOneID(id uint64) *GooseDbVersionDeleteOne {
	builder := c.Delete().Where(goosedbversion.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &GooseDbVersionDeleteOne{builder}
}

// Query returns a query builder for GooseDbVersion.
func (c *GooseDbVersionClient) Query() *GooseDbVersionQuery {
	return &GooseDbVersionQuery{
		config: c.config,
	}
}

// Get returns a GooseDbVersion entity by its id.
func (c *GooseDbVersionClient) Get(ctx context.Context, id uint64) (*GooseDbVersion, error) {
	return c.Query().Where(goosedbversion.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GooseDbVersionClient) GetX(ctx context.Context, id uint64) *GooseDbVersion {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *GooseDbVersionClient) Hooks() []Hook {
	return c.hooks.GooseDbVersion
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

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
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
func (c *UserClient) UpdateOneID(id int32) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int32) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int32) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int32) *User {
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
