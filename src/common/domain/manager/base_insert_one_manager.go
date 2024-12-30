package commonDomainManagers

import (
	"context"
	commonDomainAdapters "gorm-ddd-example/src/common/domain/adapter"
)

type BaseInsertOneManager[TCommand any, TOutput any] struct {
	insertOneAdapter commonDomainAdapters.InsertOneAdapter[TCommand, TOutput]
}

func NewBaseInsertOneManager[TCommand any, TOutput any](insertOneAdapter commonDomainAdapters.InsertOneAdapter[TCommand, TOutput]) *BaseInsertOneManager[TCommand, TOutput] {
	return &BaseInsertOneManager[TCommand, TOutput]{insertOneAdapter: insertOneAdapter}
}

func (m BaseInsertOneManager[TCommand, TOutput]) Manage(command TCommand, ctx context.Context) (TOutput, error) {
	return m.insertOneAdapter.InsertOne(command, ctx)
}
