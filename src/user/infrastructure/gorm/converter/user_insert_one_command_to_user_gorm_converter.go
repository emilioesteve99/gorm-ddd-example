package userGormConverters

import (
	"context"
	"github.com/google/uuid"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
)

type UserInsertOneCommandToUserGormConverter struct{}

func NewUserInsertOneCommandToUserGormConverter() domainConverters.Converter[userDomainCommands.UserInsertOneCommand, userGormModels.UserGorm] {
	return &UserInsertOneCommandToUserGormConverter{}
}

func (c *UserInsertOneCommandToUserGormConverter) Convert(input userDomainCommands.UserInsertOneCommand, _ context.Context) (userGormModels.UserGorm, error) {
	return userGormModels.UserGorm{
		ID:       uuid.New().String(),
		Email:    input.Email,
		Name:     input.Name,
		Password: input.Password,
	}, nil
}
