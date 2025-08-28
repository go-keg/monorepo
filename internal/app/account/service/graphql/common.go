package graphql

import (
	"context"
	"fmt"

	"github.com/go-keg/keg/contrib/cache"
	"github.com/go-keg/monorepo/internal/app/account/service/graphql/model"
	"github.com/go-keg/monorepo/internal/data/account/ent"
)

func (r *Resolver) LoginReply(_ context.Context, user *ent.User) (*model.LoginReply, error) {
	token, exp, err := r.userUseCase.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Clear Permission cache
	for _, tenantUser := range user.Edges.TenantUsers {
		cache.LocalClear(fmt.Sprintf("tenant_user:%d:permissions", tenantUser.ID))
	}
	return &model.LoginReply{
		Token: token,
		Exp:   exp,
		User:  user,
	}, nil
}
