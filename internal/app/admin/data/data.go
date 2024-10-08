package data

import (
	"github.com/IBM/sarama"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/data/example/ent"
	ent2 "github.com/go-keg/keg/contrib/ent"
	"github.com/go-keg/keg/contrib/kafka"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewEntDatabase,
	NewKafkaProducer,
)

type Data struct {
	client *ent.Client
	db     *ent.Database
}

func NewData(client *ent.Client, db *ent.Database, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{
		client: client,
		db:     db,
	}
	return d, func() {
		if err := d.client.Close(); err != nil {
			l.Error(err)
		}
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

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
