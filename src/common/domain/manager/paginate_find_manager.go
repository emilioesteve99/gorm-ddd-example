package commonDomainManagers

import (
	"context"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
)

type PaginateFindManager[TQuery commonDomainQueries.PaginateFindQuery, TOutput any] interface {
	Manage(query TQuery, context context.Context) (commonDomainModels.PaginatedItems[TOutput], error)
}
