package data

import (
	"github.com/go-keg/keg/contrib/ent/driver"
	"github.com/go-keg/monorepo/internal/app/crm/conf"
	"github.com/go-keg/monorepo/internal/data/crm/ent"
	_ "github.com/go-keg/monorepo/internal/data/crm/ent/runtime"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
)

func NewEntClient(cfg *conf.Crm) (*ent.Client, error) {
	drv, err := driver.NewDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}
