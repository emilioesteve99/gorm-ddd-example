package userDomainManagers

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	appErrors "gorm-ddd-example/src/common/application/model"
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
)

type InsertOneUserManager struct {
	commonDomainManagers.InsertOneManager[userDomainCommands.UserInsertOneCommand, userDomainModels.User]
}

func NewInsertOneUserManager(insertOneAdapter commonDomainAdapters.InsertOneAdapter[userDomainCommands.UserInsertOneCommand, userDomainModels.User]) *InsertOneUserManager {
	return &InsertOneUserManager{
		InsertOneManager: *commonDomainManagers.NewBaseInsertOneManager[
			userDomainCommands.UserInsertOneCommand,
			userDomainModels.User,
		](insertOneAdapter),
	}
}

func (m *InsertOneUserManager) Manage(command userDomainCommands.UserInsertOneCommand, ctx context.Context) (userDomainModels.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		return userDomainModels.User{}, appErrors.BuildUnknownError(err)
	}
	command.Password = string(hashedPassword)
	return m.InsertOneManager.Manage(command, ctx)
}
