package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
	}
}

// Annotations of the Role.
func (Role) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("角色"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("角色"),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Annotations(entproto.Field(2)),
		field.Int("sort").Default(1000).Annotations(entproto.Field(3)),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type).Annotations(entproto.Skip()),
		edge.From("users", User.Type).Ref("roles").Annotations(entgql.Skip(), entproto.Field(4)),
	}
}
