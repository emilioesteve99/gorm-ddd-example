package commonGormConverters

import (
	"context"
	"github.com/stretchr/testify/assert"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainConverterFixtures "gorm-ddd-example/src/common/fixture/domain/converter"
	"testing"
)

func TestBaseModelsDBToPaginatedModelsConverter_Convert(t *testing.T) {
	modelDbToModelConverterMock := &commonDomainConverterFixtures.ConverterMock[int, int]{}
	converter := NewBaseModelsDBToPaginatedModelsConverter[int, int](modelDbToModelConverterMock)

	t.Run("when called", func(t *testing.T) {
		inputFixture := []int{1}
		paginationContextFixture := commonDomainModels.PaginationContext{}
		ctx := context.TODO()

		modelDbToModelConverterMock.On("Convert", inputFixture[0], ctx).Return(inputFixture[0], nil)

		defer modelDbToModelConverterMock.ResetMock()

		res, err := converter.Convert(inputFixture, paginationContextFixture, ctx)

		t.Run("err should be nil", func(t *testing.T) {
			assert.Nil(t, err)
		})

		t.Run("res should be PaginatedItems[int]", func(t *testing.T) {
			assert.Equal(t, commonDomainModels.PaginatedItems[int]{
				Items: []int{1},
				Meta: commonDomainModels.PaginatedItemsMeta{
					CurrentPage:  0,
					ItemCount:    1,
					ItemsPerPage: 0,
					TotalItems:   0,
					TotalPages:   0,
				},
			}, res)
		})
	})
}
