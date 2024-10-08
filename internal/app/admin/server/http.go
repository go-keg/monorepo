package server

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-keg/example/internal/app/admin/conf"
	"github.com/go-keg/example/internal/app/admin/server/auth"
	"github.com/go-keg/example/internal/app/admin/service/graphql/dataloader"
	"github.com/go-keg/example/internal/data/example/ent"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/keg/contrib/metric"
	"github.com/go-keg/keg/contrib/response"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema) *http.Server {
	opts := []http.ServerOption{
		http.Address(cfg.Server.Http.Addr),
		http.Network(cfg.Server.Http.Network),
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
			metrics.Server(
				metrics.WithRequests(metric.RequestsCounter),
				metrics.WithSeconds(metric.SecondsHistogram),
			),
		),
	}

	srv := http.NewServer(opts...)
	// graphql
	gqlSrv := handler.NewDefaultServer(schema)

	gqlSrv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		resp := next(ctx)
		for _, e := range resp.Errors {
			if e.Rule != gql.CustomErrorKey {
				e.Message = fmt.Sprintf("error code: %s", response.Err2HashCode(e))
			}
		}
		return resp
	})
	loader := dataloader.NewDataLoader(client)
	srv.Handle("/query", auth.Middleware(cfg.Key, client, dataloader.Middleware(loader, gqlSrv)))
	srv.HandleFunc("/graphql-ui", playground.Handler("Admin", "/query"))
	return srv
}
