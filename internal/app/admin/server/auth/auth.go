package auth

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cast"
)

type userIDKey struct{}
type userKey struct{}

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

func Middleware(key string, client *ent.Client, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		auths := strings.SplitN(token, " ", 2)
		if len(auths) == 2 && auths[0] == "Bearer" {
			parse, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method")
				}
				return []byte(key), nil
			})
			if err != nil {
				return
			}
			subject, err := parse.Claims.GetSubject()
			if err != nil {
				return
			}
			if err == nil {
				user, err := client.User.Get(r.Context(), cast.ToInt(subject))
				if err != nil {
					return
				}
				ctx := context.WithValue(r.Context(), userIDKey{}, user.ID)
				ctx = context.WithValue(ctx, userKey{}, user)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}
