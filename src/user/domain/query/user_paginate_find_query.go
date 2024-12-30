package userDomainQueries

import commonDomainQueries "gorm-ddd-example/src/common/domain/query"

type UserPaginateFindQuery struct {
	commonDomainQueries.BasePaginateFindQuery
	Query UserFindQuery
}
