// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
)

// CoffeeBeanUpdate is the builder for updating CoffeeBean entities.
type CoffeeBeanUpdate struct {
	config
	hooks    []Hook
	mutation *CoffeeBeanMutation
}

// Where appends a list predicates to the CoffeeBeanUpdate builder.
func (cbu *CoffeeBeanUpdate) Where(ps ...predicate.CoffeeBean) *CoffeeBeanUpdate {
	cbu.mutation.Where(ps...)
	return cbu
}

// SetName sets the "name" field.
func (cbu *CoffeeBeanUpdate) SetName(s string) *CoffeeBeanUpdate {
	cbu.mutation.SetName(s)
	return cbu
}

// SetFarmName sets the "farm_name" field.
func (cbu *CoffeeBeanUpdate) SetFarmName(s string) *CoffeeBeanUpdate {
	cbu.mutation.SetFarmName(s)
	return cbu
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (cbu *CoffeeBeanUpdate) SetNillableFarmName(s *string) *CoffeeBeanUpdate {
	if s != nil {
		cbu.SetFarmName(*s)
	}
	return cbu
}

// ClearFarmName clears the value of the "farm_name" field.
func (cbu *CoffeeBeanUpdate) ClearFarmName() *CoffeeBeanUpdate {
	cbu.mutation.ClearFarmName()
	return cbu
}

// SetCountry sets the "country" field.
func (cbu *CoffeeBeanUpdate) SetCountry(s string) *CoffeeBeanUpdate {
	cbu.mutation.SetCountry(s)
	return cbu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cbu *CoffeeBeanUpdate) SetNillableCountry(s *string) *CoffeeBeanUpdate {
	if s != nil {
		cbu.SetCountry(*s)
	}
	return cbu
}

// ClearCountry clears the value of the "country" field.
func (cbu *CoffeeBeanUpdate) ClearCountry() *CoffeeBeanUpdate {
	cbu.mutation.ClearCountry()
	return cbu
}

// SetRoastDegree sets the "roast_degree" field.
func (cbu *CoffeeBeanUpdate) SetRoastDegree(s string) *CoffeeBeanUpdate {
	cbu.mutation.SetRoastDegree(s)
	return cbu
}

// SetRoastedAt sets the "roasted_at" field.
func (cbu *CoffeeBeanUpdate) SetRoastedAt(t time.Time) *CoffeeBeanUpdate {
	cbu.mutation.SetRoastedAt(t)
	return cbu
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (cbu *CoffeeBeanUpdate) SetNillableRoastedAt(t *time.Time) *CoffeeBeanUpdate {
	if t != nil {
		cbu.SetRoastedAt(*t)
	}
	return cbu
}

// ClearRoastedAt clears the value of the "roasted_at" field.
func (cbu *CoffeeBeanUpdate) ClearRoastedAt() *CoffeeBeanUpdate {
	cbu.mutation.ClearRoastedAt()
	return cbu
}

// SetCreatedAt sets the "created_at" field.
func (cbu *CoffeeBeanUpdate) SetCreatedAt(t time.Time) *CoffeeBeanUpdate {
	cbu.mutation.SetCreatedAt(t)
	return cbu
}

// SetUpdatedAt sets the "updated_at" field.
func (cbu *CoffeeBeanUpdate) SetUpdatedAt(t time.Time) *CoffeeBeanUpdate {
	cbu.mutation.SetUpdatedAt(t)
	return cbu
}

// AddUsersCoffeeBeanIDs adds the "users_coffee_beans" edge to the UsersCoffeeBean entity by IDs.
func (cbu *CoffeeBeanUpdate) AddUsersCoffeeBeanIDs(ids ...int32) *CoffeeBeanUpdate {
	cbu.mutation.AddUsersCoffeeBeanIDs(ids...)
	return cbu
}

// AddUsersCoffeeBeans adds the "users_coffee_beans" edges to the UsersCoffeeBean entity.
func (cbu *CoffeeBeanUpdate) AddUsersCoffeeBeans(u ...*UsersCoffeeBean) *CoffeeBeanUpdate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cbu.AddUsersCoffeeBeanIDs(ids...)
}

// Mutation returns the CoffeeBeanMutation object of the builder.
func (cbu *CoffeeBeanUpdate) Mutation() *CoffeeBeanMutation {
	return cbu.mutation
}

// ClearUsersCoffeeBeans clears all "users_coffee_beans" edges to the UsersCoffeeBean entity.
func (cbu *CoffeeBeanUpdate) ClearUsersCoffeeBeans() *CoffeeBeanUpdate {
	cbu.mutation.ClearUsersCoffeeBeans()
	return cbu
}

// RemoveUsersCoffeeBeanIDs removes the "users_coffee_beans" edge to UsersCoffeeBean entities by IDs.
func (cbu *CoffeeBeanUpdate) RemoveUsersCoffeeBeanIDs(ids ...int32) *CoffeeBeanUpdate {
	cbu.mutation.RemoveUsersCoffeeBeanIDs(ids...)
	return cbu
}

// RemoveUsersCoffeeBeans removes "users_coffee_beans" edges to UsersCoffeeBean entities.
func (cbu *CoffeeBeanUpdate) RemoveUsersCoffeeBeans(u ...*UsersCoffeeBean) *CoffeeBeanUpdate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cbu.RemoveUsersCoffeeBeanIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cbu *CoffeeBeanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cbu.hooks) == 0 {
		affected, err = cbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbu.mutation = mutation
			affected, err = cbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbu.hooks) - 1; i >= 0; i-- {
			if cbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbu *CoffeeBeanUpdate) SaveX(ctx context.Context) int {
	affected, err := cbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cbu *CoffeeBeanUpdate) Exec(ctx context.Context) error {
	_, err := cbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbu *CoffeeBeanUpdate) ExecX(ctx context.Context) {
	if err := cbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cbu *CoffeeBeanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coffeebean.Table,
			Columns: coffeebean.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeebean.FieldID,
			},
		},
	}
	if ps := cbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldName,
		})
	}
	if value, ok := cbu.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldFarmName,
		})
	}
	if cbu.mutation.FarmNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coffeebean.FieldFarmName,
		})
	}
	if value, ok := cbu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldCountry,
		})
	}
	if cbu.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coffeebean.FieldCountry,
		})
	}
	if value, ok := cbu.mutation.RoastDegree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldRoastDegree,
		})
	}
	if value, ok := cbu.mutation.RoastedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldRoastedAt,
		})
	}
	if cbu.mutation.RoastedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coffeebean.FieldRoastedAt,
		})
	}
	if value, ok := cbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldCreatedAt,
		})
	}
	if value, ok := cbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldUpdatedAt,
		})
	}
	if cbu.mutation.UsersCoffeeBeansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.RemovedUsersCoffeeBeansIDs(); len(nodes) > 0 && !cbu.mutation.UsersCoffeeBeansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbu.mutation.UsersCoffeeBeansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coffeebean.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// CoffeeBeanUpdateOne is the builder for updating a single CoffeeBean entity.
type CoffeeBeanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CoffeeBeanMutation
}

// SetName sets the "name" field.
func (cbuo *CoffeeBeanUpdateOne) SetName(s string) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetName(s)
	return cbuo
}

// SetFarmName sets the "farm_name" field.
func (cbuo *CoffeeBeanUpdateOne) SetFarmName(s string) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetFarmName(s)
	return cbuo
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (cbuo *CoffeeBeanUpdateOne) SetNillableFarmName(s *string) *CoffeeBeanUpdateOne {
	if s != nil {
		cbuo.SetFarmName(*s)
	}
	return cbuo
}

// ClearFarmName clears the value of the "farm_name" field.
func (cbuo *CoffeeBeanUpdateOne) ClearFarmName() *CoffeeBeanUpdateOne {
	cbuo.mutation.ClearFarmName()
	return cbuo
}

// SetCountry sets the "country" field.
func (cbuo *CoffeeBeanUpdateOne) SetCountry(s string) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetCountry(s)
	return cbuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cbuo *CoffeeBeanUpdateOne) SetNillableCountry(s *string) *CoffeeBeanUpdateOne {
	if s != nil {
		cbuo.SetCountry(*s)
	}
	return cbuo
}

// ClearCountry clears the value of the "country" field.
func (cbuo *CoffeeBeanUpdateOne) ClearCountry() *CoffeeBeanUpdateOne {
	cbuo.mutation.ClearCountry()
	return cbuo
}

// SetRoastDegree sets the "roast_degree" field.
func (cbuo *CoffeeBeanUpdateOne) SetRoastDegree(s string) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetRoastDegree(s)
	return cbuo
}

// SetRoastedAt sets the "roasted_at" field.
func (cbuo *CoffeeBeanUpdateOne) SetRoastedAt(t time.Time) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetRoastedAt(t)
	return cbuo
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (cbuo *CoffeeBeanUpdateOne) SetNillableRoastedAt(t *time.Time) *CoffeeBeanUpdateOne {
	if t != nil {
		cbuo.SetRoastedAt(*t)
	}
	return cbuo
}

// ClearRoastedAt clears the value of the "roasted_at" field.
func (cbuo *CoffeeBeanUpdateOne) ClearRoastedAt() *CoffeeBeanUpdateOne {
	cbuo.mutation.ClearRoastedAt()
	return cbuo
}

