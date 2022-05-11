// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/ent/userbrewrecipe"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
)

// UserCoffeeBeanCreate is the builder for creating a UserCoffeeBean entity.
type UserCoffeeBeanCreate struct {
	config
	mutation *UserCoffeeBeanMutation
	hooks    []Hook
}

// SetStatus sets the "status" field.
func (ucbc *UserCoffeeBeanCreate) SetStatus(i int32) *UserCoffeeBeanCreate {
	ucbc.mutation.SetStatus(i)
	return ucbc
}

// SetUserID sets the "user_id" field.
func (ucbc *UserCoffeeBeanCreate) SetUserID(i int32) *UserCoffeeBeanCreate {
	ucbc.mutation.SetUserID(i)
	return ucbc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucbc *UserCoffeeBeanCreate) SetNillableUserID(i *int32) *UserCoffeeBeanCreate {
	if i != nil {
		ucbc.SetUserID(*i)
	}
	return ucbc
}

// SetName sets the "name" field.
func (ucbc *UserCoffeeBeanCreate) SetName(s string) *UserCoffeeBeanCreate {
	ucbc.mutation.SetName(s)
	return ucbc
}

// SetFarmName sets the "farm_name" field.
func (ucbc *UserCoffeeBeanCreate) SetFarmName(s string) *UserCoffeeBeanCreate {
	ucbc.mutation.SetFarmName(s)
	return ucbc
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (ucbc *UserCoffeeBeanCreate) SetNillableFarmName(s *string) *UserCoffeeBeanCreate {
	if s != nil {
		ucbc.SetFarmName(*s)
	}
	return ucbc
}

// SetCountry sets the "country" field.
func (ucbc *UserCoffeeBeanCreate) SetCountry(s string) *UserCoffeeBeanCreate {
	ucbc.mutation.SetCountry(s)
	return ucbc
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (ucbc *UserCoffeeBeanCreate) SetNillableCountry(s *string) *UserCoffeeBeanCreate {
	if s != nil {
		ucbc.SetCountry(*s)
	}
	return ucbc
}

// SetRoastDegree sets the "roast_degree" field.
func (ucbc *UserCoffeeBeanCreate) SetRoastDegree(s string) *UserCoffeeBeanCreate {
	ucbc.mutation.SetRoastDegree(s)
	return ucbc
}

// SetRoastedAt sets the "roasted_at" field.
func (ucbc *UserCoffeeBeanCreate) SetRoastedAt(t time.Time) *UserCoffeeBeanCreate {
	ucbc.mutation.SetRoastedAt(t)
	return ucbc
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (ucbc *UserCoffeeBeanCreate) SetNillableRoastedAt(t *time.Time) *UserCoffeeBeanCreate {
	if t != nil {
		ucbc.SetRoastedAt(*t)
	}
	return ucbc
}

// SetCreatedAt sets the "created_at" field.
func (ucbc *UserCoffeeBeanCreate) SetCreatedAt(t time.Time) *UserCoffeeBeanCreate {
	ucbc.mutation.SetCreatedAt(t)
	return ucbc
}

// SetUpdatedAt sets the "updated_at" field.
func (ucbc *UserCoffeeBeanCreate) SetUpdatedAt(t time.Time) *UserCoffeeBeanCreate {
	ucbc.mutation.SetUpdatedAt(t)
	return ucbc
}

// SetID sets the "id" field.
func (ucbc *UserCoffeeBeanCreate) SetID(i int32) *UserCoffeeBeanCreate {
	ucbc.mutation.SetID(i)
	return ucbc
}

// AddUserBrewRecipeIDs adds the "user_brew_recipes" edge to the UserBrewRecipe entity by IDs.
func (ucbc *UserCoffeeBeanCreate) AddUserBrewRecipeIDs(ids ...int32) *UserCoffeeBeanCreate {
	ucbc.mutation.AddUserBrewRecipeIDs(ids...)
	return ucbc
}

// AddUserBrewRecipes adds the "user_brew_recipes" edges to the UserBrewRecipe entity.
func (ucbc *UserCoffeeBeanCreate) AddUserBrewRecipes(u ...*UserBrewRecipe) *UserCoffeeBeanCreate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucbc.AddUserBrewRecipeIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (ucbc *UserCoffeeBeanCreate) SetUser(u *User) *UserCoffeeBeanCreate {
	return ucbc.SetUserID(u.ID)
}

// Mutation returns the UserCoffeeBeanMutation object of the builder.
func (ucbc *UserCoffeeBeanCreate) Mutation() *UserCoffeeBeanMutation {
	return ucbc.mutation
}

// Save creates the UserCoffeeBean in the database.
func (ucbc *UserCoffeeBeanCreate) Save(ctx context.Context) (*UserCoffeeBean, error) {
	var (
		err  error
		node *UserCoffeeBean
	)
	if len(ucbc.hooks) == 0 {
		if err = ucbc.check(); err != nil {
			return nil, err
		}
		node, err = ucbc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoffeeBeanMutation)
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
func (ucbc *UserCoffeeBeanCreate) SaveX(ctx context.Context) *UserCoffeeBean {
	v, err := ucbc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucbc *UserCoffeeBeanCreate) Exec(ctx context.Context) error {
	_, err := ucbc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbc *UserCoffeeBeanCreate) ExecX(ctx context.Context) {
	if err := ucbc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucbc *UserCoffeeBeanCreate) check() error {
	if _, ok := ucbc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "UserCoffeeBean.status"`)}
	}
	if _, ok := ucbc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "UserCoffeeBean.name"`)}
	}
	if _, ok := ucbc.mutation.RoastDegree(); !ok {
		return &ValidationError{Name: "roast_degree", err: errors.New(`ent: missing required field "UserCoffeeBean.roast_degree"`)}
	}
	if _, ok := ucbc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserCoffeeBean.created_at"`)}
	}
	if _, ok := ucbc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserCoffeeBean.updated_at"`)}
	}
	return nil
}

