package commonGormAdapters

import (
	"context"
	appErrors "gorm-ddd-example/src/common/application/model"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	commonDomainModels "gorm-ddd-example/src/common/domain/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	"gorm.io/gorm"
	"math"
)

type PaginateFindGormAdapter[TQuery commonDomainQueries.PaginateFindQuery, TModelDB any, TModel any] struct {
	db                                 *gorm.DB
	findQueryToGormFindQueryConverter  domainConverters.ConverterWithExtraArgs[TQuery, *gorm.DB, *gorm.DB]
	modelsDbToPaginatedModelsConverter domainConverters.ConverterWithExtraArgs[[]TModelDB, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[TModel]]
}

func NewPaginateFindGormAdapter[TQuery commonDomainQueries.PaginateFindQuery, TModelDB any, TModel any](
	db *gorm.DB,
	findQueryToGormFindQueryConverter domainConverters.ConverterWithExtraArgs[TQuery, *gorm.DB, *gorm.DB],
	modelsDbToPaginatedModelsConverter domainConverters.ConverterWithExtraArgs[[]TModelDB, commonDomainModels.PaginationContext, commonDomainModels.PaginatedItems[TModel]],
) *PaginateFindGormAdapter[TQuery, TModelDB, TModel] {
	return &PaginateFindGormAdapter[TQuery, TModelDB, TModel]{
		db:                                 db,
		findQueryToGormFindQueryConverter:  findQueryToGormFindQueryConverter,
		modelsDbToPaginatedModelsConverter: modelsDbToPaginatedModelsConverter,
	}
}

func (a *PaginateFindGormAdapter[TQuery, TModelDB, TModel]) PaginateFind(input TQuery, ctx context.Context) (commonDomainModels.PaginatedItems[TModel], error) {
	result := commonDomainModels.PaginatedItems[TModel]{}

	db := a.db.WithContext(ctx)
	cursor, convertToGormFindQueryErr := a.findQueryToGormFindQueryConverter.Convert(input, db, ctx)
	if convertToGormFindQueryErr != nil {
		return result, a.convertGormErrorToAppError(convertToGormFindQueryErr)
	}

	cursor = cursor.Model(new(TModelDB))

	var totalItems int64
	countErr := cursor.Count(&totalItems).Error
	if countErr != nil {
		return result, a.convertGormErrorToAppError(countErr)
	}

	var modelsDb []TModelDB

	limit := input.Limit()
	page := input.Page()
	offset := int(math.Abs(float64(limit * (page - 1))))

	cursor = cursor.Offset(offset).Limit(limit)
	findErr := cursor.Find(&modelsDb).Error
	if findErr != nil {
		return result, a.convertGormErrorToAppError(findErr)
	}

	paginationContext := commonDomainModels.PaginationContext{
		Limit:      limit,
		Page:       page,
		TotalItems: int(totalItems),
	}

	paginatedItems, convertToPaginatedModelsErr := a.modelsDbToPaginatedModelsConverter.Convert(modelsDb, paginationContext, ctx)
	if convertToPaginatedModelsErr != nil {
		return result, convertToPaginatedModelsErr
	}

	return paginatedItems, nil
}

func (a *PaginateFindGormAdapter[TQuery, TModelDB, TModel]) convertGormErrorToAppError(gormErr error) error {
	return appErrors.BuildUnknownError(gormErr)
}
