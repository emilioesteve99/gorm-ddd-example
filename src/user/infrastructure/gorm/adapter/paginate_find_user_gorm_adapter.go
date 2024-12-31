package userGormAdapters

import (
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonGormAdapters "gorm-ddd-example/src/common/infrastructure/gorm/adapter"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
	"gorm.io/gorm"
)

type PaginateFindUserGormAdapter struct {
	commonDomainAdapters.PaginateFindAdapter[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]
}

func NewPaginateFindUserGormAdapter(
	db *gorm.DB,
	userFindQueryToGormUserFindQueryConverter domainConverters.ConverterWithExtraArgs[userDomainQueries.UserPaginateFindQuery, *gorm.DB, *gorm.DB],
	usersToPaginatedUsersConverter domainConverters.ConverterWithExtraArgs[[]userGormModels.UserGorm, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[userDomainModels.User]],
) commonDomainAdapters.PaginateFindAdapter[userDomainQueries.UserPaginateFindQuery, userDomainModels.User] {
	return &PaginateFindUserGormAdapter{
		PaginateFindAdapter: commonGormAdapters.NewPaginateFindGormAdapter(
			db,
			userFindQueryToGormUserFindQueryConverter,
			usersToPaginatedUsersConverter,
		),
	}
}
