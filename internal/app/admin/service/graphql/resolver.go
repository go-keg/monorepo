package graphql

import (
	"context"
	"encoding/json"
	"fmt"
	"image/color"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-keg/keg/contrib/cache"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/app/admin/server/auth"
	"github.com/go-keg/monorepo/internal/app/admin/service/graphql/model"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	"golang.org/x/exp/slices"
	"golang.org/x/oauth2"
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
	googleOAuth *oauth2.Config
}

// NewSchema creates a graphql executable schema.
func NewSchema(
	googleOAuth *oauth2.Config,
	logger log.Logger,
	db *ent.Database,
	client *ent.Client,
	userUseCase *biz.UserUseCase,
	userRepo biz.UserRepo,
) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:         log.NewHelper(log.With(logger, "module", "service/graphql")),
			db:          db,
			ent:         client,
			userUseCase: userUseCase,
			userRepo:    userRepo,
			captcha:     base64Captcha.NewCaptcha(base64Captcha.NewDriverString(40, 140, 0, 0, 4, "1234567890abcdefghijklmnopqrktuvwxyz", &color.RGBA{}, base64Captcha.DefaultEmbeddedFonts, nil), base64Captcha.DefaultMemStore),
			googleOAuth: googleOAuth,
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
			Permission: func(ctx context.Context, obj any, next graphql.Resolver, key string) (res any, err error) {
				u := auth.GetUser(ctx)
				if u == nil {
					return nil, gql.ErrUnauthorized
				}
				if !u.IsAdmin {
					keys, err := cache.LocalRemember(fmt.Sprintf("user:%d:permissions", u.ID), time.Minute*10, func() ([]string, error) {
						return userRepo.GetUserPermissionKeys(ctx, u.ID)
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

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func (r GoogleUserInfo) Map() map[string]any {
	data, _ := json.Marshal(r)
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}

func (r *Resolver) GoogleUserInfo(ctx context.Context, code string) (*GoogleUserInfo, error) {
	token, err := r.googleOAuth.Exchange(ctx, code)
	if err != nil {
		return nil, gql.Error("code exchange failed")
	}

	client := r.googleOAuth.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, gql.Error("failed to get user info")
	}
	defer resp.Body.Close()
	var userInfo GoogleUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, gql.Error("failed to parse user info")
	}
	return &userInfo, nil
}

func (r *Resolver) LoginReply(_ context.Context, user *ent.User) (*model.LoginReply, error) {
	token, exp, err := r.userUseCase.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	// Clear Permission cache
	cache.LocalClear(fmt.Sprintf("user:%d:permissions", user.ID))

	return &model.LoginReply{
		Token: token,
		Exp:   int(exp),
		User:  user,
	}, nil
}
