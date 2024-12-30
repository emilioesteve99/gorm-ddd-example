package userApplicationCommandHandlers

import (
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
)

type UserInsertOneCommandHandler struct {
	commonApplicationCommandHandlers.InsertOneCommandHandler[userDomainCommands.UserInsertOneCommand, userDomainModels.User]
}

func NewUserInsertOneCommandHandler(insertOneUserManager commonDomainManagers.InsertOneManager[userDomainCommands.UserInsertOneCommand, userDomainModels.User]) *UserInsertOneCommandHandler {
	return &UserInsertOneCommandHandler{
		InsertOneCommandHandler: commonApplicationCommandHandlers.NewBaseInsertOneCommandHandler[
			userDomainCommands.UserInsertOneCommand,
			userDomainModels.User,
		](insertOneUserManager),
	}
}
