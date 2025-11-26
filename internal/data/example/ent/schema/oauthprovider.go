package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// OAuthProvider holds the schema definition for the OAuthProvider entity.
type OAuthProvider struct {
	ent.Schema
}

// Fields of the OAuthProvider.
func (OAuthProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("provider").Unique(),
		field.String("name").NotEmpty().Unique().Comment("第三方平台名称"),
		field.String("client_id").NotEmpty().Comment("OAuth2 Client ID"),
		field.String("client_secret").NotEmpty().Sensitive().Comment("OAuth2 Client Secret"),
		field.String("auth_url").NotEmpty().Comment("授权地址"),
		field.String("token_url").NotEmpty().Comment("Token 获取地址"),
		field.String("user_info_url").NotEmpty().Comment("用户信息接口地址"),
		field.String("redirect_uri").NotEmpty().Comment("回调地址"),
		field.Text("scopes").Optional().Comment("请求的权限范围，例如：user,email"),
		field.Bool("enabled").Default(true).Comment("是否启用"),
	}
}

// Edges of the OAuthProvider.
func (OAuthProvider) Edges() []ent.Edge {
	return nil
}
