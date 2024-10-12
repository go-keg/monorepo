// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
)

// CreateOperationLogInput represents a mutation input for creating operationlogs.
type CreateOperationLogInput struct {
	Type    string
	Context map[string]interface{}
	UserID  int
}

// Mutate applies the CreateOperationLogInput on the OperationLogMutation builder.
func (i *CreateOperationLogInput) Mutate(m *OperationLogMutation) {
	m.SetType(i.Type)
	if v := i.Context; v != nil {
		m.SetContext(v)
	}
	m.SetUserID(i.UserID)
}

// SetInput applies the change-set in the CreateOperationLogInput on the OperationLogCreate builder.
func (c *OperationLogCreate) SetInput(i CreateOperationLogInput) *OperationLogCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOperationLogInput represents a mutation input for updating operationlogs.
type UpdateOperationLogInput struct {
	Type    *string
	Context map[string]interface{}
	UserID  *int
}

// Mutate applies the UpdateOperationLogInput on the OperationLogMutation builder.
func (i *UpdateOperationLogInput) Mutate(m *OperationLogMutation) {
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if v := i.Context; v != nil {
		m.SetContext(v)
	}
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
}

// SetInput applies the change-set in the UpdateOperationLogInput on the OperationLogUpdate builder.
func (c *OperationLogUpdate) SetInput(i UpdateOperationLogInput) *OperationLogUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOperationLogInput on the OperationLogUpdateOne builder.
func (c *OperationLogUpdateOne) SetInput(i UpdateOperationLogInput) *OperationLogUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreatePermissionInput represents a mutation input for creating permissions.
type CreatePermissionInput struct {
	Name     string
	Key      *string
	Type     permission.Type
	Path     *string
	Desc     *string
	Sort     *int
	Attrs    map[string]interface{}
	ParentID *int
	ChildIDs []int
}

