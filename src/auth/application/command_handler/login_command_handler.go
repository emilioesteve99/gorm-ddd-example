package authCommandHandlers

import (
	"context"
	authDomainCommands "gorm-ddd-example/src/auth/domain/command"
	authDomainManagers "gorm-ddd-example/src/auth/domain/manager"
	authDomainModels "gorm-ddd-example/src/auth/domain/model"
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
)

type LoginCommandHandler struct {
	loginManager authDomainManagers.ILoginManager
}

func NewLoginCommandHandler(loginManager authDomainManagers.ILoginManager) commonApplicationCommandHandlers.ILoginCommandHandler {
	return &LoginCommandHandler{
		loginManager: loginManager,
	}
}

func (h *LoginCommandHandler) Handle(command authDomainCommands.LoginCommand, ctx context.Context) (authDomainModels.AuthResponse, error) {
	return h.loginManager.Login(command, ctx)
}
