package commonDomainAdapterFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type InsertOneAdapterMock[TCommand any, TOutput any] struct {
	testUtils.BaseMock
}

func (m *InsertOneAdapterMock[TCommand, TOutput]) InsertOne(command TCommand, ctx context.Context) (TOutput, error) {
	args := m.Called(command, ctx)
	return args.Get(0).(TOutput), args.Error(1)
}
