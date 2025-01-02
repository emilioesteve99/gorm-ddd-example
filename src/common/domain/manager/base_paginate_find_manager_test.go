package commonDomainManagers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	commonDomainAdapterFixtures "gorm-ddd-example/src/common/fixture/domain/adapter"
	"testing"
)

func TestBasePaginateFindManager_Manage(t *testing.T) {
	adapterMock := &commonDomainAdapterFixtures.PaginateFindAdapterMock[commonDomainQueries.BasePaginateFindQuery, int]{}
	manager := NewBasePaginateFindManager[commonDomainQueries.BasePaginateFindQuery, int](adapterMock)

	t.Run("when called", func(t *testing.T) {
		queryFixture := commonDomainQueries.BasePaginateFindQuery{}
		ctx := context.TODO()
		outputFixture := commonDomainModels.PaginatedItems[int]{
			Items: []int{1},
		}

		adapterMock.On("PaginateFind", queryFixture, ctx).Return(outputFixture, nil).Once()

		defer adapterMock.ResetMock()

		res, err := manager.Manage(queryFixture, ctx)

		t.Run("should call .paginateFindAdapter.PaginateFind()", func(t *testing.T) {
			adapterMock.AssertNumberOfCalls(t, "PaginateFind", 1)
			adapterMock.AssertCalled(t, "PaginateFind", queryFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, outputFixture, res)
		})
	})
}
