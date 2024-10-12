package auth

import (
	"context"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type ctxKey string

const userID ctxKey = "userID"
const userKey ctxKey = "user"

func GetUserID(ctx context.Context) int {
	return cast.ToInt(ctx.Value(userID))
}

func GetUser(ctx context.Context) *ent.User {
	v := ctx.Value(userKey)
	if v == nil {
		return nil
	}
	return v.(*ent.User)
}

func Middleware(key string, client *ent.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		auths := strings.SplitN(token, " ", 2)
		if len(auths) == 2 && auths[0] == "Bearer" {
			claims := jwt.RegisteredClaims{}
			_, err := jwt.ParseWithClaims(auths[1], &claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(key), nil
			})
			if err == nil {
				u, err := client.User.Get(r.Context(), cast.ToInt(claims.Subject))
				if err != nil {
					return
				}
				ctx := context.WithValue(r.Context(), userID, u.ID)
				ctx = context.WithValue(ctx, userKey, u)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}
