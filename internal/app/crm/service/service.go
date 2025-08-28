
package service

import (
    "github.com/go-keg/monorepo/internal/app/crm/service/graphql"
    "github.com/google/wire"
)

var ProviderSet = wire.NewSet(graphql.NewSchema)
