package authioc

import (
	"github.com/golobby/container/v3"
	authCommandHandlers "gorm-ddd-example/src/auth/application/command_handler"
	authDomainConverters "gorm-ddd-example/src/auth/domain/converter"
	authDomainManagers "gorm-ddd-example/src/auth/domain/manager"
	authControllers "gorm-ddd-example/src/auth/infrastructure/http/controller"
	commonInfraUtils "gorm-ddd-example/src/common/infrastructure/utils"
)

func InitAuthDependencies(c container.Container) {
	converters := []any{
		authDomainConverters.NewLoginCommandToJwtAuthResponseConverter,
	}
	var adapters []any
	managers := []any{
		authDomainManagers.NewLoginManager,
	}
	applicationHandlers := []any{
		authCommandHandlers.NewLoginCommandHandler,
	}
	controllers := []any{
		authControllers.NewLoginHttpController,
	}
	var allFactories []any
	allFactories = append(allFactories, converters...)
	allFactories = append(allFactories, adapters...)
	allFactories = append(allFactories, managers...)
	allFactories = append(allFactories, applicationHandlers...)
	allFactories = append(allFactories, controllers...)
	commonInfraUtils.RegisterSingletonFactories(allFactories, c)
}
