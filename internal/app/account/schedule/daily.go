package schedule

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/go-keg/monorepo/internal/data/account/ent/user"
)

type Daily struct {
	db *ent.Client
	// producer sarama.SyncProducer
}

func NewDaily(
	db *ent.Client,
	// producer sarama.SyncProducer,
) *Daily {
	return &Daily{
		db: db,
		// producer: producer,
	}
}

// Run mock daily send statistic data to accounts
func (r Daily) Run(ctx context.Context) error {
	startID := 0
	for {
		accounts, err := r.db.User.Query().Where(user.IDGT(startID)).Limit(200).All(ctx)
		if err != nil {
			return err
		}
		if len(accounts) == 0 {
			break
		}
		var messages []*sarama.ProducerMessage
		for i := range accounts {
			messages = append(messages, &sarama.ProducerMessage{
				Topic: "send_daily_statistic",
				Key:   sarama.StringEncoder(fmt.Sprintf("account:%d", accounts[i].ID)),
				Value: sarama.StringEncoder("statistic data..."),
			})
		}
		// err = r.producer.SendMessages(messages)
		// if err != nil {
		// 	return err
		// }
		fmt.Println(messages)
		startID = accounts[len(accounts)-1].ID
	}
	return nil
}
