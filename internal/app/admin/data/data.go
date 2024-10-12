package data

import (
	"github.com/IBM/sarama"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/data/example/ent"
	ent2 "github.com/go-keg/keg/contrib/ent"
	"github.com/go-keg/keg/contrib/kafka"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEntClient,
	NewEntDatabase,
	NewKafkaProducer,
)

func NewEntClient(cfg *conf.Config) (*ent.Client, error) {
	drv, err := ent2.NewDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewClient(ent.Driver(drv)), nil
}

func NewEntDatabase(cfg *conf.Config) (*ent.Database, error) {
	drv, err := ent2.NewDriver(cfg.Data.Database)
	if err != nil {
		return nil, err
	}
	return ent.NewDatabase(ent.Driver(drv)), nil
}

func NewKafkaProducer(cfg *conf.Config) (sarama.SyncProducer, error) {
	return kafka.NewSyncProducerFromConfig(cfg.Data.Kafka)
}
