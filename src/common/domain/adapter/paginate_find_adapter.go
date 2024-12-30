package commonDomainAdapters

import (
	"context"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
)

type PaginateFindAdapter[TQuery any, TOutput any] interface {
	PaginateFind(input TQuery, context context.Context) (commonDomainModels.PaginatedItems[TOutput], error)
}
