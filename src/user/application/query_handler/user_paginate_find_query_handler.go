package userApplicationQueryHandlers

import (
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
)

type UserPaginateFindQueryHandler struct {
	commonApplicationQueryHandlers.PaginateFindQueryHandler[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]
}

func NewUserPaginateFindQueryHandler(paginateFindUserManager commonDomainManagers.PaginateFindManager[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]) *UserPaginateFindQueryHandler {
	return &UserPaginateFindQueryHandler{
		PaginateFindQueryHandler: commonApplicationQueryHandlers.NewBasePaginateFindQueryHandler(paginateFindUserManager),
	}
}
