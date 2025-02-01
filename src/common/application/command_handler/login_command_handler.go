package commonApplicationCommandHandlers

import (
	"context"
	authDomainCommands "gorm-ddd-example/src/auth/domain/command"
	authDomainModels "gorm-ddd-example/src/auth/domain/model"
)

type ILoginCommandHandler interface {
	Handle(input authDomainCommands.LoginCommand, ctx context.Context) (authDomainModels.AuthResponse, error)
}
