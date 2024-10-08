package data

import (
	"github.com/go-keg/example/internal/app/api/conf"
	"github.com/go-keg/example/internal/data/example/ent"
	ent2 "github.com/go-keg/keg/contrib/ent"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
)

func NewEntClient(cfg *conf.Config) (*ent.Client, error) {
	drv, err := ent2.NewDriver(cfg.Data.Database)
	//drv, err := database.NewEntDriverWithOtel(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}
