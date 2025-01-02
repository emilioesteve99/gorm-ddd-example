package commonDomainManagerFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type InsertOneManagerMock[TCommand any, TOutput any] struct {
	testUtils.BaseMock
}

func (m *InsertOneManagerMock[TCommand, TOutput]) Manage(command TCommand, ctx context.Context) (TOutput, error) {
	args := m.Called(command, ctx)
	return args.Get(0).(TOutput), args.Error(1)
}
