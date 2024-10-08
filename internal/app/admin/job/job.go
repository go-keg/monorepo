package job

import (
	"context"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/app/admin/data"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewJob)

type Job struct {
}

func NewJob(logger log.Logger, cfg *conf.Config, data *data.Data) *Job {

	return &Job{}
}

func (j Job) Start(ctx context.Context) error {
	return nil
}

func (j Job) Stop(ctx context.Context) error {
	return nil
}

type exampleJob struct {
}

func (e exampleJob) Run(ctx context.Context) error {
	panic("todo")
}
