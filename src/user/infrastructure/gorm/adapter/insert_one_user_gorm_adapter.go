package userGormAdapters

import (
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonGormAdapters "gorm-ddd-example/src/common/infrastructure/gorm/adapter"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type InsertOneUserGormAdapter struct {
	commonDomainAdapters.InsertOneAdapter[userDomainCommands.UserInsertOneCommand, userDomainModels.User]
}

func NewInsertOneUserGormAdapter(
	db *gorm.DB,
	userInsertOneCommandToUserGormConverter domainConverters.Converter[userDomainCommands.UserInsertOneCommand, userGormModels.UserGorm],
	userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User],
) *InsertOneUserGormAdapter {
	return &InsertOneUserGormAdapter{
		InsertOneAdapter: commonGormAdapters.NewInsertOneGormAdapter[
			userDomainCommands.UserInsertOneCommand,
			userGormModels.UserGorm,
			userDomainModels.User,
		](
			db,
			userInsertOneCommandToUserGormConverter,
			userGormToUserConverter,
		),
	}
}
