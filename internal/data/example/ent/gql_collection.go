// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/go-keg/monorepo/internal/data/example/ent/account"
	"github.com/go-keg/monorepo/internal/data/example/ent/operationlog"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AccountQuery) CollectFields(ctx context.Context, satisfies ...string) (*AccountQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AccountQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(account.Columns))
		selectedFields = []string{account.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "nickname":
			if _, ok := fieldSeen[account.FieldNickname]; !ok {
				selectedFields = append(selectedFields, account.FieldNickname)
				fieldSeen[account.FieldNickname] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type accountPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AccountPaginateOption
}

func newAccountPaginateArgs(rv map[string]any) *accountPaginateArgs {
	args := &accountPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*AccountWhereInput); ok {
		args.opts = append(args.opts, WithAccountFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ol *OperationLogQuery) CollectFields(ctx context.Context, satisfies ...string) (*OperationLogQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ol, nil
	}
	if err := ol.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ol, nil
}

func (ol *OperationLogQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(operationlog.Columns))
		selectedFields = []string{operationlog.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: ol.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			ol.withUser = query
			if _, ok := fieldSeen[operationlog.FieldUserID]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldUserID)
				fieldSeen[operationlog.FieldUserID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[operationlog.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldCreatedAt)
				fieldSeen[operationlog.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[operationlog.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldUpdatedAt)
				fieldSeen[operationlog.FieldUpdatedAt] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[operationlog.FieldUserID]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldUserID)
				fieldSeen[operationlog.FieldUserID] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[operationlog.FieldType]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldType)
				fieldSeen[operationlog.FieldType] = struct{}{}
			}
		case "context":
			if _, ok := fieldSeen[operationlog.FieldContext]; !ok {
				selectedFields = append(selectedFields, operationlog.FieldContext)
				fieldSeen[operationlog.FieldContext] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		ol.Select(selectedFields...)
	}
	return nil
}

type operationlogPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []OperationLogPaginateOption
}

