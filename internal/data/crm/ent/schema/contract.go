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

// Contract holds the schema definition for the Contract entity.
type Contract struct {
	ent.Schema
}

func (Contract) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the Contract.
func (Contract) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("合同"),
		entsql.WithComments(true),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the Contract.
func (Contract) Fields() []ent.Field {
	return []ent.Field{
		field.String("contract_no").Unique().NotEmpty().Comment("合同编号"),
		field.Float("amount").Default(0).Comment("合同金额"),
		field.Time("signed_at").Optional().Comment("签订日期"),
		field.Time("end_at").Optional().Comment("到期日期"),
		field.Int("created_by").Immutable().Comment("创建人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Int("updated_by").Optional().Comment("修改人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the Contract.
func (Contract) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("contracts").Unique().Required().Immutable(),
		edge.To("payments", Payment.Type),
	}
}
