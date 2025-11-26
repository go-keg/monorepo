package conf

import (
	"image/color"

	"github.com/go-keg/keg/contrib/config"
	"github.com/google/wire"
	"github.com/mojocn/base64Captcha"
	"github.com/spf13/cast"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"gopkg.in/gomail.v2"
)

var ProviderSet = wire.NewSet(
	NewGoogleOAuthConfig,
	NewCaptcha,
)

type Config struct {
	Key    string
	Name   string
	Server struct {
		HTTP config.Server `yaml:"http"`
	}
	Data struct {
		Database config.Database
		Kafka    config.Kafka
	}
	OAuth struct {
		Google struct {
			RedirectURL  string
			ClientID     string
			ClientSecret string
		}
	}
	KafkaConsumerGroup config.KafkaConsumerGroup
	Email              config.Email
	Trace              struct {
		Endpoint string
	}
	Log config.Log
}

func Load(path string, envs ...string) (*Config, error) {
	return config.Load[Config](path, envs...)
}

func MustLoad(path string, envs ...string) *Config {
	cfg, err := Load(path, envs...)
	if err != nil {
		panic(err)
	}
	return cfg
}

func NewGoogleOAuthConfig(cfg *Config) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  cfg.OAuth.Google.RedirectURL,
		ClientID:     cfg.OAuth.Google.ClientID,
		ClientSecret: cfg.OAuth.Google.ClientSecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"openid",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func NewCaptcha() *base64Captcha.Captcha {
	return base64Captcha.NewCaptcha(
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
	)
}

func NewMail(cfg *Config) *gomail.Dialer {
	return gomail.NewDialer(
		cfg.Email.Host,
		cast.ToInt(cfg.Email.Port),
		cfg.Email.Username,
		cfg.Email.Password,
	)
}
