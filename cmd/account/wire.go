//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-keg/monorepo/internal/app/account/biz"
	"github.com/go-keg/monorepo/internal/app/account/conf"
	"github.com/go-keg/monorepo/internal/app/account/data"
	"github.com/go-keg/monorepo/internal/app/account/job"
	"github.com/go-keg/monorepo/internal/app/account/schedule"
	"github.com/go-keg/monorepo/internal/app/account/server"
	"github.com/go-keg/monorepo/internal/app/account/service"
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
