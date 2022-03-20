// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/k3forx/coffee_memo/pkg/ent/coffeeshop"
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
)

// CoffeeShopDelete is the builder for deleting a CoffeeShop entity.
type CoffeeShopDelete struct {
	config
	hooks    []Hook
	mutation *CoffeeShopMutation
}

// Where appends a list predicates to the CoffeeShopDelete builder.
func (csd *CoffeeShopDelete) Where(ps ...predicate.CoffeeShop) *CoffeeShopDelete {
	csd.mutation.Where(ps...)
	return csd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (csd *CoffeeShopDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(csd.hooks) == 0 {
		affected, err = csd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CoffeeShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			csd.mutation = mutation
			affected, err = csd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(csd.hooks) - 1; i >= 0; i-- {
			if csd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = csd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, csd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (csd *CoffeeShopDelete) ExecX(ctx context.Context) int {
	n, err := csd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (csd *CoffeeShopDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: coffeeshop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: coffeeshop.FieldID,
			},
		},
	}
	if ps := csd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, csd.driver, _spec)
}

// CoffeeShopDeleteOne is the builder for deleting a single CoffeeShop entity.
type CoffeeShopDeleteOne struct {
	csd *CoffeeShopDelete
}

// Exec executes the deletion query.
func (csdo *CoffeeShopDeleteOne) Exec(ctx context.Context) error {
	n, err := csdo.csd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{coffeeshop.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (csdo *CoffeeShopDeleteOne) ExecX(ctx context.Context) {
	csdo.csd.ExecX(ctx)
}
