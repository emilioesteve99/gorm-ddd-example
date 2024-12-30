package commonApplicationCommandHandlers

import (
	"context"
	commonDomainManagers "gorm-ddd-example/src/common/domain/manager"
)

type BaseInsertOneCommandHandler[TCommand any, TOutput any] struct {
	insertOneManager commonDomainManagers.InsertOneManager[TCommand, TOutput]
}

func NewBaseInsertOneCommandHandler[TCommand any, TOutput any](insertOneManager commonDomainManagers.InsertOneManager[TCommand, TOutput]) *BaseInsertOneCommandHandler[TCommand, TOutput] {
	return &BaseInsertOneCommandHandler[TCommand, TOutput]{insertOneManager: insertOneManager}
}

func (b BaseInsertOneCommandHandler[TCommand, TOutput]) Handle(input TCommand, ctx context.Context) (TOutput, error) {
	return b.insertOneManager.Manage(input, ctx)
}
