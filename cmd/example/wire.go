//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-keg/monorepo/internal/app/example/biz"
	"github.com/go-keg/monorepo/internal/app/example/conf"
	"github.com/go-keg/monorepo/internal/app/example/data"
	"github.com/go-keg/monorepo/internal/app/example/job"
	"github.com/go-keg/monorepo/internal/app/example/schedule"
	"github.com/go-keg/monorepo/internal/app/example/server"
	"github.com/go-keg/monorepo/internal/app/example/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(log.Logger, *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(
		biz.ProviderSet,
		conf.ProviderSet,
		data.ProviderSet,
		job.ProviderSet,
		schedule.ProviderSet,
		server.ProviderSet,
		service.ProviderSet,
		newApp,
	))
}
