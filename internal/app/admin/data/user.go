package data

import (
	"context"

	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
)

type userRepo struct {
	ent *ent.Client
}

func NewUserRepo(ent *ent.Client) biz.UserRepo {
	return &userRepo{ent: ent}
}

func (r userRepo) GetUserPermissionKeys(ctx context.Context, userID int) ([]string, error) {
	return r.ent.Permission.Query().
		Where(permission.HasRolesWith(role.HasUsersWith(user.ID(userID)))).
		Select(permission.FieldKey).Strings(ctx)
}
