// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
)

// CoffeeBeanCreate is the builder for creating a CoffeeBean entity.
type CoffeeBeanCreate struct {
	config
	mutation *CoffeeBeanMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cbc *CoffeeBeanCreate) SetName(s string) *CoffeeBeanCreate {
	cbc.mutation.SetName(s)
	return cbc
}

// SetFarmName sets the "farm_name" field.
func (cbc *CoffeeBeanCreate) SetFarmName(s string) *CoffeeBeanCreate {
	cbc.mutation.SetFarmName(s)
	return cbc
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (cbc *CoffeeBeanCreate) SetNillableFarmName(s *string) *CoffeeBeanCreate {
	if s != nil {
		cbc.SetFarmName(*s)
	}
	return cbc
}

// SetCountry sets the "country" field.
func (cbc *CoffeeBeanCreate) SetCountry(s string) *CoffeeBeanCreate {
	cbc.mutation.SetCountry(s)
	return cbc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (cbc *CoffeeBeanCreate) SetNillableCountry(s *string) *CoffeeBeanCreate {
	if s != nil {
		cbc.SetCountry(*s)
	}
	return cbc
}

// SetShopID sets the "shop_id" field.
func (cbc *CoffeeBeanCreate) SetShopID(i int32) *CoffeeBeanCreate {
	cbc.mutation.SetShopID(i)
	return cbc
}

// SetRoastDegree sets the "roast_degree" field.
func (cbc *CoffeeBeanCreate) SetRoastDegree(s string) *CoffeeBeanCreate {
	cbc.mutation.SetRoastDegree(s)
	return cbc
}

// SetRoastedAt sets the "roasted_at" field.
func (cbc *CoffeeBeanCreate) SetRoastedAt(t time.Time) *CoffeeBeanCreate {
	cbc.mutation.SetRoastedAt(t)
	return cbc
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (cbc *CoffeeBeanCreate) SetNillableRoastedAt(t *time.Time) *CoffeeBeanCreate {
	if t != nil {
		cbc.SetRoastedAt(*t)
	}
	return cbc
}

// SetCreatedAt sets the "created_at" field.
func (cbc *CoffeeBeanCreate) SetCreatedAt(t time.Time) *CoffeeBeanCreate {
	cbc.mutation.SetCreatedAt(t)
	return cbc
}

// SetUpdatedAt sets the "updated_at" field.
func (cbc *CoffeeBeanCreate) SetUpdatedAt(t time.Time) *CoffeeBeanCreate {
	cbc.mutation.SetUpdatedAt(t)
	return cbc
}

// SetID sets the "id" field.
func (cbc *CoffeeBeanCreate) SetID(i int32) *CoffeeBeanCreate {
	cbc.mutation.SetID(i)
	return cbc
}

// AddUsersCoffeeBeanIDs adds the "users_coffee_beans" edge to the UsersCoffeeBean entity by IDs.
func (cbc *CoffeeBeanCreate) AddUsersCoffeeBeanIDs(ids ...int32) *CoffeeBeanCreate {
	cbc.mutation.AddUsersCoffeeBeanIDs(ids...)
	return cbc
}

// AddUsersCoffeeBeans adds the "users_coffee_beans" edges to the UsersCoffeeBean entity.
func (cbc *CoffeeBeanCreate) AddUsersCoffeeBeans(u ...*UsersCoffeeBean) *CoffeeBeanCreate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cbc.AddUsersCoffeeBeanIDs(ids...)
}

// Mutation returns the CoffeeBeanMutation object of the builder.
func (cbc *CoffeeBeanCreate) Mutation() *CoffeeBeanMutation {
	return cbc.mutation
}

// Save creates the CoffeeBean in the database.
func (cbc *CoffeeBeanCreate) Save(ctx context.Context) (*CoffeeBean, error) {
	var (
		err  error
		node *CoffeeBean
	)
	if len(cbc.hooks) == 0 {
		if err = cbc.check(); err != nil {
			return nil, err
		}
		node, err = cbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cbc.check(); err != nil {
				return nil, err
			}
			cbc.mutation = mutation
			if node, err = cbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cbc.hooks) - 1; i >= 0; i-- {
			if cbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cbc *CoffeeBeanCreate) SaveX(ctx context.Context) *CoffeeBean {
	v, err := cbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbc *CoffeeBeanCreate) Exec(ctx context.Context) error {
	_, err := cbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbc *CoffeeBeanCreate) ExecX(ctx context.Context) {
	if err := cbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cbc *CoffeeBeanCreate) check() error {
	if _, ok := cbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "CoffeeBean.name"`)}
	}
	if _, ok := cbc.mutation.ShopID(); !ok {
		return &ValidationError{Name: "shop_id", err: errors.New(`ent: missing required field "CoffeeBean.shop_id"`)}
	}
	if _, ok := cbc.mutation.RoastDegree(); !ok {
		return &ValidationError{Name: "roast_degree", err: errors.New(`ent: missing required field "CoffeeBean.roast_degree"`)}
	}
	if _, ok := cbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "CoffeeBean.created_at"`)}
	}
	if _, ok := cbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "CoffeeBean.updated_at"`)}
	}
	return nil
}

func (cbc *CoffeeBeanCreate) sqlSave(ctx context.Context) (*CoffeeBean, error) {
	_node, _spec := cbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cbc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	return _node, nil
}

func (cbc *CoffeeBeanCreate) createSpec() (*CoffeeBean, *sqlgraph.CreateSpec) {
	var (
		_node = &CoffeeBean{config: cbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: coffeebean.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeebean.FieldID,
			},
		}
	)
	if id, ok := cbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cbc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldName,
		})
		_node.Name = value
	}
	if value, ok := cbc.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := cbc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := cbc.mutation.ShopID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: coffeebean.FieldShopID,
		})
		_node.ShopID = value
	}
	if value, ok := cbc.mutation.RoastDegree(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: coffeebean.FieldRoastDegree,
		})
		_node.RoastDegree = value
	}
	if value, ok := cbc.mutation.RoastedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldRoastedAt,
		})
		_node.RoastedAt = value
	}
	if value, ok := cbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := cbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: coffeebean.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := cbc.mutation.UsersCoffeeBeansIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CoffeeBeanCreateBulk is the builder for creating many CoffeeBean entities in bulk.
type CoffeeBeanCreateBulk struct {
	config
	builders []*CoffeeBeanCreate
}

// Save creates the CoffeeBean entities in the database.
func (cbcb *CoffeeBeanCreateBulk) Save(ctx context.Context) ([]*CoffeeBean, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cbcb.builders))
	nodes := make([]*CoffeeBean, len(cbcb.builders))
	mutators := make([]Mutator, len(cbcb.builders))
	for i := range cbcb.builders {
		func(i int, root context.Context) {
			builder := cbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CoffeeBeanMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cbcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cbcb *CoffeeBeanCreateBulk) SaveX(ctx context.Context) []*CoffeeBean {
	v, err := cbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cbcb *CoffeeBeanCreateBulk) Exec(ctx context.Context) error {
	_, err := cbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbcb *CoffeeBeanCreateBulk) ExecX(ctx context.Context) {
	if err := cbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
