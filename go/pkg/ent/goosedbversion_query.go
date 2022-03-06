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
	"github.com/k3forx/coffee_memo/pkg/ent/goosedbversion"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
)

// GooseDbVersionQuery is the builder for querying GooseDbVersion entities.
type GooseDbVersionQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GooseDbVersion
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GooseDbVersionQuery builder.
func (gdvq *GooseDbVersionQuery) Where(ps ...predicate.GooseDbVersion) *GooseDbVersionQuery {
	gdvq.predicates = append(gdvq.predicates, ps...)
	return gdvq
}

// Limit adds a limit step to the query.
func (gdvq *GooseDbVersionQuery) Limit(limit int) *GooseDbVersionQuery {
	gdvq.limit = &limit
	return gdvq
}

// Offset adds an offset step to the query.
func (gdvq *GooseDbVersionQuery) Offset(offset int) *GooseDbVersionQuery {
	gdvq.offset = &offset
	return gdvq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gdvq *GooseDbVersionQuery) Unique(unique bool) *GooseDbVersionQuery {
	gdvq.unique = &unique
	return gdvq
}

// Order adds an order step to the query.
func (gdvq *GooseDbVersionQuery) Order(o ...OrderFunc) *GooseDbVersionQuery {
	gdvq.order = append(gdvq.order, o...)
	return gdvq
}

