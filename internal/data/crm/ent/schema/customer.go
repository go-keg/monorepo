package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
	mixin2 "github.com/go-keg/monorepo/internal/data/mixin"
)

// Customer holds the schema definition for the Customer entity.
type Customer struct {
	ent.Schema
}

func (Customer) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
		mixin2.OperatorMixin{},
	}
}

// Annotations of the Customer.
func (Customer) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("客户"),
		entsql.WithComments(true),
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(),
	}
}

// Fields of the Customer.
func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("客户名称"),
		field.String("industry").Optional().Comment("所属行业"),
		field.String("source").Optional().Comment("客户来源"),
		field.String("level").Optional().Comment("客户等级"),
		field.JSON("metadata", map[string]interface{}{}).
			Optional().
			Comment("自定义扩展字段"),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contacts", Contact.Type).
			Annotations(entgql.RelayConnection(), entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput)),
		edge.To("contracts", Contract.Type).
			Annotations(entgql.RelayConnection(), entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput)),
		edge.To("follow_ups", FollowUp.Type).
			Annotations(entgql.RelayConnection(), entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput)),
	}
}
