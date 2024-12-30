package userGormConverters

import (
	"context"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
)

type UserGormToUserConverter struct{}

func NewUserGormToUserConverter() *UserGormToUserConverter {
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
