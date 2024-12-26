package mysql

import (
	"e-commerce-1/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(cfg *config.MySQLConfig) (*gorm.DB, error) {
    db, err := gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("failed to connect to database: %v", err)
    }

    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }

    // Set connection pool settings
    sqlDB.SetMaxIdleConns(10)
    sqlDB.SetMaxOpenConns(100)
    
    return db, nil
}