// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type UsersCoffeeBean struct {
	ent.Schema
}

func (UsersCoffeeBean) Fields() []ent.Field {
	return []ent.Field{field.Int32("id"), field.Int32("user_id").Optional(), field.Int32("coffee_bean_id").Optional(), field.Time("created_at"), field.Time("updated_at"), field.Time("deleted_at").Optional()}
}
func (UsersCoffeeBean) Edges() []ent.Edge {
	return []ent.Edge{edge.From("coffee_bean", CoffeeBean.Type).Ref("users_coffee_beans").Unique().Field("coffee_bean_id"), edge.From("user", User.Type).Ref("users_coffee_beans").Unique().Field("user_id")}
}
func (UsersCoffeeBean) Annotations() []schema.Annotation {
	return nil
}