// Mutate applies the CreatePermissionInput on the PermissionMutation builder.
func (i *CreatePermissionInput) Mutate(m *PermissionMutation) {
	m.SetName(i.Name)
	if v := i.Key; v != nil {
		m.SetKey(*v)
	}
	m.SetType(i.Type)
	if v := i.Path; v != nil {
		m.SetPath(*v)
	}
	if v := i.Desc; v != nil {
		m.SetDesc(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.Attrs; v != nil {
		m.SetAttrs(v)
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if v := i.ChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
}

// SetInput applies the change-set in the CreatePermissionInput on the PermissionCreate builder.
func (c *PermissionCreate) SetInput(i CreatePermissionInput) *PermissionCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdatePermissionInput represents a mutation input for updating permissions.
type UpdatePermissionInput struct {
	Name           *string
	ClearKey       bool
	Key            *string
	Type           *permission.Type
	ClearPath      bool
	Path           *string
	ClearDesc      bool
	Desc           *string
	Sort           *int
	ClearAttrs     bool
	Attrs          map[string]interface{}
	ClearParent    bool
	ParentID       *int
	ClearChildren  bool
	AddChildIDs    []int
	RemoveChildIDs []int
}

// Mutate applies the UpdatePermissionInput on the PermissionMutation builder.
func (i *UpdatePermissionInput) Mutate(m *PermissionMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearKey {
		m.ClearKey()
	}
	if v := i.Key; v != nil {
		m.SetKey(*v)
	}
	if v := i.Type; v != nil {
		m.SetType(*v)
	}
	if i.ClearPath {
		m.ClearPath()
	}
	if v := i.Path; v != nil {
		m.SetPath(*v)
	}
	if i.ClearDesc {
		m.ClearDesc()
	}
	if v := i.Desc; v != nil {
		m.SetDesc(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if i.ClearAttrs {
		m.ClearAttrs()
	}
	if v := i.Attrs; v != nil {
		m.SetAttrs(v)
	}
	if i.ClearParent {
		m.ClearParent()
	}
	if v := i.ParentID; v != nil {
		m.SetParentID(*v)
	}
	if i.ClearChildren {
		m.ClearChildren()
	}
	if v := i.AddChildIDs; len(v) > 0 {
		m.AddChildIDs(v...)
	}
	if v := i.RemoveChildIDs; len(v) > 0 {
		m.RemoveChildIDs(v...)
	}
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdate builder.
func (c *PermissionUpdate) SetInput(i UpdatePermissionInput) *PermissionUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdatePermissionInput on the PermissionUpdateOne builder.
func (c *PermissionUpdateOne) SetInput(i UpdatePermissionInput) *PermissionUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateRoleInput represents a mutation input for creating roles.
type CreateRoleInput struct {
	Name          string
	Sort          *int
	PermissionIDs []int
}

// Mutate applies the CreateRoleInput on the RoleMutation builder.
func (i *CreateRoleInput) Mutate(m *RoleMutation) {
	m.SetName(i.Name)
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if v := i.PermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
}

// SetInput applies the change-set in the CreateRoleInput on the RoleCreate builder.
func (c *RoleCreate) SetInput(i CreateRoleInput) *RoleCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateRoleInput represents a mutation input for updating roles.
type UpdateRoleInput struct {
	Name                *string
	Sort                *int
	ClearPermissions    bool
	AddPermissionIDs    []int
	RemovePermissionIDs []int
}

// Mutate applies the UpdateRoleInput on the RoleMutation builder.
func (i *UpdateRoleInput) Mutate(m *RoleMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if v := i.Sort; v != nil {
		m.SetSort(*v)
	}
	if i.ClearPermissions {
		m.ClearPermissions()
	}
	if v := i.AddPermissionIDs; len(v) > 0 {
		m.AddPermissionIDs(v...)
	}
	if v := i.RemovePermissionIDs; len(v) > 0 {
		m.RemovePermissionIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateRoleInput on the RoleUpdate builder.
func (c *RoleUpdate) SetInput(i UpdateRoleInput) *RoleUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateRoleInput on the RoleUpdateOne builder.
func (c *RoleUpdateOne) SetInput(i UpdateRoleInput) *RoleUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateUserInput represents a mutation input for creating users.
type CreateUserInput struct {
	Email    string
	Nickname string
	Avatar   *string
	Password string
	Status   user.Status
	IsAdmin  *bool
	RoleIDs  []int
}

// Mutate applies the CreateUserInput on the UserMutation builder.
func (i *CreateUserInput) Mutate(m *UserMutation) {
	m.SetEmail(i.Email)
	m.SetNickname(i.Nickname)
	if v := i.Avatar; v != nil {
		m.SetAvatar(*v)
	}
	m.SetPassword(i.Password)
	m.SetStatus(i.Status)
	if v := i.IsAdmin; v != nil {
		m.SetIsAdmin(*v)
	}
	if v := i.RoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
}

// SetInput applies the change-set in the CreateUserInput on the UserCreate builder.
func (c *UserCreate) SetInput(i CreateUserInput) *UserCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateUserInput represents a mutation input for updating users.
type UpdateUserInput struct {
	Email         *string
	Nickname      *string
	ClearAvatar   bool
	Avatar        *string
	Password      *string
	Status        *user.Status
	IsAdmin       *bool
	ClearRoles    bool
	AddRoleIDs    []int
	RemoveRoleIDs []int
}

// Mutate applies the UpdateUserInput on the UserMutation builder.
func (i *UpdateUserInput) Mutate(m *UserMutation) {
	if v := i.Email; v != nil {
		m.SetEmail(*v)
	}
	if v := i.Nickname; v != nil {
		m.SetNickname(*v)
	}
	if i.ClearAvatar {
		m.ClearAvatar()
	}
	if v := i.Avatar; v != nil {
		m.SetAvatar(*v)
	}
	if v := i.Password; v != nil {
		m.SetPassword(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.IsAdmin; v != nil {
		m.SetIsAdmin(*v)
	}
	if i.ClearRoles {
		m.ClearRoles()
	}
	if v := i.AddRoleIDs; len(v) > 0 {
		m.AddRoleIDs(v...)
	}
	if v := i.RemoveRoleIDs; len(v) > 0 {
		m.RemoveRoleIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdate builder.
func (c *UserUpdate) SetInput(i UpdateUserInput) *UserUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateUserInput on the UserUpdateOne builder.
func (c *UserUpdateOne) SetInput(i UpdateUserInput) *UserUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
