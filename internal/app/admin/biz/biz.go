package biz

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewUserUseCase,
)

const GoogleOauthState = "google_oauth_state"
