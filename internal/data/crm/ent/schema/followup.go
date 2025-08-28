package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
)

// FollowUp holds the schema definition for the FollowUp entity.
type FollowUp struct {
	ent.Schema
}

func (FollowUp) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the FollowUp.
func (FollowUp) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("跟进记录"),
		entsql.WithComments(true),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the FollowUp.
func (FollowUp) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values("call", "meeting", "email", "visit", "other").
			Default("other").
			Comment("跟进类型"),
		field.String("content").NotEmpty().Comment("跟进内容"),
		field.Time("followed_at").Default(time.Now).Comment("跟进时间"),
		field.Int("created_by").Immutable().Comment("创建人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput)),
		field.Int("updated_by").Optional().Comment("修改人").
			Annotations(entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the FollowUp.
func (FollowUp) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("customer", Customer.Type).Ref("follow_ups").Unique().Required().Immutable(),
	}
}
