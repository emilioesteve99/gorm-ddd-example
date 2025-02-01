package authDomainManagers

import (
	"context"
	authDomainCommands "gorm-ddd-example/src/auth/domain/command"
	authDomainModels "gorm-ddd-example/src/auth/domain/model"
	appErrors "gorm-ddd-example/src/common/application/model"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
)

type ILoginManager interface {
	Login(command authDomainCommands.LoginCommand, ctx context.Context) (authDomainModels.AuthResponse, error)
}

type LoginManager struct {
	findOneUserManager                     commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User]
	loginCommandToJwtAuthResponseConverter domainConverters.ConverterWithExtraArgs[authDomainCommands.LoginCommand, userDomainModels.User, authDomainModels.AuthResponse]
}

func NewLoginManager(
	findOneUserManager commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User],
	loginCommandToAuthResponseConverter domainConverters.ConverterWithExtraArgs[authDomainCommands.LoginCommand, userDomainModels.User, authDomainModels.AuthResponse],
) ILoginManager {
	return &LoginManager{
		findOneUserManager:                     findOneUserManager,
		loginCommandToJwtAuthResponseConverter: loginCommandToAuthResponseConverter,
	}
}

func (m *LoginManager) Login(command authDomainCommands.LoginCommand, ctx context.Context) (authDomainModels.AuthResponse, error) {
	user, findOneUserErr := m.findOneUserManager.Manage(
		userDomainQueries.UserFindOneQuery{
			Email: &command.Email,
		},
		ctx,
	)
	if findOneUserErr != nil {
		return authDomainModels.AuthResponse{}, findOneUserErr
	}
	if user == nil {
		return authDomainModels.AuthResponse{}, appErrors.AppError{
			Code:    appErrors.UnauthorizedCode,
			Message: appErrors.UnauthorizedMsg,
		}
	}
	return m.loginCommandToJwtAuthResponseConverter.Convert(command, *user, ctx)
}
