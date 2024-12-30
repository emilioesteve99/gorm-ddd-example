package userGormConverters

import (
	"context"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"gorm.io/gorm"
)

type UserFindQueryToUserGormFindQueryConverter struct{}

func NewUserFindQueryToUserGormFindQueryConverter() *UserFindQueryToUserGormFindQueryConverter {
	return &UserFindQueryToUserGormFindQueryConverter{}
}

func (c *UserFindQueryToUserGormFindQueryConverter) Convert(input userDomainQueries.UserPaginateFindQuery, cursor *gorm.DB, _ context.Context) (*gorm.DB, error) {
	if input.Query.Email != nil {
		cursor = cursor.Where("email = ?", *input.Query.Email)
	}
	if input.Query.Ids != nil && len(*input.Query.Ids) > 0 {
		cursor = cursor.Where("id IN ?", *input.Query.Ids)
	}

	return cursor, nil
}
