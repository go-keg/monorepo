package job

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/go-keg/keg/contrib/kafka"
	"github.com/go-kratos/kratos/v2/log"
)

type kafkaConsumer struct {
	cg  *kafka.ConsumerGroup
	log *log.Helper
}

func (r kafkaConsumer) Run(ctx context.Context) error {
	return r.cg.Run(ctx, func(message *sarama.ConsumerMessage) error {
		r.log.Infow("topic", message.Topic, "partition", message.Partition, "offset", message.Partition)
		return nil
	})
}
