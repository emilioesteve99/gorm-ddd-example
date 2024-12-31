package userGormConverters

import (
	"context"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"gorm.io/gorm"
)

type UserPaginateFindQueryToUserGormFindQueryConverter struct {
	userFindQueryToUserGormFindQueryConverter domainConverters.ConverterWithExtraArgs[userDomainQueries.UserFindQuery, *gorm.DB, *gorm.DB]
}

func NewUserPaginateFindQueryToUserGormFindQueryConverter(
	userFindQueryToUserGormFindQueryConverter domainConverters.ConverterWithExtraArgs[
		userDomainQueries.UserFindQuery,
		*gorm.DB,
		*gorm.DB,
	],
) domainConverters.ConverterWithExtraArgs[userDomainQueries.UserPaginateFindQuery, *gorm.DB, *gorm.DB] {
	return &UserPaginateFindQueryToUserGormFindQueryConverter{userFindQueryToUserGormFindQueryConverter: userFindQueryToUserGormFindQueryConverter}
}

func (c *UserPaginateFindQueryToUserGormFindQueryConverter) Convert(input userDomainQueries.UserPaginateFindQuery, cursor *gorm.DB, ctx context.Context) (*gorm.DB, error) {
	return c.userFindQueryToUserGormFindQueryConverter.Convert(input.Query, cursor, ctx)
}
