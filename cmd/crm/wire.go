//go:build wireinject
// +build wireinject
package main

import (
    "github.com/go-keg/monorepo/internal/app/crm/conf"
    "github.com/go-keg/monorepo/internal/app/crm/data"
    "github.com/go-keg/monorepo/internal/app/crm/job"
    "github.com/go-keg/monorepo/internal/app/crm/schedule"
    "github.com/go-keg/monorepo/internal/app/crm/server"
    "github.com/go-keg/monorepo/internal/app/crm/service"
    

    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/google/wire"
)

func initApp(*conf.Crm, log.Logger) (*kratos.App, func(), error) {
    panic(wire.Build(data.ProviderSet, job.ProviderSet, schedule.ProviderSet, server.ProviderSet, service.ProviderSet, newApp))
}