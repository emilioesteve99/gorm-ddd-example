package userGormConverters

import (
	"context"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"gorm.io/gorm"
)

type UserFindOneQueryToUserGormFindQueryConverter struct{}

func NewUserFindOneQueryToUserGormFindQueryConverter() *UserFindOneQueryToUserGormFindQueryConverter {
	return &UserFindOneQueryToUserGormFindQueryConverter{}
}

func (c *UserFindOneQueryToUserGormFindQueryConverter) Convert(input userDomainQueries.UserFindOneQuery, cursor *gorm.DB, _ context.Context) (*gorm.DB, error) {
	if input.Email != nil {
		cursor = cursor.Where("email = ?", *input.Email)
	}
	if input.Ids != nil && len(*input.Ids) > 0 {
		cursor = cursor.Where("id IN ?", *input.Ids)
	}

	return cursor, nil
}
