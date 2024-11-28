package graphql

import (
	"context"
	"fmt"
	"image/color"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-keg/keg/contrib/cache"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/app/admin/server/auth"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/exp/slices"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log            *log.Helper
	db             *ent.Database // 使用事务时使用 Database
	ent            *ent.Client
	accountUseCase *biz.AccountUseCase
	captcha        *base64Captcha.Captcha
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger log.Logger, db *ent.Database, client *ent.Client, accountUseCase *biz.AccountUseCase) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:            log.NewHelper(log.With(logger, "module", "service/graphql")),
			db:             db,
			ent:            client,
			accountUseCase: accountUseCase,
			captcha:        base64Captcha.NewCaptcha(base64Captcha.NewDriverString(40, 140, 0, 0, 4, "1234567890abcdefghijklmnopqrktuvwxyz", &color.RGBA{}, base64Captcha.DefaultEmbeddedFonts, nil), base64Captcha.DefaultMemStore),
		},
		Directives: DirectiveRoot{
			Disabled: func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
				return nil, gql.ErrDisabled
			},
			Login: func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
				if auth.GetUser(ctx) != nil {
					return next(ctx)
				}
				return nil, gql.ErrUnauthorized
			},
			HasPermission: func(ctx context.Context, obj any, next graphql.Resolver, key string) (res any, err error) {
				u := auth.GetUser(ctx)
				if u == nil {
					return nil, gql.ErrUnauthorized
				}
				if !u.IsAdmin {
					keys, err := cache.LocalRemember(fmt.Sprintf("user:%d:permissions", u.ID), time.Minute*5, func() ([]string, error) {
						return client.Permission.Query().Where(permission.HasRolesWith(role.HasUsersWith(user.ID(u.ID)))).Select(permission.FieldKey).Strings(ctx)
					})
					if err != nil {
						return nil, err
					}
					if !slices.Contains(keys, key) {
						return res, gql.ErrNoPermission
					}
				}
				return next(ctx)
			},
		},
	})
}
