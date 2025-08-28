package biz

import (
	"context"
	"encoding/json"
	"errors"
	"io"

	"golang.org/x/oauth2"
)

const GoogleOauthState = "google_oauth_state"
const GooglePKCEKey = "google_pkce_verifier"

const (
	OAuthProviderGoogle = "google"
)

type GoogleUserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func (r GoogleUserInfo) Map() map[string]any {
	data, _ := json.Marshal(r)
	var result map[string]interface{}
	_ = json.Unmarshal(data, &result)
	return result
}

type OAuthUseCase struct {
	google *oauth2.Config
}

func NewOAuthUseCase(google *oauth2.Config) *OAuthUseCase {
	return &OAuthUseCase{google: google}
}

func (r OAuthUseCase) GoogleUserInfo(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*GoogleUserInfo, error) {
	token, err := r.google.Exchange(ctx, code, opts...)
	if err != nil {
		return nil, errors.New("code exchange failed")
	}

	client := r.google.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, errors.New("failed to get user info")
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	var userInfo GoogleUserInfo
	if err = json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, errors.New("failed to parse user info")
	}
	return &userInfo, nil
}