func (ucbc *UserCoffeeBeanCreate) sqlSave(ctx context.Context) (*UserCoffeeBean, error) {
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

func (ucbc *UserCoffeeBeanCreate) createSpec() (*UserCoffeeBean, *sqlgraph.CreateSpec) {
	var (
		_node = &UserCoffeeBean{config: ucbc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usercoffeebean.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: usercoffeebean.FieldID,
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
			Column: usercoffeebean.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := ucbc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldName,
		})
		_node.Name = value
	}
	if value, ok := ucbc.mutation.FarmName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldFarmName,
		})
		_node.FarmName = value
	}
	if value, ok := ucbc.mutation.Country(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldCountry,
		})
		_node.Country = value
	}
	if value, ok := ucbc.mutation.RoastDegree(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldRoastDegree,
		})
		_node.RoastDegree = value
	}
	if value, ok := ucbc.mutation.RoastedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldRoastedAt,
		})
		_node.RoastedAt = value
	}
	if value, ok := ucbc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := ucbc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if nodes := ucbc.mutation.UserBrewRecipesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   usercoffeebean.UserBrewRecipesTable,
			Columns: []string{usercoffeebean.UserBrewRecipesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: userbrewrecipe.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ucbc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   usercoffeebean.UserTable,
			Columns: []string{usercoffeebean.UserColumn},
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

// UserCoffeeBeanCreateBulk is the builder for creating many UserCoffeeBean entities in bulk.
type UserCoffeeBeanCreateBulk struct {
	config
	builders []*UserCoffeeBeanCreate
}

// Save creates the UserCoffeeBean entities in the database.
func (ucbcb *UserCoffeeBeanCreateBulk) Save(ctx context.Context) ([]*UserCoffeeBean, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucbcb.builders))
	nodes := make([]*UserCoffeeBean, len(ucbcb.builders))
	mutators := make([]Mutator, len(ucbcb.builders))
	for i := range ucbcb.builders {
		func(i int, root context.Context) {
			builder := ucbcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserCoffeeBeanMutation)
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
func (ucbcb *UserCoffeeBeanCreateBulk) SaveX(ctx context.Context) []*UserCoffeeBean {
	v, err := ucbcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucbcb *UserCoffeeBeanCreateBulk) Exec(ctx context.Context) error {
	_, err := ucbcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbcb *UserCoffeeBeanCreateBulk) ExecX(ctx context.Context) {
	if err := ucbcb.Exec(ctx); err != nil {
		panic(err)
	}
}
