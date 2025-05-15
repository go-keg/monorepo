package data

import (
	"github.com/go-keg/keg/contrib/ent/driver"
	"github.com/go-keg/monorepo/internal/app/admin/conf"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/google/wire"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
	NewEntDatabase,
	NewUserRepo,
	NewGoogleOAuthConfig,
	// NewKafkaProducer,
)

func NewEntClient(cfg *conf.Config) (*ent.Client, error) {
	drv, err := driver.NewDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}

func NewEntDatabase(cfg *conf.Config) (*ent.Database, error) {
	drv, err := driver.NewDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewDatabase(ent.Driver(drv)), nil
}

// func NewKafkaProducer(cfg *conf.Config) (sarama.SyncProducer, error) {
// 	return kafka.NewSyncProducerFromConfig(cfg.Data.Kafka)
// }

func NewGoogleOAuthConfig(cfg *conf.Config) *oauth2.Config {
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
