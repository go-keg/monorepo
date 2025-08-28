
package schedule

import (
    "context"
    "fmt"

    "github.com/go-keg/keg/contrib/schedule"
    "github.com/go-keg/monorepo/internal/data/crm/ent"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewSchedule)

type Schedule struct {
    *schedule.Schedule
    ent *ent.Client
}

func NewSchedule(logger log.Logger, ent *ent.Client) *Schedule {
    s := schedule.NewSchedule(logger)
    return &Schedule{Schedule: s, ent: ent}
}

func (s Schedule) Start(ctx context.Context) error {
    _, _ = s.Add("example", "* * * * *", func() error {
        return s.example(ctx)
    })
    return s.Schedule.Start()
}

func (s Schedule) Stop(ctx context.Context) error {
    return s.Schedule.Stop()
}

func (s Schedule) example(ctx context.Context) error {
    fmt.Println("todo")
    return nil
}