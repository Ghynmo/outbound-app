package config

import (
	"os"

	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OauthConfig struct {
    OAuthConfig *oauth2.Config
    Store       *session.Store
}

func NewConfig() *OauthConfig {
    store := session.New()

    oauthConfig := &oauth2.Config{
        ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
        ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
        RedirectURL:  "http://localhost:3000/auth/callback",
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
            "https://www.googleapis.com/auth/userinfo.profile",
        },
        Endpoint: google.Endpoint,
    }

    return &OauthConfig{
        OAuthConfig: oauthConfig,
        Store:       store,
    }
}