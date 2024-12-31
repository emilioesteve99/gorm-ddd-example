package gormdb

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

//type Config struct {
//	Host     string
//	User     string
//	Password string
//	DBName   string
//	Port     string
//}

func NewGormDB() *gorm.DB {
	dsn := "host=db user=user password=password dbname=database port=5432 sslmode=disable"
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
