package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

func (Permission) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the Permission.
func (Permission) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("权限"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("权限"),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_id").Optional().Nillable().Annotations(entgql.Skip(entgql.SkipMutationUpdateInput)),
		field.String("name"),
		field.String("key").Optional().Unique(),
		field.Enum("type").Values("menu", "page", "element"),
		field.String("path").Optional(),
		field.String("desc").Optional(),
		field.Int("sort").Default(1000),
		field.JSON("attrs", map[string]any{}).Optional(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("roles", Role.Type).Ref("permissions").Annotations(entgql.Skip()),
		edge.To("children", Permission.Type).From("parent").Field("parent_id").Unique(),
	}
}
