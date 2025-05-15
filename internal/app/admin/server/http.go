package server

import (
	"crypto/rand"
	"encoding/base64"
	nethttp "net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-keg/apis/api"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/admin/biz"
	"github.com/go-keg/monorepo/internal/app/admin/conf"
	"github.com/go-keg/monorepo/internal/app/admin/server/auth"
	"github.com/go-keg/monorepo/internal/app/admin/service/graphql/dataloader"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/swagger-api"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/oauth2"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, client *ent.Client, schema graphql.ExecutableSchema, oauth *oauth2.Config) *http.Server {
	srv := http.NewServer(cfg.Server.HTTP.HttpOptions(logger)...)
	// graphql
	gqlSrv := gql.NewServer(schema)
	loader := dataloader.NewDataLoader(client)
	srv.Handle("/query", auth.Middleware(cfg.Key, client, dataloader.Middleware(loader, gqlSrv)))
	srv.HandleFunc("/auth/google/authorize", HandleGoogleAuthorize(oauth))
	srv.HandleFunc("/graphql-ui", playground.Handler("Admin", "/query",
		playground.WithStoragePrefix("keg_"),
		playground.WithGraphiqlEnablePluginExplorer(true),
	))
	srv.HandlePrefix("/swagger/", swagger.Handler(api.FS))
	return srv
}

func HandleGoogleAuthorize(config *oauth2.Config) nethttp.HandlerFunc {
	return func(w nethttp.ResponseWriter, r *nethttp.Request) {
		// 生成一个随机的 state 值
		b := make([]byte, 16)
		_, _ = rand.Read(b)
		state := base64.URLEncoding.EncodeToString(b)
		nethttp.SetCookie(w, &nethttp.Cookie{
			Name:     biz.GoogleOauthState,
			Value:    state,
			MaxAge:   300,
			HttpOnly: true,
			Secure:   false, // 生产环境应为 true（HTTPS）
			Path:     "/",
		})
		url := config.AuthCodeURL(state)
		nethttp.Redirect(w, r, url, nethttp.StatusTemporaryRedirect)
	}
}
