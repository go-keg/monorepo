package middleware

import (
	"net/http"

	v1 "github.com/go-keg/apis/api/gateway/middleware/v1"
	config "github.com/go-kratos/gateway/api/gateway/config/v1"
	"github.com/go-kratos/gateway/middleware"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func AuthMiddleware(c *config.Middleware) (middleware.Middleware, error) {
	options := &v1.Auth{
		Domains: map[string]*v1.AuthOption{},
	}
	if c.Options != nil {
		if err := anypb.UnmarshalTo(c.Options, options, proto.UnmarshalOptions{Merge: true}); err != nil {
			return nil, err
		}
	}
	return func(next http.RoundTripper) http.RoundTripper {
		return middleware.RoundTripperFunc(func(req *http.Request) (*http.Response, error) {
			// TODO check auth

			// set auth header
			req.Header.Set("X-ACCOUNT-ID", "1")
			return next.RoundTrip(req)
		})
	}, nil
}
