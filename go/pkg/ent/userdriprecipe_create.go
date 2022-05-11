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
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/userdriprecipe"
)

// UserDripRecipeCreate is the builder for creating a UserDripRecipe entity.
type UserDripRecipeCreate struct {
	config
	mutation *UserDripRecipeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (udrc *UserDripRecipeCreate) SetUserID(i int32) *UserDripRecipeCreate {
	udrc.mutation.SetUserID(i)
	return udrc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (udrc *UserDripRecipeCreate) SetNillableUserID(i *int32) *UserDripRecipeCreate {
	if i != nil {
		udrc.SetUserID(*i)
	}
	return udrc
}

// SetCoffeeBeanID sets the "coffee_bean_id" field.
func (udrc *UserDripRecipeCreate) SetCoffeeBeanID(i int32) *UserDripRecipeCreate {
	udrc.mutation.SetCoffeeBeanID(i)
	return udrc
}

// SetNillableCoffeeBeanID sets the "coffee_bean_id" field if the given value is not nil.
func (udrc *UserDripRecipeCreate) SetNillableCoffeeBeanID(i *int32) *UserDripRecipeCreate {
	if i != nil {
		udrc.SetCoffeeBeanID(*i)
	}
	return udrc
}

// SetCoffeeBeanWeight sets the "coffee_bean_weight" field.
func (udrc *UserDripRecipeCreate) SetCoffeeBeanWeight(f float64) *UserDripRecipeCreate {
	udrc.mutation.SetCoffeeBeanWeight(f)
	return udrc
}

// SetCoffeeBeanGrind sets the "coffee_bean_grind" field.
func (udrc *UserDripRecipeCreate) SetCoffeeBeanGrind(s string) *UserDripRecipeCreate {
	udrc.mutation.SetCoffeeBeanGrind(s)
	return udrc
}

// SetLiquidWeight sets the "liquid_weight" field.
func (udrc *UserDripRecipeCreate) SetLiquidWeight(f float64) *UserDripRecipeCreate {
	udrc.mutation.SetLiquidWeight(f)
	return udrc
}

// SetTemperature sets the "temperature" field.
func (udrc *UserDripRecipeCreate) SetTemperature(f float64) *UserDripRecipeCreate {
	udrc.mutation.SetTemperature(f)
	return udrc
}

// SetStepOne sets the "step_one" field.
func (udrc *UserDripRecipeCreate) SetStepOne(s string) *UserDripRecipeCreate {
	udrc.mutation.SetStepOne(s)
	return udrc
}

// SetStepTwo sets the "step_two" field.
func (udrc *UserDripRecipeCreate) SetStepTwo(s string) *UserDripRecipeCreate {
	udrc.mutation.SetStepTwo(s)
	return udrc
}

// SetMemo sets the "memo" field.
func (udrc *UserDripRecipeCreate) SetMemo(s string) *UserDripRecipeCreate {
	udrc.mutation.SetMemo(s)
	return udrc
}

// SetCreatedAt sets the "created_at" field.
func (udrc *UserDripRecipeCreate) SetCreatedAt(t time.Time) *UserDripRecipeCreate {
	udrc.mutation.SetCreatedAt(t)
	return udrc
}

// SetUpdatedAt sets the "updated_at" field.
func (udrc *UserDripRecipeCreate) SetUpdatedAt(t time.Time) *UserDripRecipeCreate {
	udrc.mutation.SetUpdatedAt(t)
	return udrc
}

// SetDeletedAt sets the "deleted_at" field.
func (udrc *UserDripRecipeCreate) SetDeletedAt(t time.Time) *UserDripRecipeCreate {
	udrc.mutation.SetDeletedAt(t)
	return udrc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (udrc *UserDripRecipeCreate) SetNillableDeletedAt(t *time.Time) *UserDripRecipeCreate {
	if t != nil {
		udrc.SetDeletedAt(*t)
	}
	return udrc
}

// SetID sets the "id" field.
func (udrc *UserDripRecipeCreate) SetID(i int32) *UserDripRecipeCreate {
	udrc.mutation.SetID(i)
	return udrc
}

// SetUserCoffeeBeanID sets the "user_coffee_bean" edge to the UserCoffeeBean entity by ID.
func (udrc *UserDripRecipeCreate) SetUserCoffeeBeanID(id int32) *UserDripRecipeCreate {
	udrc.mutation.SetUserCoffeeBeanID(id)
	return udrc
}

// SetNillableUserCoffeeBeanID sets the "user_coffee_bean" edge to the UserCoffeeBean entity by ID if the given value is not nil.
func (udrc *UserDripRecipeCreate) SetNillableUserCoffeeBeanID(id *int32) *UserDripRecipeCreate {
	if id != nil {
		udrc = udrc.SetUserCoffeeBeanID(*id)
	}
	return udrc
}

// SetUserCoffeeBean sets the "user_coffee_bean" edge to the UserCoffeeBean entity.
func (udrc *UserDripRecipeCreate) SetUserCoffeeBean(u *UserCoffeeBean) *UserDripRecipeCreate {
	return udrc.SetUserCoffeeBeanID(u.ID)
}

// SetUser sets the "user" edge to the User entity.
func (udrc *UserDripRecipeCreate) SetUser(u *User) *UserDripRecipeCreate {
	return udrc.SetUserID(u.ID)
}

// Mutation returns the UserDripRecipeMutation object of the builder.
func (udrc *UserDripRecipeCreate) Mutation() *UserDripRecipeMutation {
	return udrc.mutation
}

// Save creates the UserDripRecipe in the database.
func (udrc *UserDripRecipeCreate) Save(ctx context.Context) (*UserDripRecipe, error) {
	var (
		err  error
		node *UserDripRecipe
	)
	if len(udrc.hooks) == 0 {
		if err = udrc.check(); err != nil {
			return nil, err
		}
		node, err = udrc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserDripRecipeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = udrc.check(); err != nil {
				return nil, err
			}
			udrc.mutation = mutation
			if node, err = udrc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(udrc.hooks) - 1; i >= 0; i-- {
			if udrc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = udrc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, udrc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (udrc *UserDripRecipeCreate) SaveX(ctx context.Context) *UserDripRecipe {
	v, err := udrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udrc *UserDripRecipeCreate) Exec(ctx context.Context) error {
	_, err := udrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udrc *UserDripRecipeCreate) ExecX(ctx context.Context) {
	if err := udrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (udrc *UserDripRecipeCreate) check() error {
	if _, ok := udrc.mutation.CoffeeBeanWeight(); !ok {
		return &ValidationError{Name: "coffee_bean_weight", err: errors.New(`ent: missing required field "UserDripRecipe.coffee_bean_weight"`)}
	}
	if _, ok := udrc.mutation.CoffeeBeanGrind(); !ok {
		return &ValidationError{Name: "coffee_bean_grind", err: errors.New(`ent: missing required field "UserDripRecipe.coffee_bean_grind"`)}
	}
	if _, ok := udrc.mutation.LiquidWeight(); !ok {
		return &ValidationError{Name: "liquid_weight", err: errors.New(`ent: missing required field "UserDripRecipe.liquid_weight"`)}
	}
	if _, ok := udrc.mutation.Temperature(); !ok {
		return &ValidationError{Name: "temperature", err: errors.New(`ent: missing required field "UserDripRecipe.temperature"`)}
	}
	if _, ok := udrc.mutation.StepOne(); !ok {
		return &ValidationError{Name: "step_one", err: errors.New(`ent: missing required field "UserDripRecipe.step_one"`)}
	}
	if _, ok := udrc.mutation.StepTwo(); !ok {
		return &ValidationError{Name: "step_two", err: errors.New(`ent: missing required field "UserDripRecipe.step_two"`)}
	}
	if _, ok := udrc.mutation.Memo(); !ok {
		return &ValidationError{Name: "memo", err: errors.New(`ent: missing required field "UserDripRecipe.memo"`)}
	}
	if _, ok := udrc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserDripRecipe.created_at"`)}
	}
	if _, ok := udrc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "UserDripRecipe.updated_at"`)}
	}
	return nil
}

func (udrc *UserDripRecipeCreate) sqlSave(ctx context.Context) (*UserDripRecipe, error) {
	_node, _spec := udrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, udrc.driver, _spec); err != nil {
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

func (udrc *UserDripRecipeCreate) createSpec() (*UserDripRecipe, *sqlgraph.CreateSpec) {
	var (
		_node = &UserDripRecipe{config: udrc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userdriprecipe.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: userdriprecipe.FieldID,
			},
		}
	)
	if id, ok := udrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := udrc.mutation.CoffeeBeanWeight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: userdriprecipe.FieldCoffeeBeanWeight,
		})
		_node.CoffeeBeanWeight = value
	}
	if value, ok := udrc.mutation.CoffeeBeanGrind(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userdriprecipe.FieldCoffeeBeanGrind,
		})
		_node.CoffeeBeanGrind = value
	}
	if value, ok := udrc.mutation.LiquidWeight(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: userdriprecipe.FieldLiquidWeight,
		})
		_node.LiquidWeight = value
	}
	if value, ok := udrc.mutation.Temperature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: userdriprecipe.FieldTemperature,
		})
		_node.Temperature = value
	}
	if value, ok := udrc.mutation.StepOne(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userdriprecipe.FieldStepOne,
		})
		_node.StepOne = value
	}
	if value, ok := udrc.mutation.StepTwo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userdriprecipe.FieldStepTwo,
		})
		_node.StepTwo = value
	}
	if value, ok := udrc.mutation.Memo(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: userdriprecipe.FieldMemo,
		})
		_node.Memo = value
	}
	if value, ok := udrc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userdriprecipe.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := udrc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userdriprecipe.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := udrc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userdriprecipe.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := udrc.mutation.UserCoffeeBeanIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userdriprecipe.UserCoffeeBeanTable,
			Columns: []string{userdriprecipe.UserCoffeeBeanColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt32,
					Column: usercoffeebean.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.CoffeeBeanID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := udrc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   userdriprecipe.UserTable,
			Columns: []string{userdriprecipe.UserColumn},
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

// UserDripRecipeCreateBulk is the builder for creating many UserDripRecipe entities in bulk.
type UserDripRecipeCreateBulk struct {
	config
	builders []*UserDripRecipeCreate
}

// Save creates the UserDripRecipe entities in the database.
func (udrcb *UserDripRecipeCreateBulk) Save(ctx context.Context) ([]*UserDripRecipe, error) {
	specs := make([]*sqlgraph.CreateSpec, len(udrcb.builders))
	nodes := make([]*UserDripRecipe, len(udrcb.builders))
	mutators := make([]Mutator, len(udrcb.builders))
	for i := range udrcb.builders {
		func(i int, root context.Context) {
			builder := udrcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserDripRecipeMutation)
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
					_, err = mutators[i+1].Mutate(root, udrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, udrcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, udrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (udrcb *UserDripRecipeCreateBulk) SaveX(ctx context.Context) []*UserDripRecipe {
	v, err := udrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (udrcb *UserDripRecipeCreateBulk) Exec(ctx context.Context) error {
	_, err := udrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (udrcb *UserDripRecipeCreateBulk) ExecX(ctx context.Context) {
	if err := udrcb.Exec(ctx); err != nil {
		panic(err)
	}
}
