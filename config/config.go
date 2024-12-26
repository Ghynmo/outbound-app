package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	MySQL    MySQLConfig
	Firebase FirebaseConfig
	JWT      JWTConfig
	// Redis    RedisConfig
	// Upload   UploadConfig
}

// LoadConfig loads all configurations
func LoadConfig() (*Config, error) {
	// Load .env file
	godotenv.Load()

	config := &Config{
		MySQL:    NewMySQLConfig(),
		Firebase: NewFirebaseConfig(),
		JWT:      NewJWTConfig(),
		// Redis:    NewRedisConfig(),
		// Upload:   NewUploadConfig(),
	}

	// Validate all configs
	if err := config.Validate(); err != nil {
		return nil, err
	}

	return config, nil
}

// Validate validates all configurations
func (c *Config) Validate() error {
	validators := []validator{
		c.MySQL,
		c.Firebase,
		c.JWT,
		// c.Redis,
		// c.Upload,
	}

	for _, v := range validators {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// Interface for config validation
type validator interface {
	Validate() error
}
