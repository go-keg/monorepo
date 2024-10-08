package service

import (
	"github.com/go-keg/example/internal/app/admin/service/graphql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(graphql.NewSchema)