// SetCreatedAt sets the "created_at" field.
func (cbuo *CoffeeBeanUpdateOne) SetCreatedAt(t time.Time) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetCreatedAt(t)
	return cbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cbuo *CoffeeBeanUpdateOne) SetUpdatedAt(t time.Time) *CoffeeBeanUpdateOne {
	cbuo.mutation.SetUpdatedAt(t)
	return cbuo
}

// AddUsersCoffeeBeanIDs adds the "users_coffee_beans" edge to the UsersCoffeeBean entity by IDs.
func (cbuo *CoffeeBeanUpdateOne) AddUsersCoffeeBeanIDs(ids ...int32) *CoffeeBeanUpdateOne {
	cbuo.mutation.AddUsersCoffeeBeanIDs(ids...)
	return cbuo
}

// AddUsersCoffeeBeans adds the "users_coffee_beans" edges to the UsersCoffeeBean entity.
func (cbuo *CoffeeBeanUpdateOne) AddUsersCoffeeBeans(u ...*UsersCoffeeBean) *CoffeeBeanUpdateOne {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cbuo.AddUsersCoffeeBeanIDs(ids...)
}

// Mutation returns the CoffeeBeanMutation object of the builder.
func (cbuo *CoffeeBeanUpdateOne) Mutation() *CoffeeBeanMutation {
	return cbuo.mutation
}

// ClearUsersCoffeeBeans clears all "users_coffee_beans" edges to the UsersCoffeeBean entity.
func (cbuo *CoffeeBeanUpdateOne) ClearUsersCoffeeBeans() *CoffeeBeanUpdateOne {
	cbuo.mutation.ClearUsersCoffeeBeans()
	return cbuo
}

// RemoveUsersCoffeeBeanIDs removes the "users_coffee_beans" edge to UsersCoffeeBean entities by IDs.
func (cbuo *CoffeeBeanUpdateOne) RemoveUsersCoffeeBeanIDs(ids ...int32) *CoffeeBeanUpdateOne {
	cbuo.mutation.RemoveUsersCoffeeBeanIDs(ids...)
	return cbuo
}

// RemoveUsersCoffeeBeans removes "users_coffee_beans" edges to UsersCoffeeBean entities.
func (cbuo *CoffeeBeanUpdateOne) RemoveUsersCoffeeBeans(u ...*UsersCoffeeBean) *CoffeeBeanUpdateOne {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cbuo.RemoveUsersCoffeeBeanIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cbuo *CoffeeBeanUpdateOne) Select(field string, fields ...string) *CoffeeBeanUpdateOne {
	cbuo.fields = append([]string{field}, fields...)
	return cbuo
}

// Save executes the query and returns the updated CoffeeBean entity.
func (cbuo *CoffeeBeanUpdateOne) Save(ctx context.Context) (*CoffeeBean, error) {
	var (
		err  error
		node *CoffeeBean
	)
	if len(cbuo.hooks) == 0 {
		node, err = cbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbuo.mutation = mutation
			node, err = cbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cbuo.hooks) - 1; i >= 0; i-- {
			if cbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cbuo *CoffeeBeanUpdateOne) SaveX(ctx context.Context) *CoffeeBean {
	node, err := cbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cbuo *CoffeeBeanUpdateOne) Exec(ctx context.Context) error {
	_, err := cbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbuo *CoffeeBeanUpdateOne) ExecX(ctx context.Context) {
	if err := cbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cbuo *CoffeeBeanUpdateOne) sqlSave(ctx context.Context) (_node *CoffeeBean, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   coffeebean.Table,
			Columns: coffeebean.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeebean.FieldID,
			},
		},
	}
	id, ok := cbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "CoffeeBean.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, coffeebean.FieldID)
		for _, f := range fields {
			if !coffeebean.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != coffeebean.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cbuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldName,
		})
	}
	if value, ok := cbuo.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldFarmName,
		})
	}
	if cbuo.mutation.FarmNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coffeebean.FieldFarmName,
		})
	}
	if value, ok := cbuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldCountry,
		})
	}
	if cbuo.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: coffeebean.FieldCountry,
		})
	}
	if value, ok := cbuo.mutation.RoastDegree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldRoastDegree,
		})
	}
	if value, ok := cbuo.mutation.RoastedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldRoastedAt,
		})
	}
	if cbuo.mutation.RoastedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: coffeebean.FieldRoastedAt,
		})
	}
	if value, ok := cbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldCreatedAt,
		})
	}
	if value, ok := cbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldUpdatedAt,
		})
	}
	if cbuo.mutation.UsersCoffeeBeansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.RemovedUsersCoffeeBeansIDs(); len(nodes) > 0 && !cbuo.mutation.UsersCoffeeBeansCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cbuo.mutation.UsersCoffeeBeansIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   coffeebean.UsersCoffeeBeansTable,
			Columns: []string{coffeebean.UsersCoffeeBeansColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userscoffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &CoffeeBean{config: cbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{coffeebean.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
