package service

import (
	"github.com/go-keg/monorepo/internal/app/account/service/graphql"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	graphql.NewSchema,
	graphql.NewResolver,
	graphql.NewDirectiveRoot,
)
