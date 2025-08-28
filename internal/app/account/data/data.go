package data

import (
	"github.com/go-keg/keg/contrib/ent/driver"
	"github.com/go-keg/monorepo/internal/app/account/conf"
	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
	NewEntDatabase,
	NewUserRepo,
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

