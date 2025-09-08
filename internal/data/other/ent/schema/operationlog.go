package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
	schema2 "github.com/go-keg/monorepo/internal/data/account/ent/schema"
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
		field.String("tenant_id").Optional().Comment("租户ID"),
		field.String("user_id").Comment("操作人ID"),
		field.String("module").Comment("模块名，例如：user、order"),
		field.String("action").Comment("动作，例如：create、update、delete、login"),
		field.String("object_id").Optional().Comment("被操作对象ID"),
		field.String("object_name").Optional().Comment("对象名称，如用户名、订单编号"),
		field.JSON("before", map[string]any{}).Optional().Comment("修改前数据快照"),
		field.JSON("after", map[string]any{}).Optional().Comment("修改后数据快照"),
		field.String("ip").Optional().Comment("来源IP"),
		field.String("user_agent").Optional().Comment("UA信息"),
	}
}

// Edges of the OperationLog.
func (OperationLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", schema2.User.Type).Ref("operation_logs").Unique().Required().Field("user_id"),
	}
}
