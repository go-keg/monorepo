package job

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/keg/contrib/kafka"
	"github.com/go-kratos/kratos/v2/log"
)

type kafkaConsumer struct {
	cg  *kafka.ConsumerGroup
	log *log.Helper
}

func newKafkaConsumer(cfg *conf.Config, logger log.Logger) *kafkaConsumer {
	return &kafkaConsumer{
		cg:  kafka.NewConsumerGroupFromConfig(cfg.Data.Kafka, cfg.KafkaConsumerGroup),
		log: log.NewHelper(log.With(logger, "module", "kafka")),
	}
}

func (r kafkaConsumer) Run(ctx context.Context) error {
	return r.cg.Run(ctx, func(message *sarama.ConsumerMessage) error {
		switch message.Topic {
		case "test-topic":
			//TODO
		default:
			r.log.Infow("topic", message.Topic, "partition", message.Partition, "offset", message.Partition)
		}
		return nil
	})
}
