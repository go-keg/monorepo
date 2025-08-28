
package server

import (
    "github.com/99designs/gqlgen/graphql"
    "github.com/99designs/gqlgen/graphql/playground"
    "github.com/go-keg/keg/contrib/gql"
    "github.com/go-keg/monorepo/internal/app/crm/conf"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new HTTP server.
func NewHTTPServer(c *conf.Crm, logger log.Logger, schema graphql.ExecutableSchema) *http.Server {
    server := http.NewServer(c.Server.HTTP.HttpOptions(logger)...)
    server.Handle("/crm/query", gql.NewServer(schema))
    server.HandleFunc("/crm/graphql-ui", playground.Handler("Crm", "/crm/query"))
    return server
}
