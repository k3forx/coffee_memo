// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.String("username"), field.String("email").Unique(), field.String("password"), field.Int("flags"), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}

	// Edges of the User.

}

func (User) Edges() []ent.Edge {
	return []ent.Edge{edge.To("user_brew_recipes", UserBrewRecipe.Type), edge.To("user_coffee_beans", UserCoffeeBean.Type)}
}
