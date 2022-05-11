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
	"github.com/k3forx/coffee_memo/pkg/ent/predicate"
	"github.com/k3forx/coffee_memo/pkg/ent/user"
	"github.com/k3forx/coffee_memo/pkg/ent/userbrewrecipe"
	"github.com/k3forx/coffee_memo/pkg/ent/usercoffeebean"
)

// UserCoffeeBeanUpdate is the builder for updating UserCoffeeBean entities.
type UserCoffeeBeanUpdate struct {
	config
	hooks    []Hook
	mutation *UserCoffeeBeanMutation
}

// Where appends a list predicates to the UserCoffeeBeanUpdate builder.
func (ucbu *UserCoffeeBeanUpdate) Where(ps ...predicate.UserCoffeeBean) *UserCoffeeBeanUpdate {
	ucbu.mutation.Where(ps...)
	return ucbu
}

// SetStatus sets the "status" field.
func (ucbu *UserCoffeeBeanUpdate) SetStatus(i int32) *UserCoffeeBeanUpdate {
	ucbu.mutation.ResetStatus()
	ucbu.mutation.SetStatus(i)
	return ucbu
}

// AddStatus adds i to the "status" field.
func (ucbu *UserCoffeeBeanUpdate) AddStatus(i int32) *UserCoffeeBeanUpdate {
	ucbu.mutation.AddStatus(i)
	return ucbu
}

// SetUserID sets the "user_id" field.
func (ucbu *UserCoffeeBeanUpdate) SetUserID(i int32) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetUserID(i)
	return ucbu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucbu *UserCoffeeBeanUpdate) SetNillableUserID(i *int32) *UserCoffeeBeanUpdate {
	if i != nil {
		ucbu.SetUserID(*i)
	}
	return ucbu
}

// ClearUserID clears the value of the "user_id" field.
func (ucbu *UserCoffeeBeanUpdate) ClearUserID() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearUserID()
	return ucbu
}

// SetName sets the "name" field.
func (ucbu *UserCoffeeBeanUpdate) SetName(s string) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetName(s)
	return ucbu
}

// SetFarmName sets the "farm_name" field.
func (ucbu *UserCoffeeBeanUpdate) SetFarmName(s string) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetFarmName(s)
	return ucbu
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (ucbu *UserCoffeeBeanUpdate) SetNillableFarmName(s *string) *UserCoffeeBeanUpdate {
	if s != nil {
		ucbu.SetFarmName(*s)
	}
	return ucbu
}

// ClearFarmName clears the value of the "farm_name" field.
func (ucbu *UserCoffeeBeanUpdate) ClearFarmName() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearFarmName()
	return ucbu
}

// SetCountry sets the "country" field.
func (ucbu *UserCoffeeBeanUpdate) SetCountry(s string) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetCountry(s)
	return ucbu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (ucbu *UserCoffeeBeanUpdate) SetNillableCountry(s *string) *UserCoffeeBeanUpdate {
	if s != nil {
		ucbu.SetCountry(*s)
	}
	return ucbu
}

// ClearCountry clears the value of the "country" field.
func (ucbu *UserCoffeeBeanUpdate) ClearCountry() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearCountry()
	return ucbu
}

// SetRoastDegree sets the "roast_degree" field.
func (ucbu *UserCoffeeBeanUpdate) SetRoastDegree(s string) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetRoastDegree(s)
	return ucbu
}

// SetRoastedAt sets the "roasted_at" field.
func (ucbu *UserCoffeeBeanUpdate) SetRoastedAt(t time.Time) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetRoastedAt(t)
	return ucbu
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (ucbu *UserCoffeeBeanUpdate) SetNillableRoastedAt(t *time.Time) *UserCoffeeBeanUpdate {
	if t != nil {
		ucbu.SetRoastedAt(*t)
	}
	return ucbu
}

// ClearRoastedAt clears the value of the "roasted_at" field.
func (ucbu *UserCoffeeBeanUpdate) ClearRoastedAt() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearRoastedAt()
	return ucbu
}

// SetCreatedAt sets the "created_at" field.
func (ucbu *UserCoffeeBeanUpdate) SetCreatedAt(t time.Time) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetCreatedAt(t)
	return ucbu
}

