package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70-dev

import (
	"context"

	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/admin/server/auth"
	"github.com/go-keg/monorepo/internal/app/admin/service/graphql/model"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/oauthaccount"
)

// LinkGoogleAccount is the resolver for the linkGoogleAccount field.
func (r *mutationResolver) LinkGoogleAccount(ctx context.Context, code string) (bool, error) {
	info, err := r.GoogleUserInfo(ctx, code)
	if err != nil {
		return false, err
	}
	err = r.userRepo.BindOAuthAccount(ctx, &ent.OAuthAccount{
		UserID:         auth.GetUserID(ctx),
		Provider:       oauthaccount.ProviderGoogle,
		ProviderUserID: info.ID,
		Profile:        info.Map(),
	})
	if err != nil {
		return false, gql.Error("failed to link google account")
	}
	return true, nil
}

// UnlinkGoogleAccount is the resolver for the unlinkGoogleAccount field.
func (r *mutationResolver) UnlinkGoogleAccount(ctx context.Context) (bool, error) {
	err := r.userRepo.UnBindOAuthAccount(ctx, auth.GetUserID(ctx), oauthaccount.ProviderGoogle)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GoogleOAuthLogin is the resolver for the googleOAuthLogin field.
func (r *queryResolver) GoogleOAuthLogin(ctx context.Context, code string) (*model.LoginReply, error) {
	info, err := r.GoogleUserInfo(ctx, code)
	if err != nil {
		return nil, err
	}
	user, err := r.userRepo.FindUserByOAuth(ctx, oauthaccount.ProviderGoogle, info.ID)
	if ent.IsNotFound(err) {
		return nil, gql.Error("failed to find user")
	} else if err != nil {
		return nil, err
	}
	return r.LoginReply(ctx, user)
}

// GoogleOAuthRegister is the resolver for the googleOAuthRegister field.
func (r *queryResolver) GoogleOAuthRegister(ctx context.Context, code string) (*model.LoginReply, error) {
	info, err := r.GoogleUserInfo(ctx, code)
	if err != nil {
		return nil, err
	}
	_, err = r.userRepo.FindUserByEmail(ctx, info.Email)
	if ent.IsNotFound(err) {
		// 创建新用户并绑定
		user, err := r.userRepo.CreateUser(ctx, &ent.User{
			Email:    info.Email,
			Nickname: info.Name,
			Avatar:   info.Picture,
		})
		if err != nil {
			return nil, err
		}
		return r.LoginReply(ctx, user)
	} else if err != nil {
		return nil, err
	}
	return nil, gql.Error("failed to register user")
}
