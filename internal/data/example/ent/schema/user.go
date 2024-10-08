package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/annotations"
	"github.com/go-keg/keg/contrib/ent/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{},
		mixin.SoftDeleteMixin{},
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("用户"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("用户"),
		entgql.RelayConnection(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL"), entproto.Field(2)),
		field.String("nickname").Annotations(entproto.Field(3)),
		field.String("avatar").Optional().Annotations(entproto.Field(4)),
		field.String("password").Sensitive().Annotations(entproto.Skip()),
		field.Enum("status").Values("normal", "freeze").Comment("状态").Annotations(annotations.Enums{
			"normal": "正常",
			"freeze": "冻结",
		}, entproto.Field(5), entproto.Enum(map[string]int32{
			"normal": 1,
			"freeze": 2,
		})),
		field.Bool("is_admin").Default(false).Annotations(entproto.Field(6)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type).Annotations(entproto.Field(7)),
		edge.To("operation_logs", OperationLog.Type).Annotations(entgql.Skip(), entproto.Skip()),
	}
}