// SetUpdatedAt sets the "updated_at" field.
func (ucbu *UserCoffeeBeanUpdate) SetUpdatedAt(t time.Time) *UserCoffeeBeanUpdate {
	ucbu.mutation.SetUpdatedAt(t)
	return ucbu
}

// AddUserBrewRecipeIDs adds the "user_brew_recipes" edge to the UserBrewRecipe entity by IDs.
func (ucbu *UserCoffeeBeanUpdate) AddUserBrewRecipeIDs(ids ...int32) *UserCoffeeBeanUpdate {
	ucbu.mutation.AddUserBrewRecipeIDs(ids...)
	return ucbu
}

// AddUserBrewRecipes adds the "user_brew_recipes" edges to the UserBrewRecipe entity.
func (ucbu *UserCoffeeBeanUpdate) AddUserBrewRecipes(u ...*UserBrewRecipe) *UserCoffeeBeanUpdate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucbu.AddUserBrewRecipeIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (ucbu *UserCoffeeBeanUpdate) SetUser(u *User) *UserCoffeeBeanUpdate {
	return ucbu.SetUserID(u.ID)
}

// Mutation returns the UserCoffeeBeanMutation object of the builder.
func (ucbu *UserCoffeeBeanUpdate) Mutation() *UserCoffeeBeanMutation {
	return ucbu.mutation
}

// ClearUserBrewRecipes clears all "user_brew_recipes" edges to the UserBrewRecipe entity.
func (ucbu *UserCoffeeBeanUpdate) ClearUserBrewRecipes() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearUserBrewRecipes()
	return ucbu
}

// RemoveUserBrewRecipeIDs removes the "user_brew_recipes" edge to UserBrewRecipe entities by IDs.
func (ucbu *UserCoffeeBeanUpdate) RemoveUserBrewRecipeIDs(ids ...int32) *UserCoffeeBeanUpdate {
	ucbu.mutation.RemoveUserBrewRecipeIDs(ids...)
	return ucbu
}

// RemoveUserBrewRecipes removes "user_brew_recipes" edges to UserBrewRecipe entities.
func (ucbu *UserCoffeeBeanUpdate) RemoveUserBrewRecipes(u ...*UserBrewRecipe) *UserCoffeeBeanUpdate {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucbu.RemoveUserBrewRecipeIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (ucbu *UserCoffeeBeanUpdate) ClearUser() *UserCoffeeBeanUpdate {
	ucbu.mutation.ClearUser()
	return ucbu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucbu *UserCoffeeBeanUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ucbu.hooks) == 0 {
		affected, err = ucbu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucbu.mutation = mutation
			affected, err = ucbu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ucbu.hooks) - 1; i >= 0; i-- {
			if ucbu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucbu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucbu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucbu *UserCoffeeBeanUpdate) SaveX(ctx context.Context) int {
	affected, err := ucbu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucbu *UserCoffeeBeanUpdate) Exec(ctx context.Context) error {
	_, err := ucbu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbu *UserCoffeeBeanUpdate) ExecX(ctx context.Context) {
	if err := ucbu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucbu *UserCoffeeBeanUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoffeebean.Table,
			Columns: usercoffeebean.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: usercoffeebean.FieldID,
			},
		},
	}
	if ps := ucbu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucbu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: usercoffeebean.FieldStatus,
		})
	}
	if value, ok := ucbu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: usercoffeebean.FieldStatus,
		})
	}
	if value, ok := ucbu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldName,
		})
	}
	if value, ok := ucbu.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldFarmName,
		})
	}
	if ucbu.mutation.FarmNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercoffeebean.FieldFarmName,
		})
	}
	if value, ok := ucbu.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldCountry,
		})
	}
	if ucbu.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercoffeebean.FieldCountry,
		})
	}
	if value, ok := ucbu.mutation.RoastDegree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldRoastDegree,
		})
	}
	if value, ok := ucbu.mutation.RoastedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldRoastedAt,
		})
	}
	if ucbu.mutation.RoastedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercoffeebean.FieldRoastedAt,
		})
	}
	if value, ok := ucbu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldCreatedAt,
		})
	}
	if value, ok := ucbu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldUpdatedAt,
		})
	}
	if ucbu.mutation.UserBrewRecipesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbu.mutation.RemovedUserBrewRecipesIDs(); len(nodes) > 0 && !ucbu.mutation.UserBrewRecipesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbu.mutation.UserBrewRecipesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucbu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucbu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoffeebean.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserCoffeeBeanUpdateOne is the builder for updating a single UserCoffeeBean entity.
type UserCoffeeBeanUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCoffeeBeanMutation
}

// SetStatus sets the "status" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetStatus(i int32) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ResetStatus()
	ucbuo.mutation.SetStatus(i)
	return ucbuo
}

// AddStatus adds i to the "status" field.
func (ucbuo *UserCoffeeBeanUpdateOne) AddStatus(i int32) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.AddStatus(i)
	return ucbuo
}

// SetUserID sets the "user_id" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetUserID(i int32) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetUserID(i)
	return ucbuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ucbuo *UserCoffeeBeanUpdateOne) SetNillableUserID(i *int32) *UserCoffeeBeanUpdateOne {
	if i != nil {
		ucbuo.SetUserID(*i)
	}
	return ucbuo
}

// ClearUserID clears the value of the "user_id" field.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearUserID() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearUserID()
	return ucbuo
}

// SetName sets the "name" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetName(s string) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetName(s)
	return ucbuo
}

// SetFarmName sets the "farm_name" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetFarmName(s string) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetFarmName(s)
	return ucbuo
}

// SetNillableFarmName sets the "farm_name" field if the given value is not nil.
func (ucbuo *UserCoffeeBeanUpdateOne) SetNillableFarmName(s *string) *UserCoffeeBeanUpdateOne {
	if s != nil {
		ucbuo.SetFarmName(*s)
	}
	return ucbuo
}

// ClearFarmName clears the value of the "farm_name" field.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearFarmName() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearFarmName()
	return ucbuo
}

// SetCountry sets the "country" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetCountry(s string) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetCountry(s)
	return ucbuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (ucbuo *UserCoffeeBeanUpdateOne) SetNillableCountry(s *string) *UserCoffeeBeanUpdateOne {
	if s != nil {
		ucbuo.SetCountry(*s)
	}
	return ucbuo
}

// ClearCountry clears the value of the "country" field.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearCountry() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearCountry()
	return ucbuo
}

// SetRoastDegree sets the "roast_degree" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetRoastDegree(s string) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetRoastDegree(s)
	return ucbuo
}

// SetRoastedAt sets the "roasted_at" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetRoastedAt(t time.Time) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetRoastedAt(t)
	return ucbuo
}

// SetNillableRoastedAt sets the "roasted_at" field if the given value is not nil.
func (ucbuo *UserCoffeeBeanUpdateOne) SetNillableRoastedAt(t *time.Time) *UserCoffeeBeanUpdateOne {
	if t != nil {
		ucbuo.SetRoastedAt(*t)
	}
	return ucbuo
}

// ClearRoastedAt clears the value of the "roasted_at" field.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearRoastedAt() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearRoastedAt()
	return ucbuo
}

// SetCreatedAt sets the "created_at" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetCreatedAt(t time.Time) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetCreatedAt(t)
	return ucbuo
}

// SetUpdatedAt sets the "updated_at" field.
func (ucbuo *UserCoffeeBeanUpdateOne) SetUpdatedAt(t time.Time) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.SetUpdatedAt(t)
	return ucbuo
}

// AddUserBrewRecipeIDs adds the "user_brew_recipes" edge to the UserBrewRecipe entity by IDs.
func (ucbuo *UserCoffeeBeanUpdateOne) AddUserBrewRecipeIDs(ids ...int32) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.AddUserBrewRecipeIDs(ids...)
	return ucbuo
}

// AddUserBrewRecipes adds the "user_brew_recipes" edges to the UserBrewRecipe entity.
func (ucbuo *UserCoffeeBeanUpdateOne) AddUserBrewRecipes(u ...*UserBrewRecipe) *UserCoffeeBeanUpdateOne {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucbuo.AddUserBrewRecipeIDs(ids...)
}

// SetUser sets the "user" edge to the User entity.
func (ucbuo *UserCoffeeBeanUpdateOne) SetUser(u *User) *UserCoffeeBeanUpdateOne {
	return ucbuo.SetUserID(u.ID)
}

// Mutation returns the UserCoffeeBeanMutation object of the builder.
func (ucbuo *UserCoffeeBeanUpdateOne) Mutation() *UserCoffeeBeanMutation {
	return ucbuo.mutation
}

