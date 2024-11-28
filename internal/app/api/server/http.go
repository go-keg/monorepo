package server

import (
	v1 "github.com/go-keg/apis/api/account/v1"
	"github.com/go-keg/monorepo/internal/app/api/conf"
	"github.com/go-keg/monorepo/internal/app/api/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func NewHTTPServer(cfg *conf.Config, logger log.Logger, account *service.AccountService) *http.Server {
	srv := http.NewServer(cfg.Server.Http.HttpOptions(logger)...)
	v1.RegisterAccountServiceHTTPServer(srv, account)
	return srv
}
