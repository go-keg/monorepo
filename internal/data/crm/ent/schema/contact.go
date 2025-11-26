package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
	mixin2 "github.com/go-keg/monorepo/internal/data/shared/mixin"
)

// Contact holds the schema definition for the Contact entity.
type Contact struct {
	ent.Schema
}

func (Contact) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
		mixin2.OperatorMixin{},
	}
}

// Annotations of the Contact.
func (Contact) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("客户联系方式"),
		entsql.WithComments(true),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the Contact.
func (Contact) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Comment("联系人姓名"),
		field.String("position").Optional().Comment("职位"),
		field.String("phone").Optional().Comment("手机号"),
		field.String("email").Optional().Comment("邮箱"),
	}
}

// Edges of the Contact.
func (Contact) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("contacts").Unique().Required().Immutable(),
	}
}
