// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/ent/userbrewrecipe"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
)

// UserBrewRecipeQuery is the builder for querying UserBrewRecipe entities.
type UserBrewRecipeQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserBrewRecipe
	// eager-loading edges.
	withUserCoffeeBean *UserCoffeeBeanQuery
	withUser           *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserBrewRecipeQuery builder.
func (ubrq *UserBrewRecipeQuery) Where(ps ...predicate.UserBrewRecipe) *UserBrewRecipeQuery {
	ubrq.predicates = append(ubrq.predicates, ps...)
	return ubrq
}

// Limit adds a limit step to the query.
func (ubrq *UserBrewRecipeQuery) Limit(limit int) *UserBrewRecipeQuery {
	ubrq.limit = &limit
	return ubrq
}

// Offset adds an offset step to the query.
func (ubrq *UserBrewRecipeQuery) Offset(offset int) *UserBrewRecipeQuery {
	ubrq.offset = &offset
	return ubrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ubrq *UserBrewRecipeQuery) Unique(unique bool) *UserBrewRecipeQuery {
	ubrq.unique = &unique
	return ubrq
}

// Order adds an order step to the query.
func (ubrq *UserBrewRecipeQuery) Order(o ...OrderFunc) *UserBrewRecipeQuery {
	ubrq.order = append(ubrq.order, o...)
	return ubrq
}

