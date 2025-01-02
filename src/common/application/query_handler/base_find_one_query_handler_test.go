package commonApplicationQueryHandlers

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainManagerFixtures "gorm-ddd-example/src/common/fixture/domain/manager"
	"testing"
)

func TestBaseFindOneQueryHandler_Handle(t *testing.T) {
	managerMock := &commonDomainManagerFixtures.FindOneManagerMock[int, int]{}
	queryHandler := NewBaseFindOneQueryHandler[int, int](managerMock)

	t.Run("when called", func(t *testing.T) {
		queryFixture := 0
		ctx := context.TODO()
		outputFixture := 1

		managerMock.On("Manage", queryFixture, ctx).Return(&outputFixture, nil).Once()

		defer managerMock.ResetMock()

		res, err := queryHandler.Handle(queryFixture, ctx)

		t.Run("should call .findOneManager.Manage()", func(t *testing.T) {
			managerMock.AssertNumberOfCalls(t, "Manage", 1)
			managerMock.AssertCalled(t, "Manage", queryFixture, ctx)
		})

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("should return res", func(t *testing.T) {
			assert.Equal(t, &outputFixture, res)
		})
	})
}
