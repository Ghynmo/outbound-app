package config

import (
	"e-commerce-1/helper"
	"fmt"
)

type JWTConfig struct {
	SecretKey string
}

func NewJWTConfig() JWTConfig {
	return JWTConfig{
		SecretKey: helper.GetEnv("JWT_SECRET_KEY", "your-secret-key"),
	}
}

func (c JWTConfig) Validate() error {
	if c.SecretKey == "" {
		return fmt.Errorf("JWT secret key is required")
	}

	return nil
}