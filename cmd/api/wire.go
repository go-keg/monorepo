//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-keg/monorepo/internal/app/api/conf"
	"github.com/go-keg/monorepo/internal/app/api/data"
	"github.com/go-keg/monorepo/internal/app/api/server"
	"github.com/go-keg/monorepo/internal/app/api/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func initApp(log.Logger, *conf.Config) (*kratos.App, func(), error) {
	panic(wire.Build(data.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}
