package job

import (
	"context"
	"os"
	"time"

	"github.com/go-keg/go-job"
	"github.com/go-keg/go-job/report/qyweixin"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewJob,
	NewKafkaConsumer,
)

type Job struct {
	job *job.Job
	ent *ent.Client
	log *log.Helper
}

func NewJob(logger log.Logger, ent *ent.Client, cg *KafkaConsumer) *Job {
	j := job.NewJob(logger,
		job.NewWorker("customer", cg.Run),
		job.NewWorker("example", exampleJob,
			job.WithLimiterDuration(time.Minute),
			job.WithReport(qyweixin.NewReport(os.Getenv("QY_WECHAT_TOKEN"))),
		),
	)
	return &Job{
		job: j,
		ent: ent,
		log: log.NewHelper(log.With(logger, "module", "job")),
	}
}

func (j Job) Start(ctx context.Context) error {
	return j.job.Start(ctx)
}

func (j Job) Stop(_ context.Context) error {
	j.job.Stop()
	return nil
}

func exampleJob(_ context.Context) error {
	panic("todo")
}
