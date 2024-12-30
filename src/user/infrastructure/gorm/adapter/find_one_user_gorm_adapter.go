package userGormAdapters

import (
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonGormAdapters "gorm-ddd-example/src/common/infrastructure/gorm/adapter"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type FindOneUserGormAdapter struct {
	commonDomainAdapters.FindOneAdapter[userDomainQueries.UserFindOneQuery, userDomainModels.User]
}

func NewFindOneUserGormAdapter(
	db *gorm.DB,
	userFindOneQueryToUserGormConverter domainConverters.ConverterWithExtraArgs[userDomainQueries.UserFindOneQuery, *gorm.DB, *gorm.DB],
	userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User],
) *FindOneUserGormAdapter {
	return &FindOneUserGormAdapter{
		FindOneAdapter: commonGormAdapters.NewFindOneGormAdapter[
			userDomainQueries.UserFindOneQuery,
			userGormModels.UserGorm,
			userDomainModels.User,
		](
			db,
			userFindOneQueryToUserGormConverter,
			userGormToUserConverter,
		),
	}
}