func newOperationLogPaginateArgs(rv map[string]any) *operationlogPaginateArgs {
	args := &operationlogPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &OperationLogOrder{Field: &OperationLogOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithOperationLogOrder(order))
			}
		case *OperationLogOrder:
			if v != nil {
				args.opts = append(args.opts, WithOperationLogOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*OperationLogWhereInput); ok {
		args.opts = append(args.opts, WithOperationLogFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pe *PermissionQuery) CollectFields(ctx context.Context, satisfies ...string) (*PermissionQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pe, nil
	}
	if err := pe.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pe *PermissionQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(permission.Columns))
		selectedFields = []string{permission.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: pe.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, permissionImplementors)...); err != nil {
				return err
			}
			pe.withParent = query
			if _, ok := fieldSeen[permission.FieldParentID]; !ok {
				selectedFields = append(selectedFields, permission.FieldParentID)
				fieldSeen[permission.FieldParentID] = struct{}{}
			}

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: pe.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, permissionImplementors)...); err != nil {
				return err
			}
			pe.WithNamedChildren(alias, func(wq *PermissionQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[permission.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldCreatedAt)
				fieldSeen[permission.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[permission.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, permission.FieldUpdatedAt)
				fieldSeen[permission.FieldUpdatedAt] = struct{}{}
			}
		case "parentID":
			if _, ok := fieldSeen[permission.FieldParentID]; !ok {
				selectedFields = append(selectedFields, permission.FieldParentID)
				fieldSeen[permission.FieldParentID] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[permission.FieldName]; !ok {
				selectedFields = append(selectedFields, permission.FieldName)
				fieldSeen[permission.FieldName] = struct{}{}
			}
		case "key":
			if _, ok := fieldSeen[permission.FieldKey]; !ok {
				selectedFields = append(selectedFields, permission.FieldKey)
				fieldSeen[permission.FieldKey] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[permission.FieldType]; !ok {
				selectedFields = append(selectedFields, permission.FieldType)
				fieldSeen[permission.FieldType] = struct{}{}
			}
		case "path":
			if _, ok := fieldSeen[permission.FieldPath]; !ok {
				selectedFields = append(selectedFields, permission.FieldPath)
				fieldSeen[permission.FieldPath] = struct{}{}
			}
		case "desc":
			if _, ok := fieldSeen[permission.FieldDesc]; !ok {
				selectedFields = append(selectedFields, permission.FieldDesc)
				fieldSeen[permission.FieldDesc] = struct{}{}
			}
		case "sort":
			if _, ok := fieldSeen[permission.FieldSort]; !ok {
				selectedFields = append(selectedFields, permission.FieldSort)
				fieldSeen[permission.FieldSort] = struct{}{}
			}
		case "attrs":
			if _, ok := fieldSeen[permission.FieldAttrs]; !ok {
				selectedFields = append(selectedFields, permission.FieldAttrs)
				fieldSeen[permission.FieldAttrs] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pe.Select(selectedFields...)
	}
	return nil
}

type permissionPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PermissionPaginateOption
}

func newPermissionPaginateArgs(rv map[string]any) *permissionPaginateArgs {
	args := &permissionPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PermissionOrder{Field: &PermissionOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPermissionOrder(order))
			}
		case *PermissionOrder:
			if v != nil {
				args.opts = append(args.opts, WithPermissionOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PermissionWhereInput); ok {
		args.opts = append(args.opts, WithPermissionFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (r *RoleQuery) CollectFields(ctx context.Context, satisfies ...string) (*RoleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return r, nil
	}
	if err := r.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return r, nil
}

func (r *RoleQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(role.Columns))
		selectedFields = []string{role.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "permissions":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PermissionClient{config: r.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, permissionImplementors)...); err != nil {
				return err
			}
			r.WithNamedPermissions(alias, func(wq *PermissionQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[role.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldCreatedAt)
				fieldSeen[role.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[role.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, role.FieldUpdatedAt)
				fieldSeen[role.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[role.FieldName]; !ok {
				selectedFields = append(selectedFields, role.FieldName)
				fieldSeen[role.FieldName] = struct{}{}
			}
		case "sort":
			if _, ok := fieldSeen[role.FieldSort]; !ok {
				selectedFields = append(selectedFields, role.FieldSort)
				fieldSeen[role.FieldSort] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		r.Select(selectedFields...)
	}
	return nil
}

type rolePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []RolePaginateOption
}

func newRolePaginateArgs(rv map[string]any) *rolePaginateArgs {
	args := &rolePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &RoleOrder{Field: &RoleOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithRoleOrder(order))
			}
		case *RoleOrder:
			if v != nil {
				args.opts = append(args.opts, WithRoleOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*RoleWhereInput); ok {
		args.opts = append(args.opts, WithRoleFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "roles":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&RoleClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, roleImplementors)...); err != nil {
				return err
			}
			u.WithNamedRoles(alias, func(wq *RoleQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "nickname":
			if _, ok := fieldSeen[user.FieldNickname]; !ok {
				selectedFields = append(selectedFields, user.FieldNickname)
				fieldSeen[user.FieldNickname] = struct{}{}
			}
		case "avatar":
			if _, ok := fieldSeen[user.FieldAvatar]; !ok {
				selectedFields = append(selectedFields, user.FieldAvatar)
				fieldSeen[user.FieldAvatar] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[user.FieldStatus]; !ok {
				selectedFields = append(selectedFields, user.FieldStatus)
				fieldSeen[user.FieldStatus] = struct{}{}
			}
		case "isAdmin":
			if _, ok := fieldSeen[user.FieldIsAdmin]; !ok {
				selectedFields = append(selectedFields, user.FieldIsAdmin)
				fieldSeen[user.FieldIsAdmin] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok || v == nil {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
