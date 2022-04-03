// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeebean"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
)

// CoffeeBeanDelete is the builder for deleting a CoffeeBean entity.
type CoffeeBeanDelete struct {
	config
	hooks    []Hook
	mutation *CoffeeBeanMutation
}

// Where appends a list predicates to the CoffeeBeanDelete builder.
func (cbd *CoffeeBeanDelete) Where(ps ...predicate.CoffeeBean) *CoffeeBeanDelete {
	cbd.mutation.Where(ps...)
	return cbd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cbd *CoffeeBeanDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cbd.hooks) == 0 {
		affected, err = cbd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cbd.mutation = mutation
			affected, err = cbd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cbd.hooks) - 1; i >= 0; i-- {
			if cbd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cbd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cbd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cbd *CoffeeBeanDelete) ExecX(ctx context.Context) int {
	n, err := cbd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cbd *CoffeeBeanDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coffeebean.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeebean.FieldID,
			},
		},
	}
	if ps := cbd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cbd.driver, _spec)
}

// CoffeeBeanDeleteOne is the builder for deleting a single CoffeeBean entity.
type CoffeeBeanDeleteOne struct {
	cbd *CoffeeBeanDelete
}

// Exec executes the deletion query.
func (cbdo *CoffeeBeanDeleteOne) Exec(ctx context.Context) error {
	n, err := cbdo.cbd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coffeebean.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cbdo *CoffeeBeanDeleteOne) ExecX(ctx context.Context) {
	cbdo.cbd.ExecX(ctx)
}
