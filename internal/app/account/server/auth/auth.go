package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-keg/monorepo/internal/data/account/ent"
	"github.com/go-keg/monorepo/internal/data/account/ent/tenantuser"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

type userIDKey struct{}
type userKey struct{}

type tenantIDKey struct{}
type tenantUserIDKey struct{}
type tenantUserKey struct{}

func GetUserID(ctx context.Context) int {
	return cast.ToInt(ctx.Value(userIDKey{}))
}

func GetUser(ctx context.Context) *ent.User {
	v := ctx.Value(userKey{})
	if v == nil {
		return nil
	}
	return v.(*ent.User)
}

func GetTenantUserID(ctx context.Context) int {
	return cast.ToInt(ctx.Value(tenantUserIDKey{}))
}

func GetTenantUser(ctx context.Context) *ent.TenantUser {
	v := ctx.Value(tenantUserKey{})
	if v == nil {
		return nil
	}
	return v.(*ent.TenantUser)
}

func GetTenantID(ctx context.Context) int {
	v := ctx.Value(tenantIDKey{})
	if v == nil {
		return 0
	}
	return v.(int)
}

func Middleware(key string, client *ent.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tenantID := cast.ToInt(r.Header.Get("X-Tenant-ID"))
		ctx := r.Context()
		ctx = context.WithValue(ctx, tenantIDKey{}, tenantID)
		token := r.Header.Get("Authorization")
		auths := strings.SplitN(token, " ", 2)
		if len(auths) == 2 && auths[0] == "Bearer" {
			parse, err := jwt.Parse(auths[1], func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(key), nil
			})
			if err == nil {
				subject, err := parse.Claims.GetSubject()
				if err == nil {
					u, err := client.User.Get(ctx, cast.ToInt(subject))
					if err == nil {
						ctx = context.WithValue(ctx, userIDKey{}, u.ID)
						ctx = context.WithValue(ctx, userKey{}, u)
						tenantUser, err := client.TenantUser.Query().Where(
							tenantuser.TenantID(tenantID),
							tenantuser.UserID(tenantID),
						).First(ctx)
						if err == nil {
							ctx = context.WithValue(ctx, tenantUserKey{}, tenantUser)
							ctx = context.WithValue(ctx, tenantUserIDKey{}, tenantUser.ID)
						}
					}
				}
			}
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
