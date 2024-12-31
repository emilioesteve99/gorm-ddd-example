package userApplicationQueryHandlers

import (
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
)

type UserFindOneQueryHandler struct {
	commonApplicationQueryHandlers.FindOneQueryHandler[userDomainQueries.UserFindOneQuery, userDomainModels.User]
}

func NewUserFindOneQueryHandler(findOneUserManager commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User]) commonApplicationQueryHandlers.FindOneQueryHandler[userDomainQueries.UserFindOneQuery, userDomainModels.User] {
	return &UserFindOneQueryHandler{
		FindOneQueryHandler: commonApplicationQueryHandlers.NewBaseFindOneQueryHandler[
			userDomainQueries.UserFindOneQuery,
			userDomainModels.User,
		](findOneUserManager),
	}
}
