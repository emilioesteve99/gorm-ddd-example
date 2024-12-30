package commonApplicationQueryHandlers

import (
	"context"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
)

type BaseFindOneQueryHandler[TQuery any, TOutput any] struct {
	findOneManager commonDomainManagers.FindOneManager[TQuery, TOutput]
}

func NewBaseFindOneQueryHandler[TQuery any, TOutput any](findOneManager commonDomainManagers.FindOneManager[TQuery, TOutput]) *BaseFindOneQueryHandler[TQuery, TOutput] {
	return &BaseFindOneQueryHandler[TQuery, TOutput]{findOneManager: findOneManager}
}

func (b *BaseFindOneQueryHandler[TQuery, TOutput]) Handle(input TQuery, ctx context.Context) (*TOutput, error) {
	return b.findOneManager.Manage(input, ctx)
}
