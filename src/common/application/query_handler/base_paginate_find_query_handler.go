package commonApplicationQueryHandlers

import (
	"context"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
)

type BasePaginateFindQueryHandler[TQuery commonDomainQueries.PaginateFindQuery, TOutput any] struct {
	paginateFindManager commonDomainManagers.PaginateFindManager[TQuery, TOutput]
}

func NewBasePaginateFindQueryHandler[TQuery commonDomainQueries.PaginateFindQuery, TOutput any](paginateFindManager commonDomainManagers.PaginateFindManager[TQuery, TOutput]) *BasePaginateFindQueryHandler[TQuery, TOutput] {
	return &BasePaginateFindQueryHandler[TQuery, TOutput]{paginateFindManager: paginateFindManager}
}

func (b *BasePaginateFindQueryHandler[TQuery, TOutput]) Handle(input TQuery, ctx context.Context) (commonDomainModels.PaginatedItems[TOutput], error) {
	return b.paginateFindManager.Manage(input, ctx)
}
