package commonDependencies

import (
	"github.com/golobby/container/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

func InitCommonGormDependencies(Container container.Container) {
	container.MustSingleton(Container, func() *gorm.DB {
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
	})
}
