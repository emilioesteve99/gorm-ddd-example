package commonDomainAdapterFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type FindOneAdapterMock[TQuery any, TOutput any] struct {
	testUtils.BaseMock
}

func (m *FindOneAdapterMock[TQuery, TOutput]) FindOne(query TQuery, ctx context.Context) (*TOutput, error) {
	args := m.Called(query, ctx)
	return args.Get(0).(*TOutput), args.Error(1)
}
