package commonApplicationQueryHandlers

import (
	"context"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
)

type PaginateFindQueryHandler[TQuery commonDomainQueries.PaginateFindQuery, TOutput any] interface {
	Handle(input TQuery, ctx context.Context) (commonDomainModels.PaginatedItems[TOutput], error)
}
