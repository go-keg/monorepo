package mixin

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type OperatorMixin struct {
	mixin.Schema
}

func (r OperatorMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("created_by").Immutable().Comment("创建人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Int("updated_by").Optional().Comment("修改人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}
