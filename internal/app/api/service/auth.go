package service

import (
	"context"
	v1 "github.com/go-keg/apis/api/account/v1"
	"github.com/go-keg/apis/api/common"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
)

type AuthService struct {
	v1.UnimplementedAuthServiceServer
	ent *ent.Client
}

func NewAuthService(ent *ent.Client) *AuthService {
	return &AuthService{ent: ent}
}

func (r AuthService) Permissions(ctx context.Context, req *common.ID) (*v1.PermissionsReply, error) {
	permissionKeys, err := r.ent.Permission.Query().Where(
		permission.HasRolesWith(
			role.HasUsersWith(user.ID(int(req.Id))),
		),
	).Select(permission.FieldKey).Strings(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.PermissionsReply{Keywords: permissionKeys}, nil
}
