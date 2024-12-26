package config

import (
	"e-commerce-1/helper"
	"fmt"
	"time"
)

type MySQLConfig struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
	SSLMode      string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifetime  time.Duration
}

func NewMySQLConfig() MySQLConfig {
	return MySQLConfig{
		Host:         helper.GetEnv("MYSQL_HOST", "localhost"),
		Port:         helper.GetEnv("MYSQL_PORT", "3306"),
		User:         helper.GetEnv("MYSQL_USER", "root"),
		Password:     helper.GetEnv("MYSQL_PASSWORD", ""),
		DatabaseName: helper.GetEnv("MYSQL_DATABASE", "your_db"),
		SSLMode:      helper.GetEnv("MYSQL_SSL_MODE", "disable"),
		MaxIdleConns: helper.GetEnvAsInt("MYSQL_MAX_IDLE_CONNS", 10),
		MaxOpenConns: helper.GetEnvAsInt("MYSQL_MAX_OPEN_CONNS", 100),
		MaxLifetime:  time.Duration(helper.GetEnvAsInt("MYSQL_CONN_MAX_LIFETIME", 3600)) * time.Second,
	}
}

func (c MySQLConfig) GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.User,
		c.Password,
		c.Host,
		c.Port,
		c.DatabaseName,
	)
}

func (c MySQLConfig) Validate() error {
	if c.DatabaseName == "" {
		return fmt.Errorf("mysql database name is required")
	}
	return nil
}