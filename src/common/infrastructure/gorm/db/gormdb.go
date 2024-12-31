package gormdb

import (
	"fmt"
	"gorm-ddd-example/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func NewGormDB(cfg config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		cfg.Db.Host,
		cfg.Db.Username,
		cfg.Db.Password,
		cfg.Db.Name,
		cfg.Db.Port,
	)
	db, openDbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})
	if openDbErr != nil {
		panic(openDbErr)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db
}
