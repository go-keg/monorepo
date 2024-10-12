package middleware

import (
	"context"
	"fmt"
	accountv1 "github.com/go-keg/apis/api/account/v1"
	"github.com/go-keg/apis/api/common"
	v1 "github.com/go-keg/apis/api/gateway/middleware/v1"
	"github.com/go-keg/example/internal/pkg/auth"
	"github.com/go-keg/keg/contrib/cache"
	config "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/middleware"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/spf13/cast"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"net/http"
	"os"
)

type PermissionReader func(req *http.Request) []string

func NewAuthService() (accountv1.AuthServiceClient, error) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(os.Getenv("AUTH_SERVICE_GRPC_ENDPOINT")),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		return nil, err
	}
	return accountv1.NewAuthServiceClient(conn), nil
}
func accountPermissionReader() (PermissionReader, error) {
	as, err := NewAuthService()
	if err != nil {
		return nil, err
	}
	return func(req *http.Request) []string {
		accountID := cast.ToUint32(req.Header.Get(auth.HeaderAccountID))
		cacheKey := fmt.Sprintf("account:%d:keywords", accountID)
		keywords, err := cache.LocalRemember(cacheKey, cache.DefaultExpiration, func() ([]string, error) {
			reply, err := as.Permissions(req.Context(), &common.ID{Id: accountID})
			if err != nil {
				return nil, err
			}
			return reply.Keywords, nil
		})
		if err != nil {
			log.Errorw("caller", "accountPermissionReader", "err", err)
			return nil
		}
		return keywords
	}, nil
}

func PermissionMiddleware(c *config.Middleware) (middleware.Middleware, error) {
	options := &v1.Permission{
		Domain:    "",
		Whitelist: nil,
		Blacklist: nil,
	}
	if c.Options != nil {
		if err := anypb.UnmarshalTo(c.Options, options, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, err
		}
	}
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			// TODO check permission

			return next.RoundTrip(req)
		})
	}, nil
}
