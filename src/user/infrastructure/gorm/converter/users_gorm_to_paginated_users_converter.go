package userGormConverters

import (
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonGormConverters "gorm-ddd-example/src/common/infrastructure/gorm/converter"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
)

type UsersGormToPaginatedUsersConverter struct {
	domainConverters.ConverterWithExtraArgs[[]userGormModels.UserGorm, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[userDomainModels.User]]
}

func NewUsersGormToPaginatedUsersConverter(
	userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User],
) *UsersGormToPaginatedUsersConverter {
	return &UsersGormToPaginatedUsersConverter{
		ConverterWithExtraArgs: commonGormConverters.NewBaseModelsDBToPaginatedModelsConverter(userGormToUserConverter),
	}
}
