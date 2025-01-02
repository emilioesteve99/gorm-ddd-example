package commonDomainManagers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainAdapterFixtures "gorm-ddd-example/src/common/fixture/domain/adapter"
	"testing"
)

func TestBaseFindOneManager_Manage(t *testing.T) {
	adapterMock := &commonDomainAdapterFixtures.FindOneAdapterMock[int, int]{}
	manager := NewBaseFindOneManager[int, int](adapterMock)

	t.Run("when called", func(t *testing.T) {
		queryFixture := 1
		ctx := context.TODO()
		outputFixture := 2

		adapterMock.On("FindOne", queryFixture, ctx).Return(&outputFixture, nil).Once()

		defer adapterMock.ResetMock()

		res, err := manager.Manage(queryFixture, ctx)

		t.Run("should call .findOneAdapter.FindOne()", func(t *testing.T) {
			adapterMock.AssertNumberOfCalls(t, "FindOne", 1)
			adapterMock.AssertCalled(t, "FindOne", queryFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, &outputFixture, res)
		})
	})
}