// First returns the first GooseDbVersion entity from the query.
// Returns a *NotFoundError when no GooseDbVersion was found.
func (gdvq *GooseDbVersionQuery) First(ctx context.Context) (*GooseDbVersion, error) {
	nodes, err := gdvq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goosedbversion.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) FirstX(ctx context.Context) *GooseDbVersion {
	node, err := gdvq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GooseDbVersion ID from the query.
// Returns a *NotFoundError when no GooseDbVersion ID was found.
func (gdvq *GooseDbVersionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gdvq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goosedbversion.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := gdvq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GooseDbVersion entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one GooseDbVersion entity is found.
// Returns a *NotFoundError when no GooseDbVersion entities are found.
func (gdvq *GooseDbVersionQuery) Only(ctx context.Context) (*GooseDbVersion, error) {
	nodes, err := gdvq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goosedbversion.Label}
	default:
		return nil, &NotSingularError{goosedbversion.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) OnlyX(ctx context.Context) *GooseDbVersion {
	node, err := gdvq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GooseDbVersion ID in the query.
// Returns a *NotSingularError when more than one GooseDbVersion ID is found.
// Returns a *NotFoundError when no entities are found.
func (gdvq *GooseDbVersionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = gdvq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = &NotSingularError{goosedbversion.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := gdvq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GooseDbVersions.
func (gdvq *GooseDbVersionQuery) All(ctx context.Context) ([]*GooseDbVersion, error) {
	if err := gdvq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gdvq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) AllX(ctx context.Context) []*GooseDbVersion {
	nodes, err := gdvq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GooseDbVersion IDs.
func (gdvq *GooseDbVersionQuery) IDs(ctx context.Context) ([]uint64, error) {
	var ids []uint64
	if err := gdvq.Select(goosedbversion.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := gdvq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gdvq *GooseDbVersionQuery) Count(ctx context.Context) (int, error) {
	if err := gdvq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gdvq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) CountX(ctx context.Context) int {
	count, err := gdvq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gdvq *GooseDbVersionQuery) Exist(ctx context.Context) (bool, error) {
	if err := gdvq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gdvq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gdvq *GooseDbVersionQuery) ExistX(ctx context.Context) bool {
	exist, err := gdvq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GooseDbVersionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gdvq *GooseDbVersionQuery) Clone() *GooseDbVersionQuery {
	if gdvq == nil {
		return nil
	}
	return &GooseDbVersionQuery{
		config:     gdvq.config,
		limit:      gdvq.limit,
		offset:     gdvq.offset,
		order:      append([]OrderFunc{}, gdvq.order...),
		predicates: append([]predicate.GooseDbVersion{}, gdvq.predicates...),
		// clone intermediate query.
		sql:    gdvq.sql.Clone(),
		path:   gdvq.path,
		unique: gdvq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		VersionID int `json:"version_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GooseDbVersion.Query().
//		GroupBy(goosedbversion.FieldVersionID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (gdvq *GooseDbVersionQuery) GroupBy(field string, fields ...string) *GooseDbVersionGroupBy {
	group := &GooseDbVersionGroupBy{config: gdvq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gdvq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gdvq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		VersionID int `json:"version_id,omitempty"`
//	}
//
//	client.GooseDbVersion.Query().
//		Select(goosedbversion.FieldVersionID).
//		Scan(ctx, &v)
//
func (gdvq *GooseDbVersionQuery) Select(fields ...string) *GooseDbVersionSelect {
	gdvq.fields = append(gdvq.fields, fields...)
	return &GooseDbVersionSelect{GooseDbVersionQuery: gdvq}
}

func (gdvq *GooseDbVersionQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gdvq.fields {
		if !goosedbversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gdvq.path != nil {
		prev, err := gdvq.path(ctx)
		if err != nil {
			return err
		}
		gdvq.sql = prev
	}
	return nil
}

func (gdvq *GooseDbVersionQuery) sqlAll(ctx context.Context) ([]*GooseDbVersion, error) {
	var (
		nodes = []*GooseDbVersion{}
		_spec = gdvq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GooseDbVersion{config: gdvq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gdvq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gdvq *GooseDbVersionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gdvq.querySpec()
	_spec.Node.Columns = gdvq.fields
	if len(gdvq.fields) > 0 {
		_spec.Unique = gdvq.unique != nil && *gdvq.unique
	}
	return sqlgraph.CountNodes(ctx, gdvq.driver, _spec)
}

func (gdvq *GooseDbVersionQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gdvq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (gdvq *GooseDbVersionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goosedbversion.Table,
			Columns: goosedbversion.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: goosedbversion.FieldID,
			},
		},
		From:   gdvq.sql,
		Unique: true,
	}
	if unique := gdvq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gdvq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goosedbversion.FieldID)
		for i := range fields {
			if fields[i] != goosedbversion.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gdvq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gdvq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gdvq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gdvq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gdvq *GooseDbVersionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gdvq.driver.Dialect())
	t1 := builder.Table(goosedbversion.Table)
	columns := gdvq.fields
	if len(columns) == 0 {
		columns = goosedbversion.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gdvq.sql != nil {
		selector = gdvq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gdvq.unique != nil && *gdvq.unique {
		selector.Distinct()
	}
	for _, p := range gdvq.predicates {
		p(selector)
	}
	for _, p := range gdvq.order {
		p(selector)
	}
	if offset := gdvq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gdvq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GooseDbVersionGroupBy is the group-by builder for GooseDbVersion entities.
type GooseDbVersionGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gdvgb *GooseDbVersionGroupBy) Aggregate(fns ...AggregateFunc) *GooseDbVersionGroupBy {
	gdvgb.fns = append(gdvgb.fns, fns...)
	return gdvgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gdvgb *GooseDbVersionGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gdvgb.path(ctx)
	if err != nil {
		return err
	}
	gdvgb.sql = query
	return gdvgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gdvgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gdvgb.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gdvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) StringsX(ctx context.Context) []string {
	v, err := gdvgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gdvgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) StringX(ctx context.Context) string {
	v, err := gdvgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gdvgb.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gdvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) IntsX(ctx context.Context) []int {
	v, err := gdvgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gdvgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) IntX(ctx context.Context) int {
	v, err := gdvgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gdvgb.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gdvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gdvgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gdvgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gdvgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gdvgb.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gdvgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gdvgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gdvgb *GooseDbVersionGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gdvgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gdvgb *GooseDbVersionGroupBy) BoolX(ctx context.Context) bool {
	v, err := gdvgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gdvgb *GooseDbVersionGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gdvgb.fields {
		if !goosedbversion.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gdvgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gdvgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gdvgb *GooseDbVersionGroupBy) sqlQuery() *sql.Selector {
	selector := gdvgb.sql.Select()
	aggregation := make([]string, 0, len(gdvgb.fns))
	for _, fn := range gdvgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gdvgb.fields)+len(gdvgb.fns))
		for _, f := range gdvgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gdvgb.fields...)...)
}

// GooseDbVersionSelect is the builder for selecting fields of GooseDbVersion entities.
type GooseDbVersionSelect struct {
	*GooseDbVersionQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gdvs *GooseDbVersionSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gdvs.prepareQuery(ctx); err != nil {
		return err
	}
	gdvs.sql = gdvs.GooseDbVersionQuery.sqlQuery(ctx)
	return gdvs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gdvs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gdvs.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gdvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) StringsX(ctx context.Context) []string {
	v, err := gdvs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gdvs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) StringX(ctx context.Context) string {
	v, err := gdvs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gdvs.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gdvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) IntsX(ctx context.Context) []int {
	v, err := gdvs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gdvs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) IntX(ctx context.Context) int {
	v, err := gdvs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gdvs.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gdvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gdvs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gdvs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) Float64X(ctx context.Context) float64 {
	v, err := gdvs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gdvs.fields) > 1 {
		return nil, errors.New("ent: GooseDbVersionSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gdvs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) BoolsX(ctx context.Context) []bool {
	v, err := gdvs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gdvs *GooseDbVersionSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gdvs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goosedbversion.Label}
	default:
		err = fmt.Errorf("ent: GooseDbVersionSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gdvs *GooseDbVersionSelect) BoolX(ctx context.Context) bool {
	v, err := gdvs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gdvs *GooseDbVersionSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gdvs.sql.Query()
	if err := gdvs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}