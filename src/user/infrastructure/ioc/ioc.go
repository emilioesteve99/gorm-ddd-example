package userioc

import (
	"github.com/golobby/container/v3"
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonHttpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	userApplicationCommandHandlers "gorm-ddd-example/src/user/application/command_handler"
	userApplicationQueryHandlers "gorm-ddd-example/src/user/application/query_handler"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainManagers "gorm-ddd-example/src/user/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	userGormAdapters "gorm-ddd-example/src/user/infrastructure/gorm/adapter"
	userGormConverters "gorm-ddd-example/src/user/infrastructure/gorm/converter"
	userGormModels "gorm-ddd-example/src/user/infrastructure/gorm/model"
	userControllers "gorm-ddd-example/src/user/infrastructure/http/controller"
	"gorm.io/gorm"
)

func initGormAdapters(Container container.Container) {
	container.MustSingleton(Container, func() domainConverters.Converter[userDomainCommands.UserInsertOneCommand, userGormModels.UserGorm] {
		return userGormConverters.NewUserInsertOneCommandToUserGormConverter()
	})
	container.MustSingleton(Container, func() domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User] {
		return userGormConverters.NewUserGormToUserConverter()
	})
	container.MustSingleton(Container, func(
		db *gorm.DB,
		userInsertOneCommandToUserGormConverter domainConverters.Converter[userDomainCommands.UserInsertOneCommand, userGormModels.UserGorm],
		userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User],
	) commonDomainAdapters.InsertOneAdapter[userDomainCommands.UserInsertOneCommand, userDomainModels.User] {
		return userGormAdapters.NewInsertOneUserGormAdapter(
			db,
			userInsertOneCommandToUserGormConverter,
			userGormToUserConverter,
		)
	})
	container.MustSingleton(Container, func() domainConverters.ConverterWithExtraArgs[userDomainQueries.UserFindOneQuery, *gorm.DB, *gorm.DB] {
		return userGormConverters.NewUserFindOneQueryToUserGormFindQueryConverter()
	})
	container.MustSingleton(Container, func(
		db *gorm.DB,
		userFindOneQueryToUserGormConverter domainConverters.ConverterWithExtraArgs[userDomainQueries.UserFindOneQuery, *gorm.DB, *gorm.DB],
		userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User],
	) commonDomainAdapters.FindOneAdapter[userDomainQueries.UserFindOneQuery, userDomainModels.User] {
		return userGormAdapters.NewFindOneUserGormAdapter(
			db,
			userFindOneQueryToUserGormConverter,
			userGormToUserConverter,
		)
	})
	container.MustSingleton(Container, func() domainConverters.ConverterWithExtraArgs[userDomainQueries.UserPaginateFindQuery, *gorm.DB, *gorm.DB] {
		return userGormConverters.NewUserFindQueryToUserGormFindQueryConverter()
	})
	container.MustSingleton(Container, func(userGormToUserConverter domainConverters.Converter[userGormModels.UserGorm, userDomainModels.User]) domainConverters.ConverterWithExtraArgs[[]userGormModels.UserGorm, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[userDomainModels.User]] {
		return userGormConverters.NewUsersGormToPaginatedUsersConverter(userGormToUserConverter)
	})
	container.MustSingleton(Container, func(
		db *gorm.DB,
		userFindQueryToGormUserFindQueryConverter domainConverters.ConverterWithExtraArgs[userDomainQueries.UserPaginateFindQuery, *gorm.DB, *gorm.DB],
		usersToPaginatedUsersConverter domainConverters.ConverterWithExtraArgs[[]userGormModels.UserGorm, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[userDomainModels.User]],
	) commonDomainAdapters.PaginateFindAdapter[userDomainQueries.UserPaginateFindQuery, userDomainModels.User] {
		return userGormAdapters.NewPaginateFindUserGormAdapter(
			db,
			userFindQueryToGormUserFindQueryConverter,
			usersToPaginatedUsersConverter,
		)
	})
}

