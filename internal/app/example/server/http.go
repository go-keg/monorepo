package server

import (
	"fmt"
	nethttp "net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-keg/apis/api"
	"github.com/go-keg/keg/contrib/gql"
	"github.com/go-keg/monorepo/internal/app/example/biz"
	"github.com/go-keg/monorepo/internal/app/example/conf"
	"github.com/go-keg/monorepo/internal/app/example/server/auth"
	"github.com/go-keg/monorepo/internal/app/example/service/graphql/dataloader"
	"github.com/go-keg/monorepo/internal/data/example/ent"
	"github.com/go-keg/swagger-api"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"github.com/gorilla/websocket"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"golang.org/x/oauth2"
)

func NewHTTPServer(
	logger log.Logger,
	cfg *conf.Config,
	client *ent.Client,
	schema graphql.ExecutableSchema,
	oauth *oauth2.Config,
	oauthUc *biz.OAuthUseCase,
	ur biz.UserRepo,
	uc *biz.UserUseCase,
) *http.Server {
	srv := http.NewServer(cfg.Server.HTTP.HttpOptions(logger, http.Filter(handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}),
	)))...)
	// graphql
	gqlSrv := gql.NewServer(schema, gql.WithWebsocket(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *nethttp.Request) bool {
				return true
			},
		},
	}))
	loader := dataloader.NewDataLoader(client)
	srv.Handle("/query", auth.Middleware(cfg.Key, client, dataloader.Middleware(loader, gqlSrv)))
	srv.HandlePrefix("/auth/google/", HandleGoogle(oauth, oauthUc, ur, uc))
	srv.HandleFunc("/graphql-ui", playground.Handler(
		"Account",
		"/query",
		playground.WithStoragePrefix("account_"),
		playground.WithGraphiqlEnablePluginExplorer(true),
	))
	srv.HandlePrefix("/swagger/", swagger.Handler(api.FS))
	return srv
}
func HandleGoogle(config *oauth2.Config, oauthUc *biz.OAuthUseCase, ur biz.UserRepo, uc *biz.UserUseCase) nethttp.Handler {
	r := gin.Default()
	store := cookie.NewStore([]byte(os.Getenv("SESSION_SECRET")))
	store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   false,
		SameSite: nethttp.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("sess", store))
	r.GET("/auth/google/login", func(c *gin.Context) {
		sess := sessions.Default(c)
		state := lo.RandomString(16, lo.AlphanumericCharset)
		verifier := lo.RandomString(32, lo.AlphanumericCharset)

		sess.Set(biz.GoogleOauthState, state)
		sess.Set(biz.GooglePKCEKey, verifier)
		_ = sess.Save()

		authURL := config.AuthCodeURL(
			state,
			oauth2.AccessTypeOffline,
			oauth2.ApprovalForce,
			oauth2.S256ChallengeOption(verifier),
		)
		c.Redirect(nethttp.StatusFound, authURL)
	})
	r.GET("/auth/google/callback", func(c *gin.Context) {
		sess := sessions.Default(c)
		// 校验 state
		stateParam := c.Query("state")
		code := c.Query("code")
		if code == "" || stateParam == "" {
			c.String(nethttp.StatusBadRequest, "缺少 code 或 state")
			return
		}
		stateSaved := sess.Get(biz.GoogleOauthState)
		if stateSaved == nil || stateParam != stateSaved.(string) {
			c.String(nethttp.StatusBadRequest, "state 校验失败")
			return
		}
		verifier := cast.ToString(sess.Get(biz.GooglePKCEKey))

		sess.Delete(biz.GoogleOauthState)
		sess.Delete(biz.GooglePKCEKey)
		_ = sess.Save()

		userInfo, err := oauthUc.GoogleUserInfo(c, code, oauth2.SetAuthURLParam("code_verifier", verifier))
		if err != nil {
			return
		}
		user, err := ur.FindUserByOAuth(c, biz.OAuthProviderGoogle, userInfo.ID)
		if err != nil {
			user, err = ur.CreateUser(c, ent.User{
				Email:    userInfo.Email,
				Nickname: userInfo.Name,
				Avatar:   userInfo.Picture,
			})
			if err != nil {
				return
			}
			err = ur.BindOAuthAccount(c, ent.OAuthAccount{
				UserID:         user.ID,
				Provider:       "google",
				ProviderUserID: userInfo.ID,
				Profile:        userInfo.Map(),
			})
			if err != nil {
				return
			}
		}
		token, ttl, err := uc.GenerateToken(user.ID)
		if err != nil {
			return
		}
		redirectURL := fmt.Sprintf("/login?token=%s&ttl=%d", token, ttl)
		c.Redirect(nethttp.StatusFound, redirectURL)
	})

	return r
}
