package biz

import (
	"context"
	"fmt"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/app/admin/service/graphql/model"
	"github.com/go-keg/example/internal/data/example/ent"
	"github.com/go-keg/keg/contrib/cache"
	"github.com/golang-jwt/jwt/v5"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
	"time"
)

type AccountUseCase struct {
	cfg    *conf.Config
	client *ent.Client
	dialer *gomail.Dialer
}

func NewAccountUseCase(cfg *conf.Config) *AccountUseCase {
	return &AccountUseCase{
		cfg:    cfg,
		dialer: gomail.NewDialer(cfg.Email.Host, cast.ToInt(cfg.Email.Port), cfg.Email.Username, cfg.Email.Password),
	}
}

func (r AccountUseCase) GenerateToken(_ context.Context, userId int) (string, int64, error) {
	exp := time.Now().Add(time.Hour * 24 * 7).Unix()
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": cast.ToString(userId),
		"exp": exp,               // Expiration Time
		"iat": time.Now().Unix(), // Issued At OPTIONAL
	}).SignedString([]byte(r.cfg.Key))
	if err != nil {
		return "", 0, err
	}
	return token, exp, nil
}

func (r AccountUseCase) VerifyPassword(account *ent.User, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password)) == nil
}

func (r AccountUseCase) GeneratePassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(hashed)
}

func (r AccountUseCase) SendEmail(email string, emailType model.VerifyCodeType) error {
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

func (r AccountUseCase) CheckEmailVerifyCode(email string, emailType model.VerifyCodeType, code string) bool {
	v, ok := cache.LocalGet(fmt.Sprintf("send_email:%s:%s", emailType, email))
	if !ok {
		return false
	}
	return cast.ToString(v) == code
}
