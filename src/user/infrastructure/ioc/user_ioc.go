package userioc

import (
	"github.com/golobby/container/v3"
	commonInfraUtils "gorm-ddd-example/src/common/infrastructure/utils"
	userApplicationCommandHandlers "gorm-ddd-example/src/user/application/command_handler"
	userApplicationQueryHandlers "gorm-ddd-example/src/user/application/query_handler"
	userDomainManagers "gorm-ddd-example/src/user/domain/manager"
	userGormAdapters "gorm-ddd-example/src/user/infrastructure/gorm/adapter"
	userGormConverters "gorm-ddd-example/src/user/infrastructure/gorm/converter"
	userControllers "gorm-ddd-example/src/user/infrastructure/http/controller"
)

func InitUserDependencies(c container.Container) {
	converters := []any{
		userGormConverters.NewUserInsertOneCommandToUserGormConverter,
		userGormConverters.NewUserGormToUserConverter,
		userGormConverters.NewUserFindOneQueryToUserGormFindQueryConverter,
		userGormConverters.NewUserFindQueryToUserGormFindQueryConverter,
		userGormConverters.NewUserPaginateFindQueryToUserGormFindQueryConverter,
		userGormConverters.NewUsersGormToPaginatedUsersConverter,
	}
	adapters := []any{
		userGormAdapters.NewInsertOneUserGormAdapter,
		userGormAdapters.NewFindOneUserGormAdapter,
		userGormAdapters.NewPaginateFindUserGormAdapter,
	}
	managers := []any{
		userDomainManagers.NewInsertOneUserManager,
		userDomainManagers.NewFindOneUserManager,
		userDomainManagers.NewPaginateFindUserManager,
	}
	applicationHandlers := []any{
		userApplicationCommandHandlers.NewUserInsertOneCommandHandler,
		userApplicationQueryHandlers.NewUserFindOneQueryHandler,
		userApplicationQueryHandlers.NewUserPaginateFindQueryHandler,
	}
	controllers := []any{
		userControllers.NewInsertOneUserHttpController,
		userControllers.NewFindOneUserHttpController,
		userControllers.NewPaginateFindUserHttpController,
	}
	var allFactories []any
	allFactories = append(allFactories, converters...)
	allFactories = append(allFactories, adapters...)
	allFactories = append(allFactories, managers...)
	allFactories = append(allFactories, applicationHandlers...)
	allFactories = append(allFactories, controllers...)
	commonInfraUtils.RegisterSingletonFactories(allFactories, c)
}
