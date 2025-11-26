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

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the Permission.
func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("组织架构"),
		entsql.WithComments(true),
		entgql.Mutations(),
	}
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("tenant_id").Immutable(),
		field.String("name").NotEmpty(),
		field.Int("parent_id").Optional().Nillable(),
		field.Enum("type").Values("company", "department", "team").Default("department"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("organizations").
			Unique().
			Required().
			Immutable().
			Field("tenant_id").
			Annotations(entgql.Skip()),
		edge.To("children", Organization.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)).
			From("parent").Unique().Field("parent_id").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
		edge.To("memberships", Membership.Type).
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput | entgql.SkipMutationUpdateInput)),
	}
}
