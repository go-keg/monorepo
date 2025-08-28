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

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

func (Payment) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the Payment.
func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("收款记录"),
		entsql.WithComments(true),
		entgql.Mutations(),
	}
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Float("amount").Default(0).Comment("收款金额"),
		field.Time("received_at").Optional().Comment("收款日期"),
		field.Int("created_by").Immutable().Comment("创建人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Int("updated_by").Optional().Comment("修改人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contract", Contract.Type).Ref("payments").Unique().Required().Immutable(),
	}
}
