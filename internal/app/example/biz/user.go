package biz

import (
	"context"
	"fmt"
	"time"

	"github.com/go-keg/keg/contrib/cache"
	"github.com/go-keg/monorepo/internal/app/example/conf"
	"github.com/go-keg/monorepo/internal/app/example/service/graphql/model"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"gopkg.in/gomail.v2"
)

type UserRepo interface {
	FindUserByOAuth(ctx context.Context, provider string, providerUserID string) (*ent.User, error)
	FindUserByEmail(ctx context.Context, email string) (*ent.User, error)
	BindOAuthAccount(ctx context.Context, data ent.OAuthAccount) error
	UnBindOAuthAccount(ctx context.Context, userID int, provider string) error
	CreateUser(ctx context.Context, user ent.User) (*ent.User, error)
	GetTenantUserPermissionKeys(ctx context.Context, tenantUserID int) ([]string, error)
}
type UserUseCase struct {
	cfg    *conf.Config
	dialer *gomail.Dialer
}

func NewUserUseCase(cfg *conf.Config) *UserUseCase {
	return &UserUseCase{
		cfg:    cfg,
		dialer: gomail.NewDialer(cfg.Email.Host, cast.ToInt(cfg.Email.Port), cfg.Email.Username, cfg.Email.Password),
	}
}

func (r UserUseCase) GenerateToken(userID int) (string, int, error) {
	ttl := time.Hour * 48
	exp := time.Now().Add(ttl).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": cast.ToString(userID),
		"exp": exp,               // Expiration Time
		"iat": time.Now().Unix(), // Issued At OPTIONAL
	}).SignedString([]byte(r.cfg.Key))
	if err != nil {
		return "", 0, err
	}
	return token, int(ttl / time.Second), nil
}

func (r UserUseCase) SendEmail(email string, emailType model.VerifyCodeType) error {
	code := lo.RandomString(6, lo.NumbersCharset)
	cache.LocalSet(fmt.Sprintf("send_email:%s:%s", emailType, email), code, time.Minute*15)
	m := gomail.NewMessage()
	m.SetHeader("From", r.cfg.Email.From)
	m.SetHeader("To", email)
	switch emailType {
	case model.VerifyCodeTypeRegister:
		m.SetHeader("Subject", fmt.Sprintf("Register - %s", r.cfg.Name))
		m.SetBody("text/html", fmt.Sprintf("verify code: %s", code))
	case model.VerifyCodeTypeForgetPassword:
		m.SetHeader("Subject", fmt.Sprintf("Forget Password - %s", r.cfg.Name))
		m.SetBody("text/html", fmt.Sprintf("verify code: %s", code))
	}
	return r.dialer.DialAndSend(m)
}

func (r UserUseCase) CheckEmailVerifyCode(email string, emailType model.VerifyCodeType, code string) bool {
	v, ok := cache.LocalGet(fmt.Sprintf("send_email:%s:%s", emailType, email))
	if !ok {
		return false
	}
	return cast.ToString(v) == code
}
