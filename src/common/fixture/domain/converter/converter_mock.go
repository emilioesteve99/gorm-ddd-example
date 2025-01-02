package commonDomainConverterFixtures

import (
	"context"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type ConverterMock[TInput any, TOutput any] struct {
	testUtils.BaseMock
}

func (m *ConverterMock[TInput, TOutput]) Convert(input TInput, ctx context.Context) (TOutput, error) {
	args := m.Called(input, ctx)
	return args.Get(0).(TOutput), args.Error(1)
}
