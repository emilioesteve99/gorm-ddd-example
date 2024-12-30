package main

import (
	"github.com/golobby/container/v3"
	commonDependencies "gorm-ddd-example/src/common/infrastructure/ioc"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
	"gorm.io/gorm"
)

func main() {
	commonDependencies.InitDependencies()

	var db *gorm.DB
	container.MustResolve(commonDependencies.Container, &db)

	migrateErr := db.AutoMigrate(&userGormModels.UserGorm{})
	if migrateErr != nil {
		panic(migrateErr)
	}
}
