
package job

import (
    "context"

    "github.com/go-keg/go-job"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJob)

type Job struct {
    *job.Job
}

func NewJob(logger log.Logger) *Job {
    j := job.NewJob(logger, job.NewWorker("example", example))
    return &Job{j}
}

func (j Job) Start(ctx context.Context) error {
    return j.Job.Start(ctx)
}

func (j Job) Stop(ctx context.Context) error {
    return nil
}

func example(ctx context.Context) error {
    panic("todo")
}