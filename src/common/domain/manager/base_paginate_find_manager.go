package commonDomainManagers

import (
	"context"
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
)

type BasePaginateFindManager[TQuery commonDomainQueries.PaginateFindQuery, TOutput any] struct {
	paginateFindAdapter commonDomainAdapters.PaginateFindAdapter[TQuery, TOutput]
}

func NewBasePaginateFindManager[TQuery commonDomainQueries.PaginateFindQuery, TOutput any](paginateFindAdapter commonDomainAdapters.PaginateFindAdapter[TQuery, TOutput]) *BasePaginateFindManager[TQuery, TOutput] {
	return &BasePaginateFindManager[TQuery, TOutput]{
		paginateFindAdapter: paginateFindAdapter,
	}
}

func (m *BasePaginateFindManager[TQuery, TOutput]) Manage(query TQuery, ctx context.Context) (commonDomainModels.PaginatedItems[TOutput], error) {
	return m.paginateFindAdapter.PaginateFind(query, ctx)
}
