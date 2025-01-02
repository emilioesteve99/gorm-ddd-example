package commonApplicationQueryHandlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	commonDomainManagerFixtures "gorm-ddd-example/src/common/fixture/domain/manager"
	"testing"
)

func TestBasePaginateFindQueryHandler_Handle(t *testing.T) {
	managerMock := &commonDomainManagerFixtures.PaginateFindManagerMock[commonDomainQueries.BasePaginateFindQuery, int]{}
	queryHandler := NewBasePaginateFindQueryHandler[commonDomainQueries.BasePaginateFindQuery, int](managerMock)

	t.Run("when called", func(t *testing.T) {
		queryFixture := commonDomainQueries.BasePaginateFindQuery{}
		ctx := context.TODO()
		outputFixture := commonDomainModels.PaginatedItems[int]{
			Items: []int{1},
		}

		managerMock.On("Manage", queryFixture, ctx).Return(outputFixture, nil).Once()

		defer managerMock.ResetMock()

		res, err := queryHandler.Handle(queryFixture, ctx)

		t.Run("should call .paginateFindManager.Manage()", func(t *testing.T) {
			managerMock.AssertNumberOfCalls(t, "Manage", 1)
			managerMock.AssertCalled(t, "Manage", queryFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, outputFixture, res)
		})
	})
}
