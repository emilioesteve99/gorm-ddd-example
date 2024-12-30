package userDomainManagers

import (
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
)

type FindOneUserManager struct {
	commonDomainManagers.FindOneManager[userDomainQueries.UserFindOneQuery, userDomainModels.User]
}

func NewFindOneUserManager(findOneAdapter commonDomainAdapters.FindOneAdapter[userDomainQueries.UserFindOneQuery, userDomainModels.User]) *FindOneUserManager {
	return &FindOneUserManager{
		FindOneManager: commonDomainManagers.NewBaseFindOneManager[
			userDomainQueries.UserFindOneQuery,
			userDomainModels.User,
		](findOneAdapter),
	}
}
