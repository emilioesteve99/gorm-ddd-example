package commonApplicationCommandHandlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainManagerFixtures "gorm-ddd-example/src/common/fixture/domain/manager"
	"testing"
)

func TestBaseInsertOneCommandHandler_Handle(t *testing.T) {
	managerMock := &commonDomainManagerFixtures.InsertOneManagerMock[int, int]{}
	commandHandler := NewBaseInsertOneCommandHandler[int, int](managerMock)

	t.Run("when called", func(t *testing.T) {
		commandFixture := 0
		ctx := context.TODO()
		outputFixture := 1

		managerMock.On("Manage", commandFixture, ctx).Return(outputFixture, nil).Once()

		defer managerMock.ResetMock()

		res, err := commandHandler.Handle(commandFixture, ctx)

		t.Run("should call .insertOneManager.Manage()", func(t *testing.T) {
			managerMock.AssertNumberOfCalls(t, "Manage", 1)
			managerMock.AssertCalled(t, "Manage", commandFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, outputFixture, res)
		})
	})
}
