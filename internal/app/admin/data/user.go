package data

import (
	"context"

	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/monorepo/internal/data/example/ent/oauthaccount"
	"github.com/go-keg/monorepo/internal/data/example/ent/permission"
	"github.com/go-keg/monorepo/internal/data/example/ent/role"
	"github.com/go-keg/monorepo/internal/data/example/ent/user"
	"github.com/samber/lo"
)

type userRepo struct {
	ent *ent.Client
}

func (r userRepo) UnBindOAuthAccount(ctx context.Context, userID int, provider string) error {
	_, err := r.ent.OAuthAccount.Delete().Where(
		oauthaccount.UserID(userID),
		oauthaccount.Provider(provider),
	).Exec(ctx)
	return err
}

func (r userRepo) FindUserByOAuth(ctx context.Context, provider, providerUserID string) (*ent.User, error) {
	return r.ent.User.Query().Where(user.HasOauthAccountsWith(
		oauthaccount.Provider(provider),
		oauthaccount.ProviderUserID(providerUserID),
	)).First(ctx)
}

func (r userRepo) FindUserByEmail(ctx context.Context, email string) (*ent.User, error) {
	return r.ent.User.Query().Where(user.Email(email)).First(ctx)
}

func (r userRepo) BindOAuthAccount(ctx context.Context, data *ent.OAuthAccount) error {
	err := r.UnBindOAuthAccount(ctx, data.UserID, data.Provider)
	if err != nil {
		return err
	}
	return r.ent.OAuthAccount.Create().
		SetUserID(data.UserID).
		SetProvider(data.Provider).
		SetProviderUserID(data.ProviderUserID).
		SetNillableAccessToken(lo.EmptyableToPtr(data.AccessToken)).
		SetNillableRefreshToken(lo.EmptyableToPtr(data.RefreshToken)).
		SetNillableTokenExpiry(lo.EmptyableToPtr(data.TokenExpiry)).
		SetProfile(data.Profile).
		Exec(ctx)
}

func (r userRepo) CreateUser(ctx context.Context, user *ent.User) (*ent.User, error) {
	return r.ent.User.Create().
		SetEmail(user.Email).
		SetNickname(user.Nickname).
		SetNillableAvatar(lo.EmptyableToPtr(user.Avatar)).
		SetNillablePassword(lo.EmptyableToPtr(user.Password)).
		Save(ctx)
}

func NewUserRepo(ent *ent.Client) biz.UserRepo {
	return &userRepo{ent: ent}
}

func (r userRepo) GetUserPermissionKeys(ctx context.Context, userID int) ([]string, error) {
	return r.ent.Permission.Query().
		Where(permission.HasRolesWith(role.HasUsersWith(user.ID(userID)))).
		Select(permission.FieldKey).Strings(ctx)
}
