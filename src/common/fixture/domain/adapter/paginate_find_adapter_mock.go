package commonDomainAdapterFixtures

import (
	"context"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	testUtils "gorm-ddd-example/src/common/fixture/utils"
)

type PaginateFindAdapterMock[TQuery commonDomainQueries.PaginateFindQuery, TOutput any] struct {
	testUtils.BaseMock
}

func (m *PaginateFindAdapterMock[TQuery, TOutput]) PaginateFind(query TQuery, ctx context.Context) (commonDomainModels.PaginatedItems[TOutput], error) {
	args := m.Called(query, ctx)
	return args.Get(0).(commonDomainModels.PaginatedItems[TOutput]), args.Error(1)
}
