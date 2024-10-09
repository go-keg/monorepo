package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/go-keg/example/internal/app/admin/service/graphql/dataloader"
	"github.com/go-keg/example/internal/data/example/ent"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	input.Password = r.accountUseCase.GeneratePassword(input.Password)
	return r.db.User(ctx).Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	if input.Password != nil {
		hashed := r.accountUseCase.GeneratePassword(*input.Password)
		input.Password = &hashed
	}
	return r.db.User(ctx).UpdateOneID(id).SetInput(input).Save(ctx)
}

// CreateRole is the resolver for the createRole field.
func (r *mutationResolver) CreateRole(ctx context.Context, input ent.CreateRoleInput) (*ent.Role, error) {
	return r.db.Role(ctx).Create().SetInput(input).Save(ctx)
}

// UpdateRole is the resolver for the updateRole field.
func (r *mutationResolver) UpdateRole(ctx context.Context, id int, input ent.UpdateRoleInput) (*ent.Role, error) {
	return r.db.Role(ctx).UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeleteRole is the resolver for the deleteRole field.
func (r *mutationResolver) DeleteRole(ctx context.Context, id int) (bool, error) {
	err := r.db.Role(ctx).DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CreatePermission is the resolver for the createPermission field.
func (r *mutationResolver) CreatePermission(ctx context.Context, input ent.CreatePermissionInput) (*ent.Permission, error) {
	return r.db.Permission(ctx).Create().SetInput(input).Save(ctx)
}

// UpdatePermission is the resolver for the updatePermission field.
func (r *mutationResolver) UpdatePermission(ctx context.Context, id int, input ent.UpdatePermissionInput) (*ent.Permission, error) {
	return r.db.Permission(ctx).UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePermission is the resolver for the deletePermission field.
func (r *mutationResolver) DeletePermission(ctx context.Context, id int) (bool, error) {
	err := r.db.Permission(ctx).DeleteOneID(id).Exec(ctx)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ChildrenCount is the resolver for the childrenCount field.
func (r *permissionResolver) ChildrenCount(ctx context.Context, obj *ent.Permission) (int, error) {
	return dataloader.For(ctx).GetPermissionChildrenCount(ctx, obj.ID)
}
