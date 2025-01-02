package commonDomainManagers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainAdapterFixtures "gorm-ddd-example/src/common/fixture/domain/adapter"
	"testing"
)

func TestBaseInsertOneManager_Manage(t *testing.T) {
	adapterMock := &commonDomainAdapterFixtures.InsertOneAdapterMock[int, int]{}
	manager := NewBaseInsertOneManager[int, int](adapterMock)

	t.Run("when called", func(t *testing.T) {
		commandFixture := 1
		ctx := context.TODO()
		outputFixture := 2

		adapterMock.On("InsertOne", commandFixture, ctx).Return(outputFixture, nil).Once()

		defer adapterMock.ResetMock()

		res, err := manager.Manage(commandFixture, ctx)

		t.Run("should call .insertOneAdapter.InsertOne()", func(t *testing.T) {
			adapterMock.AssertNumberOfCalls(t, "InsertOne", 1)
			adapterMock.AssertCalled(t, "InsertOne", commandFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, outputFixture, res)
		})
	})
}