func initDomainManagers(Container container.Container) {
	container.MustSingleton(Container, func(
		insertOneUserAdapter commonDomainAdapters.InsertOneAdapter[userDomainCommands.UserInsertOneCommand, userDomainModels.User],
	) commonDomainManagers.InsertOneManager[userDomainCommands.UserInsertOneCommand, userDomainModels.User] {
		return userDomainManagers.NewInsertOneUserManager(insertOneUserAdapter)
	})
	container.MustSingleton(Container, func(
		findOneUserAdapter commonDomainAdapters.FindOneAdapter[userDomainQueries.UserFindOneQuery, userDomainModels.User],
	) commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User] {
		return userDomainManagers.NewFindOneUserManager(findOneUserAdapter)
	})
	container.MustSingleton(Container, func(
		paginateFindUserAdapter commonDomainAdapters.PaginateFindAdapter[userDomainQueries.UserPaginateFindQuery, userDomainModels.User],
	) commonDomainManagers.PaginateFindManager[userDomainQueries.UserPaginateFindQuery, userDomainModels.User] {
		return userDomainManagers.NewPaginateFindUserManager(paginateFindUserAdapter)
	})
}

func initApplicationHandlers(Container container.Container) {
	container.MustSingleton(Container, func(
		insertOneUserManager commonDomainManagers.InsertOneManager[userDomainCommands.UserInsertOneCommand, userDomainModels.User],
	) commonApplicationCommandHandlers.InsertOneCommandHandler[userDomainCommands.UserInsertOneCommand, userDomainModels.User] {
		return userApplicationCommandHandlers.NewUserInsertOneCommandHandler(insertOneUserManager)
	})
	container.MustSingleton(Container, func(
		findOneUserManager commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User],
	) commonApplicationQueryHandlers.FindOneQueryHandler[userDomainQueries.UserFindOneQuery, userDomainModels.User] {
		return userApplicationQueryHandlers.NewUserFindOneQueryHandler(findOneUserManager)
	})
	container.MustSingleton(Container, func(
		paginateFindUserManager commonDomainManagers.PaginateFindManager[userDomainQueries.UserPaginateFindQuery, userDomainModels.User],
	) commonApplicationQueryHandlers.PaginateFindQueryHandler[userDomainQueries.UserPaginateFindQuery, userDomainModels.User] {
		return userApplicationQueryHandlers.NewUserPaginateFindQueryHandler(paginateFindUserManager)
	})
}

func initHttpControllers(Container container.Container) {
	container.MustSingleton(Container, func(
		baseHttpController *commonHttpControllers.BaseHttpController,
		userInsertOneCommandHandler commonApplicationCommandHandlers.InsertOneCommandHandler[userDomainCommands.UserInsertOneCommand, userDomainModels.User],
	) *userControllers.InsertOneUserHttpController {
		return userControllers.NewInsertOneUserHttpController(
			baseHttpController,
			userInsertOneCommandHandler,
		)
	})
	container.MustSingleton(Container, func(
		baseHttpController *commonHttpControllers.BaseHttpController,
		userFindOneQueryHandler commonApplicationQueryHandlers.FindOneQueryHandler[userDomainQueries.UserFindOneQuery, userDomainModels.User],
	) *userControllers.FindOneUserHttpController {
		return userControllers.NewFindOneUserHttpController(
			baseHttpController,
			userFindOneQueryHandler,
		)
	})
	container.MustSingleton(Container, func(
		baseHttpController *commonHttpControllers.BaseHttpController,
		userPaginateFindQueryHandler commonApplicationQueryHandlers.PaginateFindQueryHandler[userDomainQueries.UserPaginateFindQuery, userDomainModels.User],
	) *userControllers.PaginateFindUserHttpController {
		return userControllers.NewPaginateFindUserHttpController(
			baseHttpController,
			userPaginateFindQueryHandler,
		)
	})
}

func InitUserDependencies(Container container.Container) {
	initGormAdapters(Container)
	initDomainManagers(Container)
	initApplicationHandlers(Container)
	initHttpControllers(Container)
}
