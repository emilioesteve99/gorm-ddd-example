package commonDomainManagers

import (
	"context"
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
)

type BaseFindOneManager[TQuery any, TOutput any] struct {
	findOneAdapter commonDomainAdapters.FindOneAdapter[TQuery, TOutput]
}

func NewBaseFindOneManager[TQuery any, TOutput any](findOneAdapter commonDomainAdapters.FindOneAdapter[TQuery, TOutput]) *BaseFindOneManager[TQuery, TOutput] {
	return &BaseFindOneManager[TQuery, TOutput]{findOneAdapter: findOneAdapter}
}

func (m BaseFindOneManager[TQuery, TOutput]) Manage(query TQuery, ctx context.Context) (*TOutput, error) {
	return m.findOneAdapter.FindOne(query, ctx)
}
