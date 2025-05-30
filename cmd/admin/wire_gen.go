// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/app/admin/conf"
	"github.com/go-keg/monorepo/internal/app/admin/data"
	"github.com/go-keg/monorepo/internal/app/admin/job"
	"github.com/go-keg/monorepo/internal/app/admin/schedule"
	"github.com/go-keg/monorepo/internal/app/admin/server"
	"github.com/go-keg/monorepo/internal/app/admin/service/graphql"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "github.com/go-keg/monorepo/internal/data/example/ent/runtime"
)

// Injectors from wire.go:

func initApp(logger log.Logger, config *conf.Config) (*kratos.App, func(), error) {
	client, err := data.NewEntClient(config)
	if err != nil {
		return nil, nil, err
	}
	oauth2Config := data.NewGoogleOAuthConfig(config)
	database, err := data.NewEntDatabase(config)
	if err != nil {
		return nil, nil, err
	}
	userUseCase := biz.NewUserUseCase(config)
	userRepo := data.NewUserRepo(client)
	executableSchema := graphql.NewSchema(oauth2Config, logger, database, client, userUseCase, userRepo)
	httpServer := server.NewHTTPServer(config, logger, client, executableSchema, oauth2Config)
	jobJob := job.NewJob(logger, client, config)
	daily := schedule.NewDaily(client)
	scheduleSchedule := schedule.NewSchedule(logger, client, daily)
	app := newApp(logger, httpServer, jobJob, scheduleSchedule)
	return app, func() {
	}, nil
}
