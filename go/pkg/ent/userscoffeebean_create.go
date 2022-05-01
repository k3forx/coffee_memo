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
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/ent/userscoffeebean"
)

// UsersCoffeeBeanCreate is the builder for creating a UsersCoffeeBean entity.
type UsersCoffeeBeanCreate struct {
	config
	mutation *UsersCoffeeBeanMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (ucbc *UsersCoffeeBeanCreate) SetStatus(i int32) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetStatus(i)
	return ucbc
}

// SetUserID sets the "user_id" field.
func (ucbc *UsersCoffeeBeanCreate) SetUserID(i int32) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetUserID(i)
	return ucbc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucbc *UsersCoffeeBeanCreate) SetNillableUserID(i *int32) *UsersCoffeeBeanCreate {
	if i != nil {
		ucbc.SetUserID(*i)
	}
	return ucbc
}

// SetCoffeeBeanID sets the "coffee_bean_id" field.
func (ucbc *UsersCoffeeBeanCreate) SetCoffeeBeanID(i int32) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetCoffeeBeanID(i)
	return ucbc
}

// SetNillableCoffeeBeanID sets the "coffee_bean_id" field if the given value is not nil.
func (ucbc *UsersCoffeeBeanCreate) SetNillableCoffeeBeanID(i *int32) *UsersCoffeeBeanCreate {
	if i != nil {
		ucbc.SetCoffeeBeanID(*i)
	}
	return ucbc
}

// SetCreatedAt sets the "created_at" field.
func (ucbc *UsersCoffeeBeanCreate) SetCreatedAt(t time.Time) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetCreatedAt(t)
	return ucbc
}

// SetUpdatedAt sets the "updated_at" field.
func (ucbc *UsersCoffeeBeanCreate) SetUpdatedAt(t time.Time) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetUpdatedAt(t)
	return ucbc
}

// SetDeletedAt sets the "deleted_at" field.
func (ucbc *UsersCoffeeBeanCreate) SetDeletedAt(t time.Time) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetDeletedAt(t)
	return ucbc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ucbc *UsersCoffeeBeanCreate) SetNillableDeletedAt(t *time.Time) *UsersCoffeeBeanCreate {
	if t != nil {
		ucbc.SetDeletedAt(*t)
	}
	return ucbc
}

// SetID sets the "id" field.
func (ucbc *UsersCoffeeBeanCreate) SetID(i int32) *UsersCoffeeBeanCreate {
	ucbc.mutation.SetID(i)
	return ucbc
}

// SetCoffeeBean sets the "coffee_bean" edge to the CoffeeBean entity.
func (ucbc *UsersCoffeeBeanCreate) SetCoffeeBean(c *CoffeeBean) *UsersCoffeeBeanCreate {
	return ucbc.SetCoffeeBeanID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ucbc *UsersCoffeeBeanCreate) SetUser(u *User) *UsersCoffeeBeanCreate {
	return ucbc.SetUserID(u.ID)
}

// Mutation returns the UsersCoffeeBeanMutation object of the builder.
func (ucbc *UsersCoffeeBeanCreate) Mutation() *UsersCoffeeBeanMutation {
	return ucbc.mutation
}

// Save creates the UsersCoffeeBean in the database.
func (ucbc *UsersCoffeeBeanCreate) Save(ctx context.Context) (*UsersCoffeeBean, error) {
	var (
		err  error
		node *UsersCoffeeBean
	)
	if len(ucbc.hooks) == 0 {
		if err = ucbc.check(); err != nil {
			return nil, err
		}
		node, err = ucbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersCoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ucbc.check(); err != nil {
				return nil, err
			}
			ucbc.mutation = mutation
			if node, err = ucbc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ucbc.hooks) - 1; i >= 0; i-- {
			if ucbc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucbc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucbc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ucbc *UsersCoffeeBeanCreate) SaveX(ctx context.Context) *UsersCoffeeBean {
	v, err := ucbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucbc *UsersCoffeeBeanCreate) Exec(ctx context.Context) error {
	_, err := ucbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbc *UsersCoffeeBeanCreate) ExecX(ctx context.Context) {
	if err := ucbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucbc *UsersCoffeeBeanCreate) check() error {
	if _, ok := ucbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UsersCoffeeBean.status"`)}
	}
	if _, ok := ucbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UsersCoffeeBean.created_at"`)}
	}
	if _, ok := ucbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UsersCoffeeBean.updated_at"`)}
	}
	return nil
}

func (ucbc *UsersCoffeeBeanCreate) sqlSave(ctx context.Context) (*UsersCoffeeBean, error) {
	_node, _spec := ucbc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ucbc.driver, _spec); err != nil {
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

func (ucbc *UsersCoffeeBeanCreate) createSpec() (*UsersCoffeeBean, *sqlgraph.CreateSpec) {
	var (
		_node = &UsersCoffeeBean{config: ucbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userscoffeebean.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: userscoffeebean.FieldID,
			},
		}
	)
	if id, ok := ucbc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ucbc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: userscoffeebean.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ucbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userscoffeebean.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ucbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userscoffeebean.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := ucbc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userscoffeebean.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := ucbc.mutation.CoffeeBeanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userscoffeebean.CoffeeBeanTable,
			Columns: []string{userscoffeebean.CoffeeBeanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: coffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CoffeeBeanID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ucbc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userscoffeebean.UserTable,
			Columns: []string{userscoffeebean.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UsersCoffeeBeanCreateBulk is the builder for creating many UsersCoffeeBean entities in bulk.
type UsersCoffeeBeanCreateBulk struct {
	config
	builders []*UsersCoffeeBeanCreate
}

// Save creates the UsersCoffeeBean entities in the database.
func (ucbcb *UsersCoffeeBeanCreateBulk) Save(ctx context.Context) ([]*UsersCoffeeBean, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucbcb.builders))
	nodes := make([]*UsersCoffeeBean, len(ucbcb.builders))
	mutators := make([]Mutator, len(ucbcb.builders))
	for i := range ucbcb.builders {
		func(i int, root context.Context) {
			builder := ucbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersCoffeeBeanMutation)
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
					_, err = mutators[i+1].Mutate(root, ucbcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucbcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ucbcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucbcb *UsersCoffeeBeanCreateBulk) SaveX(ctx context.Context) []*UsersCoffeeBean {
	v, err := ucbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucbcb *UsersCoffeeBeanCreateBulk) Exec(ctx context.Context) error {
	_, err := ucbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbcb *UsersCoffeeBeanCreateBulk) ExecX(ctx context.Context) {
	if err := ucbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
