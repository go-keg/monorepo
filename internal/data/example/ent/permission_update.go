// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/predicate"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
)

// PermissionUpdate is the builder for updating Permission entities.
type PermissionUpdate struct {
	config
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (pu *PermissionUpdate) Where(ps ...predicate.Permission) *PermissionUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdatedAt sets the "updated_at" field.
func (pu *PermissionUpdate) SetUpdatedAt(t time.Time) *PermissionUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (pu *PermissionUpdate) ClearUpdatedAt() *PermissionUpdate {
	pu.mutation.ClearUpdatedAt()
	return pu
}

// SetParentID sets the "parent_id" field.
func (pu *PermissionUpdate) SetParentID(i int) *PermissionUpdate {
	pu.mutation.SetParentID(i)
	return pu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableParentID(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetParentID(*i)
	}
	return pu
}

// ClearParentID clears the value of the "parent_id" field.
func (pu *PermissionUpdate) ClearParentID() *PermissionUpdate {
	pu.mutation.ClearParentID()
	return pu
}

// SetName sets the "name" field.
func (pu *PermissionUpdate) SetName(s string) *PermissionUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableName(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetKey sets the "key" field.
func (pu *PermissionUpdate) SetKey(s string) *PermissionUpdate {
	pu.mutation.SetKey(s)
	return pu
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableKey(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetKey(*s)
	}
	return pu
}

// ClearKey clears the value of the "key" field.
func (pu *PermissionUpdate) ClearKey() *PermissionUpdate {
	pu.mutation.ClearKey()
	return pu
}

// SetType sets the "type" field.
func (pu *PermissionUpdate) SetType(pe permission.Type) *PermissionUpdate {
	pu.mutation.SetType(pe)
	return pu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableType(pe *permission.Type) *PermissionUpdate {
	if pe != nil {
		pu.SetType(*pe)
	}
	return pu
}

// SetPath sets the "path" field.
func (pu *PermissionUpdate) SetPath(s string) *PermissionUpdate {
	pu.mutation.SetPath(s)
	return pu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillablePath(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetPath(*s)
	}
	return pu
}

// ClearPath clears the value of the "path" field.
func (pu *PermissionUpdate) ClearPath() *PermissionUpdate {
	pu.mutation.ClearPath()
	return pu
}

// SetDescription sets the "description" field.
func (pu *PermissionUpdate) SetDescription(s string) *PermissionUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableDescription(s *string) *PermissionUpdate {
	if s != nil {
		pu.SetDescription(*s)
	}
	return pu
}

// ClearDescription clears the value of the "description" field.
func (pu *PermissionUpdate) ClearDescription() *PermissionUpdate {
	pu.mutation.ClearDescription()
	return pu
}

// SetSort sets the "sort" field.
func (pu *PermissionUpdate) SetSort(i int) *PermissionUpdate {
	pu.mutation.ResetSort()
	pu.mutation.SetSort(i)
	return pu
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (pu *PermissionUpdate) SetNillableSort(i *int) *PermissionUpdate {
	if i != nil {
		pu.SetSort(*i)
	}
	return pu
}

// AddSort adds i to the "sort" field.
func (pu *PermissionUpdate) AddSort(i int) *PermissionUpdate {
	pu.mutation.AddSort(i)
	return pu
}

// SetAttrs sets the "attrs" field.
func (pu *PermissionUpdate) SetAttrs(m map[string]interface{}) *PermissionUpdate {
	pu.mutation.SetAttrs(m)
	return pu
}

// ClearAttrs clears the value of the "attrs" field.
func (pu *PermissionUpdate) ClearAttrs() *PermissionUpdate {
	pu.mutation.ClearAttrs()
	return pu
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (pu *PermissionUpdate) AddRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddRoleIDs(ids...)
	return pu
}

// AddRoles adds the "roles" edges to the Role entity.
func (pu *PermissionUpdate) AddRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Permission entity.
func (pu *PermissionUpdate) SetParent(p *Permission) *PermissionUpdate {
	return pu.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the Permission entity by IDs.
func (pu *PermissionUpdate) AddChildIDs(ids ...int) *PermissionUpdate {
	pu.mutation.AddChildIDs(ids...)
	return pu
}

// AddChildren adds the "children" edges to the Permission entity.
func (pu *PermissionUpdate) AddChildren(p ...*Permission) *PermissionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddChildIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (pu *PermissionUpdate) Mutation() *PermissionMutation {
	return pu.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (pu *PermissionUpdate) ClearRoles() *PermissionUpdate {
	pu.mutation.ClearRoles()
	return pu
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (pu *PermissionUpdate) RemoveRoleIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveRoleIDs(ids...)
	return pu
}

// RemoveRoles removes "roles" edges to Role entities.
func (pu *PermissionUpdate) RemoveRoles(r ...*Role) *PermissionUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pu.RemoveRoleIDs(ids...)
}

// ClearParent clears the "parent" edge to the Permission entity.
func (pu *PermissionUpdate) ClearParent() *PermissionUpdate {
	pu.mutation.ClearParent()
	return pu
}

// ClearChildren clears all "children" edges to the Permission entity.
func (pu *PermissionUpdate) ClearChildren() *PermissionUpdate {
	pu.mutation.ClearChildren()
	return pu
}

// RemoveChildIDs removes the "children" edge to Permission entities by IDs.
func (pu *PermissionUpdate) RemoveChildIDs(ids ...int) *PermissionUpdate {
	pu.mutation.RemoveChildIDs(ids...)
	return pu
}

// RemoveChildren removes "children" edges to Permission entities.
func (pu *PermissionUpdate) RemoveChildren(p ...*Permission) *PermissionUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PermissionUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PermissionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PermissionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PermissionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *PermissionUpdate) defaults() {
	if _, ok := pu.mutation.UpdatedAt(); !ok && !pu.mutation.UpdatedAtCleared() {
		v := permission.UpdateDefaultUpdatedAt()
		pu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PermissionUpdate) check() error {
	if v, ok := pu.mutation.GetType(); ok {
		if err := permission.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Permission.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (pu *PermissionUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdate {
	pu.modifiers = append(pu.modifiers, modifiers...)
	return pu
}

func (pu *PermissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if pu.mutation.CreatedAtCleared() {
		_spec.ClearField(permission.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if pu.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Key(); ok {
		_spec.SetField(permission.FieldKey, field.TypeString, value)
	}
	if pu.mutation.KeyCleared() {
		_spec.ClearField(permission.FieldKey, field.TypeString)
	}
	if value, ok := pu.mutation.GetType(); ok {
		_spec.SetField(permission.FieldType, field.TypeEnum, value)
	}
	if value, ok := pu.mutation.Path(); ok {
		_spec.SetField(permission.FieldPath, field.TypeString, value)
	}
	if pu.mutation.PathCleared() {
		_spec.ClearField(permission.FieldPath, field.TypeString)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if pu.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := pu.mutation.Sort(); ok {
		_spec.SetField(permission.FieldSort, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedSort(); ok {
		_spec.AddField(permission.FieldSort, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Attrs(); ok {
		_spec.SetField(permission.FieldAttrs, field.TypeJSON, value)
	}
	if pu.mutation.AttrsCleared() {
		_spec.ClearField(permission.FieldAttrs, field.TypeJSON)
	}
	if pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedRolesIDs(); len(nodes) > 0 && !pu.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.ParentTable,
			Columns: []string{permission.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.ParentTable,
			Columns: []string{permission.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !pu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(pu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PermissionUpdateOne is the builder for updating a single Permission entity.
type PermissionUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PermissionMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (puo *PermissionUpdateOne) SetUpdatedAt(t time.Time) *PermissionUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (puo *PermissionUpdateOne) ClearUpdatedAt() *PermissionUpdateOne {
	puo.mutation.ClearUpdatedAt()
	return puo
}

// SetParentID sets the "parent_id" field.
func (puo *PermissionUpdateOne) SetParentID(i int) *PermissionUpdateOne {
	puo.mutation.SetParentID(i)
	return puo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableParentID(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetParentID(*i)
	}
	return puo
}

// ClearParentID clears the value of the "parent_id" field.
func (puo *PermissionUpdateOne) ClearParentID() *PermissionUpdateOne {
	puo.mutation.ClearParentID()
	return puo
}

// SetName sets the "name" field.
func (puo *PermissionUpdateOne) SetName(s string) *PermissionUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableName(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetKey sets the "key" field.
func (puo *PermissionUpdateOne) SetKey(s string) *PermissionUpdateOne {
	puo.mutation.SetKey(s)
	return puo
}

// SetNillableKey sets the "key" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableKey(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetKey(*s)
	}
	return puo
}

// ClearKey clears the value of the "key" field.
func (puo *PermissionUpdateOne) ClearKey() *PermissionUpdateOne {
	puo.mutation.ClearKey()
	return puo
}

// SetType sets the "type" field.
func (puo *PermissionUpdateOne) SetType(pe permission.Type) *PermissionUpdateOne {
	puo.mutation.SetType(pe)
	return puo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableType(pe *permission.Type) *PermissionUpdateOne {
	if pe != nil {
		puo.SetType(*pe)
	}
	return puo
}

// SetPath sets the "path" field.
func (puo *PermissionUpdateOne) SetPath(s string) *PermissionUpdateOne {
	puo.mutation.SetPath(s)
	return puo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillablePath(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetPath(*s)
	}
	return puo
}

// ClearPath clears the value of the "path" field.
func (puo *PermissionUpdateOne) ClearPath() *PermissionUpdateOne {
	puo.mutation.ClearPath()
	return puo
}

// SetDescription sets the "description" field.
func (puo *PermissionUpdateOne) SetDescription(s string) *PermissionUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableDescription(s *string) *PermissionUpdateOne {
	if s != nil {
		puo.SetDescription(*s)
	}
	return puo
}

// ClearDescription clears the value of the "description" field.
func (puo *PermissionUpdateOne) ClearDescription() *PermissionUpdateOne {
	puo.mutation.ClearDescription()
	return puo
}

// SetSort sets the "sort" field.
func (puo *PermissionUpdateOne) SetSort(i int) *PermissionUpdateOne {
	puo.mutation.ResetSort()
	puo.mutation.SetSort(i)
	return puo
}

// SetNillableSort sets the "sort" field if the given value is not nil.
func (puo *PermissionUpdateOne) SetNillableSort(i *int) *PermissionUpdateOne {
	if i != nil {
		puo.SetSort(*i)
	}
	return puo
}

// AddSort adds i to the "sort" field.
func (puo *PermissionUpdateOne) AddSort(i int) *PermissionUpdateOne {
	puo.mutation.AddSort(i)
	return puo
}

// SetAttrs sets the "attrs" field.
func (puo *PermissionUpdateOne) SetAttrs(m map[string]interface{}) *PermissionUpdateOne {
	puo.mutation.SetAttrs(m)
	return puo
}

// ClearAttrs clears the value of the "attrs" field.
func (puo *PermissionUpdateOne) ClearAttrs() *PermissionUpdateOne {
	puo.mutation.ClearAttrs()
	return puo
}

// AddRoleIDs adds the "roles" edge to the Role entity by IDs.
func (puo *PermissionUpdateOne) AddRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddRoleIDs(ids...)
	return puo
}

// AddRoles adds the "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) AddRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.AddRoleIDs(ids...)
}

// SetParent sets the "parent" edge to the Permission entity.
func (puo *PermissionUpdateOne) SetParent(p *Permission) *PermissionUpdateOne {
	return puo.SetParentID(p.ID)
}

// AddChildIDs adds the "children" edge to the Permission entity by IDs.
func (puo *PermissionUpdateOne) AddChildIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.AddChildIDs(ids...)
	return puo
}

// AddChildren adds the "children" edges to the Permission entity.
func (puo *PermissionUpdateOne) AddChildren(p ...*Permission) *PermissionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddChildIDs(ids...)
}

// Mutation returns the PermissionMutation object of the builder.
func (puo *PermissionUpdateOne) Mutation() *PermissionMutation {
	return puo.mutation
}

// ClearRoles clears all "roles" edges to the Role entity.
func (puo *PermissionUpdateOne) ClearRoles() *PermissionUpdateOne {
	puo.mutation.ClearRoles()
	return puo
}

// RemoveRoleIDs removes the "roles" edge to Role entities by IDs.
func (puo *PermissionUpdateOne) RemoveRoleIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveRoleIDs(ids...)
	return puo
}

// RemoveRoles removes "roles" edges to Role entities.
func (puo *PermissionUpdateOne) RemoveRoles(r ...*Role) *PermissionUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return puo.RemoveRoleIDs(ids...)
}

// ClearParent clears the "parent" edge to the Permission entity.
func (puo *PermissionUpdateOne) ClearParent() *PermissionUpdateOne {
	puo.mutation.ClearParent()
	return puo
}

// ClearChildren clears all "children" edges to the Permission entity.
func (puo *PermissionUpdateOne) ClearChildren() *PermissionUpdateOne {
	puo.mutation.ClearChildren()
	return puo
}

// RemoveChildIDs removes the "children" edge to Permission entities by IDs.
func (puo *PermissionUpdateOne) RemoveChildIDs(ids ...int) *PermissionUpdateOne {
	puo.mutation.RemoveChildIDs(ids...)
	return puo
}

// RemoveChildren removes "children" edges to Permission entities.
func (puo *PermissionUpdateOne) RemoveChildren(p ...*Permission) *PermissionUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveChildIDs(ids...)
}

// Where appends a list predicates to the PermissionUpdate builder.
func (puo *PermissionUpdateOne) Where(ps ...predicate.Permission) *PermissionUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PermissionUpdateOne) Select(field string, fields ...string) *PermissionUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Permission entity.
func (puo *PermissionUpdateOne) Save(ctx context.Context) (*Permission, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PermissionUpdateOne) SaveX(ctx context.Context) *Permission {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PermissionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PermissionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *PermissionUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdatedAt(); !ok && !puo.mutation.UpdatedAtCleared() {
		v := permission.UpdateDefaultUpdatedAt()
		puo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PermissionUpdateOne) check() error {
	if v, ok := puo.mutation.GetType(); ok {
		if err := permission.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Permission.type": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (puo *PermissionUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *PermissionUpdateOne {
	puo.modifiers = append(puo.modifiers, modifiers...)
	return puo
}

func (puo *PermissionUpdateOne) sqlSave(ctx context.Context) (_node *Permission, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(permission.Table, permission.Columns, sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Permission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, permission.FieldID)
		for _, f := range fields {
			if !permission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != permission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if puo.mutation.CreatedAtCleared() {
		_spec.ClearField(permission.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(permission.FieldUpdatedAt, field.TypeTime, value)
	}
	if puo.mutation.UpdatedAtCleared() {
		_spec.ClearField(permission.FieldUpdatedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(permission.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Key(); ok {
		_spec.SetField(permission.FieldKey, field.TypeString, value)
	}
	if puo.mutation.KeyCleared() {
		_spec.ClearField(permission.FieldKey, field.TypeString)
	}
	if value, ok := puo.mutation.GetType(); ok {
		_spec.SetField(permission.FieldType, field.TypeEnum, value)
	}
	if value, ok := puo.mutation.Path(); ok {
		_spec.SetField(permission.FieldPath, field.TypeString, value)
	}
	if puo.mutation.PathCleared() {
		_spec.ClearField(permission.FieldPath, field.TypeString)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(permission.FieldDescription, field.TypeString, value)
	}
	if puo.mutation.DescriptionCleared() {
		_spec.ClearField(permission.FieldDescription, field.TypeString)
	}
	if value, ok := puo.mutation.Sort(); ok {
		_spec.SetField(permission.FieldSort, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedSort(); ok {
		_spec.AddField(permission.FieldSort, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Attrs(); ok {
		_spec.SetField(permission.FieldAttrs, field.TypeJSON, value)
	}
	if puo.mutation.AttrsCleared() {
		_spec.ClearField(permission.FieldAttrs, field.TypeJSON)
	}
	if puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedRolesIDs(); len(nodes) > 0 && !puo.mutation.RolesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   permission.RolesTable,
			Columns: permission.RolesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(role.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.ParentTable,
			Columns: []string{permission.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   permission.ParentTable,
			Columns: []string{permission.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !puo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   permission.ChildrenTable,
			Columns: []string{permission.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(permission.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(puo.modifiers...)
	_node = &Permission{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{permission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
