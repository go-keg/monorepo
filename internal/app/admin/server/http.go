package server

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-keg/apis/api"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/admin/conf"
	"github.com/go-keg/monorepo/internal/app/admin/server/auth"
	"github.com/go-keg/monorepo/internal/app/admin/service/graphql/dataloader"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/swagger-api"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	// graphql
	gqlSrv := gql.NewServer(schema)
	loader := dataloader.NewDataLoader(client)
	srv.Handle("/query", auth.Middleware(cfg.Key, client, dataloader.Middleware(loader, gqlSrv)))
	srv.HandleFunc("/graphql-ui", playground.Handler("Admin", "/query",
		playground.WithStoragePrefix("keg_"),
		playground.WithGraphiqlEnablePluginExplorer(true),
	))
	srv.HandlePrefix("/swagger/", swagger.Handler(api.FS))
	return srv
}
