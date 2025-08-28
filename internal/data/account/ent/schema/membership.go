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

// Membership holds the schema definition for the Membership entity.
type Membership struct {
	ent.Schema
}

func (Membership) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the Membership.
func (Membership) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("组织架构成员"),
		entsql.WithComments(true),
		entgql.Mutations(),
	}
}

// Fields of the Membership.
func (Membership) Fields() []ent.Field {
	return []ent.Field{
		field.Int("organization_id").Immutable(),
		field.Int("tenant_user_id").Immutable(),
		field.Bool("is_leader").Default(false),
	}
}

// Edges of the Membership.
func (Membership) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant_user", TenantUser.Type).
			Ref("memberships").
			Unique().
			Required().
			Immutable().
			Field("tenant_user_id"),
		edge.From("organization", Organization.Type).
			Ref("memberships").
			Unique().
			Required().
			Immutable().
			Field("organization_id"),
	}
}
