package middleware

import "github.com/go-kratos/gateway/middleware"

func init() {
	middleware.Register("keg-example-auth", AuthMiddleware)
	middleware.Register("keg-example-permission", PermissionMiddleware)
}
