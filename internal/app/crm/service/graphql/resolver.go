package graphql

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-keg/monorepo/internal/data/crm/ent"
	"github.com/go-keg/monorepo/internal/pkg/auth"
	"github.com/go-kratos/kratos/v2/log"
)

type Resolver struct {
	client *ent.Client
	log    *log.Helper
}

// NewSchema creates a graphql executable schema.
func NewSchema(logger log.Logger, client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{
			log:    log.NewHelper(log.With(logger, "module", "service/graphql")),
			client: client,
		},
	})
}

func (r *Resolver) TenantUserID(ctx context.Context) int {
	return auth.GetAccountID(ctx)
}
