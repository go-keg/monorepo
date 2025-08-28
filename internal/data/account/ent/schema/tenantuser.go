package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TenantUser holds the schema definition for the TenantUser entity.
type TenantUser struct {
	ent.Schema
}

// Fields of the TenantUser.
func (TenantUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("tenant_id").Immutable(),
		field.Int("user_id").Immutable(),
		field.Bool("is_owner").Default(false),
		field.Bool("is_active").Default(true),
		field.Bool("last_login_tenant").Default(false).Comment("是否为最后登录的租户"),
		field.Time("last_login_at").Optional().Comment("最后登录时间"),
	}
}

// Edges of the TenantUser.
func (TenantUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("tenant_users").Immutable().Unique().Required().Field("user_id"),
		edge.From("tenant", Tenant.Type).Ref("tenant_users").Immutable().Unique().Required().Field("tenant_id"),
		edge.To("roles", TenantRole.Type), // 每个租户内角色绑定
		edge.To("memberships", Membership.Type),
	}
}