// QueryUserCoffeeBean chains the current query on the "user_coffee_bean" edge.
func (ubrq *UserBrewRecipeQuery) QueryUserCoffeeBean() *UserCoffeeBeanQuery {
	query := &UserCoffeeBeanQuery{config: ubrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbrewrecipe.Table, userbrewrecipe.FieldID, selector),
			sqlgraph.To(usercoffeebean.Table, usercoffeebean.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userbrewrecipe.UserCoffeeBeanTable, userbrewrecipe.UserCoffeeBeanColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (ubrq *UserBrewRecipeQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: ubrq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ubrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ubrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(userbrewrecipe.Table, userbrewrecipe.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, userbrewrecipe.UserTable, userbrewrecipe.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ubrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserBrewRecipe entity from the query.
// Returns a *NotFoundError when no UserBrewRecipe was found.
func (ubrq *UserBrewRecipeQuery) First(ctx context.Context) (*UserBrewRecipe, error) {
	nodes, err := ubrq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userbrewrecipe.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) FirstX(ctx context.Context) *UserBrewRecipe {
	node, err := ubrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserBrewRecipe ID from the query.
// Returns a *NotFoundError when no UserBrewRecipe ID was found.
func (ubrq *UserBrewRecipeQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ubrq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userbrewrecipe.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) FirstIDX(ctx context.Context) int32 {
	id, err := ubrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserBrewRecipe entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserBrewRecipe entity is found.
// Returns a *NotFoundError when no UserBrewRecipe entities are found.
func (ubrq *UserBrewRecipeQuery) Only(ctx context.Context) (*UserBrewRecipe, error) {
	nodes, err := ubrq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userbrewrecipe.Label}
	default:
		return nil, &NotSingularError{userbrewrecipe.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) OnlyX(ctx context.Context) *UserBrewRecipe {
	node, err := ubrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserBrewRecipe ID in the query.
// Returns a *NotSingularError when more than one UserBrewRecipe ID is found.
// Returns a *NotFoundError when no entities are found.
func (ubrq *UserBrewRecipeQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = ubrq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userbrewrecipe.Label}
	default:
		err = &NotSingularError{userbrewrecipe.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := ubrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserBrewRecipes.
func (ubrq *UserBrewRecipeQuery) All(ctx context.Context) ([]*UserBrewRecipe, error) {
	if err := ubrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ubrq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) AllX(ctx context.Context) []*UserBrewRecipe {
	nodes, err := ubrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserBrewRecipe IDs.
func (ubrq *UserBrewRecipeQuery) IDs(ctx context.Context) ([]int32, error) {
	var ids []int32
	if err := ubrq.Select(userbrewrecipe.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) IDsX(ctx context.Context) []int32 {
	ids, err := ubrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ubrq *UserBrewRecipeQuery) Count(ctx context.Context) (int, error) {
	if err := ubrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ubrq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) CountX(ctx context.Context) int {
	count, err := ubrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ubrq *UserBrewRecipeQuery) Exist(ctx context.Context) (bool, error) {
	if err := ubrq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ubrq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ubrq *UserBrewRecipeQuery) ExistX(ctx context.Context) bool {
	exist, err := ubrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserBrewRecipeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ubrq *UserBrewRecipeQuery) Clone() *UserBrewRecipeQuery {
	if ubrq == nil {
		return nil
	}
	return &UserBrewRecipeQuery{
		config:             ubrq.config,
		limit:              ubrq.limit,
		offset:             ubrq.offset,
		order:              append([]OrderFunc{}, ubrq.order...),
		predicates:         append([]predicate.UserBrewRecipe{}, ubrq.predicates...),
		withUserCoffeeBean: ubrq.withUserCoffeeBean.Clone(),
		withUser:           ubrq.withUser.Clone(),
		// clone intermediate query.
		sql:    ubrq.sql.Clone(),
		path:   ubrq.path,
		unique: ubrq.unique,
	}
}

// WithUserCoffeeBean tells the query-builder to eager-load the nodes that are connected to
// the "user_coffee_bean" edge. The optional arguments are used to configure the query builder of the edge.
func (ubrq *UserBrewRecipeQuery) WithUserCoffeeBean(opts ...func(*UserCoffeeBeanQuery)) *UserBrewRecipeQuery {
	query := &UserCoffeeBeanQuery{config: ubrq.config}
	for _, opt := range opts {
		opt(query)
	}
	ubrq.withUserCoffeeBean = query
	return ubrq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ubrq *UserBrewRecipeQuery) WithUser(opts ...func(*UserQuery)) *UserBrewRecipeQuery {
	query := &UserQuery{config: ubrq.config}
	for _, opt := range opts {
		opt(query)
	}
	ubrq.withUser = query
	return ubrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Status int32 `json:"status,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserBrewRecipe.Query().
//		GroupBy(userbrewrecipe.FieldStatus).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (ubrq *UserBrewRecipeQuery) GroupBy(field string, fields ...string) *UserBrewRecipeGroupBy {
	grbuild := &UserBrewRecipeGroupBy{config: ubrq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ubrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ubrq.sqlQuery(ctx), nil
	}
	grbuild.label = userbrewrecipe.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Status int32 `json:"status,omitempty"`
//	}
//
//	client.UserBrewRecipe.Query().
//		Select(userbrewrecipe.FieldStatus).
//		Scan(ctx, &v)
//
func (ubrq *UserBrewRecipeQuery) Select(fields ...string) *UserBrewRecipeSelect {
	ubrq.fields = append(ubrq.fields, fields...)
	selbuild := &UserBrewRecipeSelect{UserBrewRecipeQuery: ubrq}
	selbuild.label = userbrewrecipe.Label
	selbuild.flds, selbuild.scan = &ubrq.fields, selbuild.Scan
	return selbuild
}

func (ubrq *UserBrewRecipeQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ubrq.fields {
		if !userbrewrecipe.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ubrq.path != nil {
		prev, err := ubrq.path(ctx)
		if err != nil {
			return err
		}
		ubrq.sql = prev
	}
	return nil
}

func (ubrq *UserBrewRecipeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserBrewRecipe, error) {
	var (
		nodes       = []*UserBrewRecipe{}
		_spec       = ubrq.querySpec()
		loadedTypes = [2]bool{
			ubrq.withUserCoffeeBean != nil,
			ubrq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*UserBrewRecipe).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &UserBrewRecipe{config: ubrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ubrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := ubrq.withUserCoffeeBean; query != nil {
		ids := make([]int32, 0, len(nodes))
		nodeids := make(map[int32][]*UserBrewRecipe)
		for i := range nodes {
			fk := nodes[i].UserCoffeeBeanID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(usercoffeebean.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_coffee_bean_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.UserCoffeeBean = n
			}
		}
	}

	if query := ubrq.withUser; query != nil {
		ids := make([]int32, 0, len(nodes))
		nodeids := make(map[int32][]*UserBrewRecipe)
		for i := range nodes {
			fk := nodes[i].UserID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (ubrq *UserBrewRecipeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ubrq.querySpec()
	_spec.Node.Columns = ubrq.fields
	if len(ubrq.fields) > 0 {
		_spec.Unique = ubrq.unique != nil && *ubrq.unique
	}
	return sqlgraph.CountNodes(ctx, ubrq.driver, _spec)
}

func (ubrq *UserBrewRecipeQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ubrq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (ubrq *UserBrewRecipeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userbrewrecipe.Table,
			Columns: userbrewrecipe.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: userbrewrecipe.FieldID,
			},
		},
		From:   ubrq.sql,
		Unique: true,
	}
	if unique := ubrq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ubrq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userbrewrecipe.FieldID)
		for i := range fields {
			if fields[i] != userbrewrecipe.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ubrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ubrq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ubrq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ubrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ubrq *UserBrewRecipeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ubrq.driver.Dialect())
	t1 := builder.Table(userbrewrecipe.Table)
	columns := ubrq.fields
	if len(columns) == 0 {
		columns = userbrewrecipe.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ubrq.sql != nil {
		selector = ubrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ubrq.unique != nil && *ubrq.unique {
		selector.Distinct()
	}
	for _, p := range ubrq.predicates {
		p(selector)
	}
	for _, p := range ubrq.order {
		p(selector)
	}
	if offset := ubrq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ubrq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserBrewRecipeGroupBy is the group-by builder for UserBrewRecipe entities.
type UserBrewRecipeGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ubrgb *UserBrewRecipeGroupBy) Aggregate(fns ...AggregateFunc) *UserBrewRecipeGroupBy {
	ubrgb.fns = append(ubrgb.fns, fns...)
	return ubrgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ubrgb *UserBrewRecipeGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ubrgb.path(ctx)
	if err != nil {
		return err
	}
	ubrgb.sql = query
	return ubrgb.sqlScan(ctx, v)
}

func (ubrgb *UserBrewRecipeGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ubrgb.fields {
		if !userbrewrecipe.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ubrgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ubrgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ubrgb *UserBrewRecipeGroupBy) sqlQuery() *sql.Selector {
	selector := ubrgb.sql.Select()
	aggregation := make([]string, 0, len(ubrgb.fns))
	for _, fn := range ubrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ubrgb.fields)+len(ubrgb.fns))
		for _, f := range ubrgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ubrgb.fields...)...)
}

// UserBrewRecipeSelect is the builder for selecting fields of UserBrewRecipe entities.
type UserBrewRecipeSelect struct {
	*UserBrewRecipeQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ubrs *UserBrewRecipeSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ubrs.prepareQuery(ctx); err != nil {
		return err
	}
	ubrs.sql = ubrs.UserBrewRecipeQuery.sqlQuery(ctx)
	return ubrs.sqlScan(ctx, v)
}

func (ubrs *UserBrewRecipeSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ubrs.sql.Query()
	if err := ubrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
