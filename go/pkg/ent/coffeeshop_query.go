// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeeshop"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
)

// CoffeeShopQuery is the builder for querying CoffeeShop entities.
type CoffeeShopQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.CoffeeShop
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CoffeeShopQuery builder.
func (csq *CoffeeShopQuery) Where(ps ...predicate.CoffeeShop) *CoffeeShopQuery {
	csq.predicates = append(csq.predicates, ps...)
	return csq
}

// Limit adds a limit step to the query.
func (csq *CoffeeShopQuery) Limit(limit int) *CoffeeShopQuery {
	csq.limit = &limit
	return csq
}

// Offset adds an offset step to the query.
func (csq *CoffeeShopQuery) Offset(offset int) *CoffeeShopQuery {
	csq.offset = &offset
	return csq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (csq *CoffeeShopQuery) Unique(unique bool) *CoffeeShopQuery {
	csq.unique = &unique
	return csq
}

// Order adds an order step to the query.
func (csq *CoffeeShopQuery) Order(o ...OrderFunc) *CoffeeShopQuery {
	csq.order = append(csq.order, o...)
	return csq
}

// First returns the first CoffeeShop entity from the query.
// Returns a *NotFoundError when no CoffeeShop was found.
func (csq *CoffeeShopQuery) First(ctx context.Context) (*CoffeeShop, error) {
	nodes, err := csq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{coffeeshop.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (csq *CoffeeShopQuery) FirstX(ctx context.Context) *CoffeeShop {
	node, err := csq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CoffeeShop ID from the query.
// Returns a *NotFoundError when no CoffeeShop ID was found.
func (csq *CoffeeShopQuery) FirstID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = csq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{coffeeshop.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (csq *CoffeeShopQuery) FirstIDX(ctx context.Context) int32 {
	id, err := csq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CoffeeShop entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CoffeeShop entity is found.
// Returns a *NotFoundError when no CoffeeShop entities are found.
func (csq *CoffeeShopQuery) Only(ctx context.Context) (*CoffeeShop, error) {
	nodes, err := csq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{coffeeshop.Label}
	default:
		return nil, &NotSingularError{coffeeshop.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (csq *CoffeeShopQuery) OnlyX(ctx context.Context) *CoffeeShop {
	node, err := csq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CoffeeShop ID in the query.
// Returns a *NotSingularError when more than one CoffeeShop ID is found.
// Returns a *NotFoundError when no entities are found.
func (csq *CoffeeShopQuery) OnlyID(ctx context.Context) (id int32, err error) {
	var ids []int32
	if ids, err = csq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = &NotSingularError{coffeeshop.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (csq *CoffeeShopQuery) OnlyIDX(ctx context.Context) int32 {
	id, err := csq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CoffeeShops.
func (csq *CoffeeShopQuery) All(ctx context.Context) ([]*CoffeeShop, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return csq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (csq *CoffeeShopQuery) AllX(ctx context.Context) []*CoffeeShop {
	nodes, err := csq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CoffeeShop IDs.
func (csq *CoffeeShopQuery) IDs(ctx context.Context) ([]int32, error) {
	var ids []int32
	if err := csq.Select(coffeeshop.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (csq *CoffeeShopQuery) IDsX(ctx context.Context) []int32 {
	ids, err := csq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (csq *CoffeeShopQuery) Count(ctx context.Context) (int, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return csq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (csq *CoffeeShopQuery) CountX(ctx context.Context) int {
	count, err := csq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (csq *CoffeeShopQuery) Exist(ctx context.Context) (bool, error) {
	if err := csq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return csq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (csq *CoffeeShopQuery) ExistX(ctx context.Context) bool {
	exist, err := csq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CoffeeShopQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (csq *CoffeeShopQuery) Clone() *CoffeeShopQuery {
	if csq == nil {
		return nil
	}
	return &CoffeeShopQuery{
		config:     csq.config,
		limit:      csq.limit,
		offset:     csq.offset,
		order:      append([]OrderFunc{}, csq.order...),
		predicates: append([]predicate.CoffeeShop{}, csq.predicates...),
		// clone intermediate query.
		sql:    csq.sql.Clone(),
		path:   csq.path,
		unique: csq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CoffeeShop.Query().
//		GroupBy(coffeeshop.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (csq *CoffeeShopQuery) GroupBy(field string, fields ...string) *CoffeeShopGroupBy {
	group := &CoffeeShopGroupBy{config: csq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := csq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return csq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.CoffeeShop.Query().
//		Select(coffeeshop.FieldName).
//		Scan(ctx, &v)
//
func (csq *CoffeeShopQuery) Select(fields ...string) *CoffeeShopSelect {
	csq.fields = append(csq.fields, fields...)
	return &CoffeeShopSelect{CoffeeShopQuery: csq}
}

func (csq *CoffeeShopQuery) prepareQuery(ctx context.Context) error {
	for _, f := range csq.fields {
		if !coffeeshop.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if csq.path != nil {
		prev, err := csq.path(ctx)
		if err != nil {
			return err
		}
		csq.sql = prev
	}
	return nil
}

func (csq *CoffeeShopQuery) sqlAll(ctx context.Context) ([]*CoffeeShop, error) {
	var (
		nodes = []*CoffeeShop{}
		_spec = csq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &CoffeeShop{config: csq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, csq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (csq *CoffeeShopQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := csq.querySpec()
	_spec.Node.Columns = csq.fields
	if len(csq.fields) > 0 {
		_spec.Unique = csq.unique != nil && *csq.unique
	}
	return sqlgraph.CountNodes(ctx, csq.driver, _spec)
}

func (csq *CoffeeShopQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := csq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (csq *CoffeeShopQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coffeeshop.Table,
			Columns: coffeeshop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeeshop.FieldID,
			},
		},
		From:   csq.sql,
		Unique: true,
	}
	if unique := csq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := csq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coffeeshop.FieldID)
		for i := range fields {
			if fields[i] != coffeeshop.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := csq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := csq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := csq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := csq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (csq *CoffeeShopQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(csq.driver.Dialect())
	t1 := builder.Table(coffeeshop.Table)
	columns := csq.fields
	if len(columns) == 0 {
		columns = coffeeshop.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if csq.sql != nil {
		selector = csq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if csq.unique != nil && *csq.unique {
		selector.Distinct()
	}
	for _, p := range csq.predicates {
		p(selector)
	}
	for _, p := range csq.order {
		p(selector)
	}
	if offset := csq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := csq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CoffeeShopGroupBy is the group-by builder for CoffeeShop entities.
type CoffeeShopGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (csgb *CoffeeShopGroupBy) Aggregate(fns ...AggregateFunc) *CoffeeShopGroupBy {
	csgb.fns = append(csgb.fns, fns...)
	return csgb
}

// Scan applies the group-by query and scans the result into the given value.
func (csgb *CoffeeShopGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := csgb.path(ctx)
	if err != nil {
		return err
	}
	csgb.sql = query
	return csgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := csgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) StringsX(ctx context.Context) []string {
	v, err := csgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = csgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) StringX(ctx context.Context) string {
	v, err := csgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) IntsX(ctx context.Context) []int {
	v, err := csgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = csgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) IntX(ctx context.Context) int {
	v, err := csgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := csgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = csgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) Float64X(ctx context.Context) float64 {
	v, err := csgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(csgb.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := csgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := csgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (csgb *CoffeeShopGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = csgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (csgb *CoffeeShopGroupBy) BoolX(ctx context.Context) bool {
	v, err := csgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (csgb *CoffeeShopGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range csgb.fields {
		if !coffeeshop.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := csgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := csgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (csgb *CoffeeShopGroupBy) sqlQuery() *sql.Selector {
	selector := csgb.sql.Select()
	aggregation := make([]string, 0, len(csgb.fns))
	for _, fn := range csgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(csgb.fields)+len(csgb.fns))
		for _, f := range csgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(csgb.fields...)...)
}

// CoffeeShopSelect is the builder for selecting fields of CoffeeShop entities.
type CoffeeShopSelect struct {
	*CoffeeShopQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (css *CoffeeShopSelect) Scan(ctx context.Context, v interface{}) error {
	if err := css.prepareQuery(ctx); err != nil {
		return err
	}
	css.sql = css.CoffeeShopQuery.sqlQuery(ctx)
	return css.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (css *CoffeeShopSelect) ScanX(ctx context.Context, v interface{}) {
	if err := css.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Strings(ctx context.Context) ([]string, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (css *CoffeeShopSelect) StringsX(ctx context.Context) []string {
	v, err := css.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = css.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (css *CoffeeShopSelect) StringX(ctx context.Context) string {
	v, err := css.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Ints(ctx context.Context) ([]int, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (css *CoffeeShopSelect) IntsX(ctx context.Context) []int {
	v, err := css.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = css.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (css *CoffeeShopSelect) IntX(ctx context.Context) int {
	v, err := css.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (css *CoffeeShopSelect) Float64sX(ctx context.Context) []float64 {
	v, err := css.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = css.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (css *CoffeeShopSelect) Float64X(ctx context.Context) float64 {
	v, err := css.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(css.fields) > 1 {
		return nil, errors.New("ent: CoffeeShopSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := css.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (css *CoffeeShopSelect) BoolsX(ctx context.Context) []bool {
	v, err := css.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (css *CoffeeShopSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = css.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{coffeeshop.Label}
	default:
		err = fmt.Errorf("ent: CoffeeShopSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (css *CoffeeShopSelect) BoolX(ctx context.Context) bool {
	v, err := css.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (css *CoffeeShopSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := css.sql.Query()
	if err := css.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
