package commonGormConverters

import (
	"context"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	"math"
)

type BaseModelsDBToPaginatedModelsConverter[TModelDB any, TModel any] struct {
	modelDbToModelConverter domainConverters.Converter[TModelDB, TModel]
}

func NewBaseModelsDBToPaginatedModelsConverter[TModelDB any, TModel any](
	modelDbToModelConverter domainConverters.Converter[TModelDB, TModel],
) *BaseModelsDBToPaginatedModelsConverter[TModelDB, TModel] {
	return &BaseModelsDBToPaginatedModelsConverter[TModelDB, TModel]{
		modelDbToModelConverter: modelDbToModelConverter,
	}
}

func (c *BaseModelsDBToPaginatedModelsConverter[TModelDB, TModel]) Convert(input []TModelDB, paginationContext commonDomainModels.PaginationContext, ctx context.Context) (commonDomainModels.PaginatedItems[TModel], error) {
	var result commonDomainModels.PaginatedItems[TModel]
	totalPages := int(math.Ceil(float64(paginationContext.TotalItems) / float64(paginationContext.Limit)))
	if totalPages < 0 {
		totalPages = 0
	}
	result = commonDomainModels.PaginatedItems[TModel]{
		Items: []TModel{},
		Meta: commonDomainModels.PaginatedItemsMeta{
			CurrentPage:  paginationContext.Page,
			ItemCount:    len(input),
			ItemsPerPage: paginationContext.Limit,
			TotalItems:   paginationContext.TotalItems,
			TotalPages:   totalPages,
		},
	}

	for _, modelDb := range input {
		model, convertErr := c.modelDbToModelConverter.Convert(modelDb, ctx)
		if convertErr == nil {
			result.Items = append(result.Items, model)
		}
	}

	return result, nil
}
