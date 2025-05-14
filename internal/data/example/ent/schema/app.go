package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/keg/contrib/ent/mixin"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.TimeMixin{SortFieldCaseStyle: mixin.NamingStyleCamelCase},
	}
}

// Annotations of the App.
func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{
		schema.Comment("AI应用"),
		entsql.WithComments(true),
		entgql.QueryField().Directives().Description("AI应用"),
		entgql.RelayConnection(),
		entgql.Mutations(),
	}
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description").Optional(),
		field.String("token").Optional(),
		field.Enum("type").Values("chat", "agent_chat", "workflow").Comment("应用类型"),
		field.Bool("usable").Default(false).Comment("是否可用"),
		field.Time("expires_at").Optional().Comment("到期时间"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return nil
}
