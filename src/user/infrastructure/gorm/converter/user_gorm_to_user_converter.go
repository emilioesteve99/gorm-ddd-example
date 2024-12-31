package userGormConverters

import (
	"context"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
)

type UserGormToUserConverter struct{}

func NewUserGormToUserConverter() domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User] {
	return &UserGormToUserConverter{}
}

func (c *UserGormToUserConverter) Convert(input userGormModels.UserGorm, _ context.Context) (userDomainModels.User, error) {
	return userDomainModels.User{
		Id:       input.ID,
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}, nil
}
