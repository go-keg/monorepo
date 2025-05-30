// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"

	"github.com/go-keg/monorepo/internal/data/example/ent/oauthaccount"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
)

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitempty"`     // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

// Node implements Noder interface
func (a *App) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     a.ID,
		Type:   "App",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(a.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Token); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "token",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Type); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "app.Type",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.Usable); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "bool",
		Name:  "usable",
		Value: string(buf),
	}
	if buf, err = json.Marshal(a.ExpiresAt); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "time.Time",
		Name:  "expires_at",
		Value: string(buf),
	}
	return node, nil
}

// Node implements Noder interface
func (oa *OAuthAccount) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     oa.ID,
		Type:   "OAuthAccount",
		Fields: make([]*Field, 7),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(oa.UserID); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.Provider); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "oauthaccount.Provider",
		Name:  "provider",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.ProviderUserID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "provider_user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.AccessToken); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "access_token",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.RefreshToken); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "refresh_token",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.TokenExpiry); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "time.Time",
		Name:  "token_expiry",
		Value: string(buf),
	}
	if buf, err = json.Marshal(oa.Profile); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "map[string]interface {}",
		Name:  "profile",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = oa.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (ol *OperationLog) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     ol.ID,
		Type:   "OperationLog",
		Fields: make([]*Field, 6),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(ol.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ol.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ol.UserID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "user_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ol.Type); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ol.Content); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "content",
		Value: string(buf),
	}
	if buf, err = json.Marshal(ol.Metadata); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "map[string]interface {}",
		Name:  "metadata",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "User",
		Name: "user",
	}
	err = ol.QueryUser().
		Select(user.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (pe *Permission) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pe.ID,
		Type:   "Permission",
		Fields: make([]*Field, 10),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(pe.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.ParentID); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "int",
		Name:  "parent_id",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Name); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Key); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "key",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Type); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "permission.Type",
		Name:  "type",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Path); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "string",
		Name:  "path",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Description); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Sort); err != nil {
		return nil, err
	}
	node.Fields[8] = &Field{
		Type:  "int",
		Name:  "sort",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.Attrs); err != nil {
		return nil, err
	}
	node.Fields[9] = &Field{
		Type:  "map[string]interface {}",
		Name:  "attrs",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Permission",
		Name: "parent",
	}
	err = pe.QueryParent().
		Select(permission.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "Permission",
		Name: "children",
	}
	err = pe.QueryChildren().
		Select(permission.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (r *Role) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     r.ID,
		Type:   "Role",
		Fields: make([]*Field, 5),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(r.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Name); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "name",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Description); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "description",
		Value: string(buf),
	}
	if buf, err = json.Marshal(r.Sort); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "int",
		Name:  "sort",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Permission",
		Name: "permissions",
	}
	err = r.QueryPermissions().
		Select(permission.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node implements Noder interface
func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 8),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(u.CreatedAt); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "time.Time",
		Name:  "created_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.UpdatedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "updated_at",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Email); err != nil {
		return nil, err
	}
	node.Fields[2] = &Field{
		Type:  "string",
		Name:  "email",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Nickname); err != nil {
		return nil, err
	}
	node.Fields[3] = &Field{
		Type:  "string",
		Name:  "nickname",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Avatar); err != nil {
		return nil, err
	}
	node.Fields[4] = &Field{
		Type:  "string",
		Name:  "avatar",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Password); err != nil {
		return nil, err
	}
	node.Fields[5] = &Field{
		Type:  "string",
		Name:  "password",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.Status); err != nil {
		return nil, err
	}
	node.Fields[6] = &Field{
		Type:  "user.Status",
		Name:  "status",
		Value: string(buf),
	}
	if buf, err = json.Marshal(u.IsAdmin); err != nil {
		return nil, err
	}
	node.Fields[7] = &Field{
		Type:  "bool",
		Name:  "is_admin",
		Value: string(buf),
	}
	node.Edges[0] = &Edge{
		Type: "Role",
		Name: "roles",
	}
	err = u.QueryRoles().
		Select(role.FieldID).
		Scan(ctx, &node.Edges[0].IDs)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		Type: "OAuthAccount",
		Name: "oauth_accounts",
	}
	err = u.QueryOauthAccounts().
		Select(oauthaccount.FieldID).
		Scan(ctx, &node.Edges[1].IDs)
	if err != nil {
		return nil, err
	}
	return node, nil
}

// Node returns the node with given global ID.
//
// This API helpful in case you want to build
// an administrator tool to browser all types in system.
func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	n, err := c.Noder(ctx, id)
	if err != nil {
		return nil, err
	}
	return n.Node(ctx)
}
