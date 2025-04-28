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

// OperationLog holds the schema definition for the OperationLog entity.
type OperationLog struct {
	ent.Schema
}

func (OperationLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the OperationLog.
func (OperationLog) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("操作日志"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("操作日志"),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the OperationLog.
func (OperationLog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("type").Comment("操作类型"),
		field.JSON("context", map[string]any{}).Annotations(entgql.Type("Map")),
	}
}

// Edges of the OperationLog.
func (OperationLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("operation_logs").Unique().Required().Field("user_id"),
	}
}
