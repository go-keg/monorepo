package graphql

import (
	"context"
	"fmt"
	"image/color"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-keg/keg/contrib/cache"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/keg/contrib/gql/pubsub"
	zlog "github.com/go-keg/keg/contrib/log"
	"github.com/go-keg/monorepo/internal/app/example/biz"
	"github.com/go-keg/monorepo/internal/app/example/server/auth"
	"github.com/go-keg/monorepo/internal/app/example/service/graphql/model"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/exp/slices"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	log         *log.Helper
	db          *ent.Database // 使用事务时使用 Database
	ent         *ent.Client
	userUseCase *biz.UserUseCase
	userRepo    biz.UserRepo
	captcha     *base64Captcha.Captcha
	oauth       *biz.OAuthUseCase
	pubSub      *pubsub.PubSub[*model.Message]
}

func NewResolver(
	logger log.Logger,
	db *ent.Database,
	client *ent.Client,
	uc *biz.UserUseCase,
	ur biz.UserRepo,
	oauth *biz.OAuthUseCase,
) *Resolver {
	return &Resolver{
		log:         log.NewHelper(log.With(logger, "module", "service/graphql", "traceId", zlog.TraceID())),
		db:          db,
		ent:         client,
		userUseCase: uc,
		userRepo:    ur,
		captcha: base64Captcha.NewCaptcha(
			base64Captcha.NewDriverString(
				40,
				140,
				0,
				0,
				4,
				"1234567890abcdefghijklmnopqrktuvwxyz",
				&color.RGBA{},
				base64Captcha.DefaultEmbeddedFonts,
				nil),
			base64Captcha.DefaultMemStore,
		),
		oauth:  oauth,
		pubSub: pubsub.New[*model.Message](),
	}
}

func NewDirectiveRoot(ur biz.UserRepo) DirectiveRoot {
	return DirectiveRoot{
		Login: func(ctx context.Context, obj any, next graphql.Resolver) (res any, err error) {
			if auth.GetUser(ctx) != nil {
				return next(ctx)
			}
			return nil, gql.ErrUnauthorized
		},
		HasPermission: func(ctx context.Context, obj any, next graphql.Resolver, key string) (res any, err error) {
			u := auth.GetTenantUser(ctx)
			if u == nil {
				return nil, gql.ErrUnauthorized
			}
			if !u.IsOwner {
				keys, err := cache.LocalRemember(fmt.Sprintf("tenant_user:%d:permissions", u.ID), time.Minute*10, func() ([]string, error) {
					return ur.GetTenantUserPermissionKeys(ctx, u.ID)
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
		HasRole: func(ctx context.Context, obj any, next graphql.Resolver, role []model.UserRole) (res any, err error) {
			for _, userRole := range role {
				switch userRole {
				case model.UserRoleTenantUser:
					if auth.GetTenantUser(ctx) != nil {
						return next(ctx)
					}
				case model.UserRoleTenantOwner:
					if u := auth.GetTenantUser(ctx); u != nil && u.IsOwner {
						return next(ctx)
					}
				case model.UserRoleSystemAdmin:
					if u := auth.GetUser(ctx); u != nil && u.IsAdmin {
						return next(ctx)
					}
				}
			}
			return nil, gql.ErrAccessDenied
		},
	}
}

// NewSchema creates a graphql executable schema.
func NewSchema(
	resolver *Resolver,
	directive DirectiveRoot,
) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers:  resolver,
		Directives: directive,
	})
}