// ClearUserBrewRecipes clears all "user_brew_recipes" edges to the UserBrewRecipe entity.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearUserBrewRecipes() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearUserBrewRecipes()
	return ucbuo
}

// RemoveUserBrewRecipeIDs removes the "user_brew_recipes" edge to UserBrewRecipe entities by IDs.
func (ucbuo *UserCoffeeBeanUpdateOne) RemoveUserBrewRecipeIDs(ids ...int32) *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.RemoveUserBrewRecipeIDs(ids...)
	return ucbuo
}

// RemoveUserBrewRecipes removes "user_brew_recipes" edges to UserBrewRecipe entities.
func (ucbuo *UserCoffeeBeanUpdateOne) RemoveUserBrewRecipes(u ...*UserBrewRecipe) *UserCoffeeBeanUpdateOne {
	ids := make([]int32, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return ucbuo.RemoveUserBrewRecipeIDs(ids...)
}

// ClearUser clears the "user" edge to the User entity.
func (ucbuo *UserCoffeeBeanUpdateOne) ClearUser() *UserCoffeeBeanUpdateOne {
	ucbuo.mutation.ClearUser()
	return ucbuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucbuo *UserCoffeeBeanUpdateOne) Select(field string, fields ...string) *UserCoffeeBeanUpdateOne {
	ucbuo.fields = append([]string{field}, fields...)
	return ucbuo
}

// Save executes the query and returns the updated UserCoffeeBean entity.
func (ucbuo *UserCoffeeBeanUpdateOne) Save(ctx context.Context) (*UserCoffeeBean, error) {
	var (
		err  error
		node *UserCoffeeBean
	)
	if len(ucbuo.hooks) == 0 {
		node, err = ucbuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserCoffeeBeanMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ucbuo.mutation = mutation
			node, err = ucbuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ucbuo.hooks) - 1; i >= 0; i-- {
			if ucbuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ucbuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ucbuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ucbuo *UserCoffeeBeanUpdateOne) SaveX(ctx context.Context) *UserCoffeeBean {
	node, err := ucbuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucbuo *UserCoffeeBeanUpdateOne) Exec(ctx context.Context) error {
	_, err := ucbuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucbuo *UserCoffeeBeanUpdateOne) ExecX(ctx context.Context) {
	if err := ucbuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ucbuo *UserCoffeeBeanUpdateOne) sqlSave(ctx context.Context) (_node *UserCoffeeBean, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   usercoffeebean.Table,
			Columns: usercoffeebean.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt32,
				Column: usercoffeebean.FieldID,
			},
		},
	}
	id, ok := ucbuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCoffeeBean.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucbuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercoffeebean.FieldID)
		for _, f := range fields {
			if !usercoffeebean.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercoffeebean.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucbuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucbuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: usercoffeebean.FieldStatus,
		})
	}
	if value, ok := ucbuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: usercoffeebean.FieldStatus,
		})
	}
	if value, ok := ucbuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldName,
		})
	}
	if value, ok := ucbuo.mutation.FarmName(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldFarmName,
		})
	}
	if ucbuo.mutation.FarmNameCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercoffeebean.FieldFarmName,
		})
	}
	if value, ok := ucbuo.mutation.Country(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldCountry,
		})
	}
	if ucbuo.mutation.CountryCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: usercoffeebean.FieldCountry,
		})
	}
	if value, ok := ucbuo.mutation.RoastDegree(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: usercoffeebean.FieldRoastDegree,
		})
	}
	if value, ok := ucbuo.mutation.RoastedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldRoastedAt,
		})
	}
	if ucbuo.mutation.RoastedAtCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: usercoffeebean.FieldRoastedAt,
		})
	}
	if value, ok := ucbuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldCreatedAt,
		})
	}
	if value, ok := ucbuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: usercoffeebean.FieldUpdatedAt,
		})
	}
	if ucbuo.mutation.UserBrewRecipesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbuo.mutation.RemovedUserBrewRecipesIDs(); len(nodes) > 0 && !ucbuo.mutation.UserBrewRecipesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbuo.mutation.UserBrewRecipesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ucbuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucbuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCoffeeBean{config: ucbuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucbuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercoffeebean.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
