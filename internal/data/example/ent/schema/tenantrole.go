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

// TenantRole holds the schema definition for the TenantRole entity.
type TenantRole struct {
	ent.Schema
}

func (TenantRole) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the TenantRole.
func (TenantRole) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("角色"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("角色"),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the TenantRole.
func (TenantRole) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("tenant_id").Immutable(),
		field.String("name").NotEmpty().Annotations(entgql.OrderField("name")),
		field.String("description").Optional(),
		field.Int("sort").Default(1000).Annotations(entgql.OrderField("sort")),
	}
}

// Edges of the TenantRole.
func (TenantRole) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("permissions", Permission.Type),
		edge.From("tenant_users", TenantUser.Type).Ref("roles").Annotations(entgql.Skip()),
		edge.From("tenant", Tenant.Type).Ref("roles").Immutable().Unique().Required().Field("tenant_id").
			Annotations(entgql.Skip()),
	}
}
