// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type CoffeeBean struct {
	ent.Schema
}

func (CoffeeBean) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.String("name"), field.String("farm_name").Optional(), field.String("country").Optional(), field.String("shop_id"), field.String("roasted_degree"), field.Time("roasted_at").Optional(), field.Time("created_at"), field.Time("updated_at")}
}
func (CoffeeBean) Edges() []ent.Edge {
	return nil
}
func (CoffeeBean) Annotations() []schema.Annotation {
	return nil
}
