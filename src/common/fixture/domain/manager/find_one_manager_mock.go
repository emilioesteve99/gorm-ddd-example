package commonDomainManagerFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type FindOneManagerMock[TQuery any, TOutput any] struct {
	testUtils.BaseMock
}

func (m *FindOneManagerMock[TQuery, TOutput]) Manage(query TQuery, ctx context.Context) (*TOutput, error) {
	args := m.Called(query, ctx)
	return args.Get(0).(*TOutput), args.Error(1)
}
