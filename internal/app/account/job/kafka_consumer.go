package job

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/go-keg/keg/contrib/kafka"
	"github.com/go-keg/monorepo/internal/app/account/conf"
	"github.com/go-kratos/kratos/v2/log"
)

type KafkaConsumer struct {
	cg  *kafka.ConsumerGroup
	log *log.Helper
}

func NewKafkaConsumer(cfg *conf.Config, logger log.Logger) (*KafkaConsumer, error) {
	r := &KafkaConsumer{
		log: log.NewHelper(log.With(logger, "module", "kafka")),
	}
	cg, err := r.consumerGroup(cfg)
	if err != nil {
		r.log.Error(err)
	}
	r.cg = cg
	return r, nil
}

func (r KafkaConsumer) consumerGroup(cfg *conf.Config) (*kafka.ConsumerGroup, error) {
	handler := kafka.NewConsumerGroupHandler(func(message *sarama.ConsumerMessage) error {
		switch message.Topic {
		case "test-topic":
			// TODO
		default:
			r.log.Infow("topic", message.Topic, "partition", message.Partition, "offset", message.Partition)
		}
		return nil
	})
	return kafka.NewConsumerGroupFromConfig(cfg.Data.Kafka, cfg.KafkaConsumerGroup, handler)
}

func (r KafkaConsumer) Run(ctx context.Context) error {
	if r.cg != nil {
		return r.cg.Run(ctx)
	}
	return nil
}
