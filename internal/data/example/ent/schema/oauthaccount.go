package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OAuthAccount holds the schema definition for the OAuthAccount entity.
type OAuthAccount struct {
	ent.Schema
}

func (OAuthAccount) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("provider", "provider_user_id").Unique(),
	}
}

// Fields of the OAuthAccount.
func (OAuthAccount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.String("provider"),
		field.String("provider_user_id"),
		field.String("access_token").Optional(),
		field.String("refresh_token").Optional(),
		field.Time("token_expiry").Optional(),
		field.JSON("profile", map[string]any{}).Optional(),
	}

}

// Edges of the OAuthAccount.
func (OAuthAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("oauth_accounts").Unique().Required().Field("user_id"),
	}
}
