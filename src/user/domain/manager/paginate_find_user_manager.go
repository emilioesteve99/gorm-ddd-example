package userDomainManagers

import (
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
)

type PaginateFindUserManager struct {
	commonDomainManagers.PaginateFindManager[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]
}

func NewPaginateFindUserManager(paginateFindUserAdapter commonDomainAdapters.PaginateFindAdapter[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]) *PaginateFindUserManager {
	return &PaginateFindUserManager{
		PaginateFindManager: commonDomainManagers.NewBasePaginateFindManager(paginateFindUserAdapter),
	}
}
