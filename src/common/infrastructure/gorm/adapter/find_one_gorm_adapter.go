package commonGormAdapters

import (
	"context"
	appErrors "gorm-ddd-example/src/common/application/model"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	"gorm.io/gorm"
)

type FindOneGormAdapter[TQuery any, TModelDB any, TModel any] struct {
	db                                                        *gorm.DB
	anyEntityFindOneQueryToAnyEntityGormFindOneQueryConverter domainConverters.ConverterWithExtraArgs[TQuery, *gorm.DB, *gorm.DB]
	anyEntityGormToAnyEntityConverter                         domainConverters.Converter[TModelDB, TModel]
}

func NewFindOneGormAdapter[TQuery any, TModelDB any, TModel any](
	db *gorm.DB,
	anyEntityFindOneQueryToAnyEntityGormFindOneQueryConverter domainConverters.ConverterWithExtraArgs[TQuery, *gorm.DB, *gorm.DB],
	anyEntityGormToAnyEntityConverter domainConverters.Converter[TModelDB, TModel],
) *FindOneGormAdapter[TQuery, TModelDB, TModel] {
	return &FindOneGormAdapter[TQuery, TModelDB, TModel]{
		db: db,
		anyEntityFindOneQueryToAnyEntityGormFindOneQueryConverter: anyEntityFindOneQueryToAnyEntityGormFindOneQueryConverter,
		anyEntityGormToAnyEntityConverter:                         anyEntityGormToAnyEntityConverter,
	}
}

func (a *FindOneGormAdapter[TQuery, TModelDB, TModel]) FindOne(input TQuery, ctx context.Context) (*TModel, error) {
	var result *TModel
	result = nil

	var anyEntityGorm TModelDB

	db := a.db.WithContext(ctx)
	cursor, convertToAnyEntityGormFindOneQueryErr := a.anyEntityFindOneQueryToAnyEntityGormFindOneQueryConverter.Convert(input, db, ctx)
	if convertToAnyEntityGormFindOneQueryErr != nil {
		return result, a.convertGormErrorToAppError(convertToAnyEntityGormFindOneQueryErr)
	}

	cursor = cursor.Limit(1)

	findErr := cursor.Find(&anyEntityGorm).Error
	if findErr != nil {
		return result, a.convertGormErrorToAppError(findErr)
	}

	if cursor.RowsAffected == 0 {
		return result, nil
	}

	anyEntity, convertToAnyEntityErr := a.anyEntityGormToAnyEntityConverter.Convert(anyEntityGorm, ctx)
	if convertToAnyEntityErr != nil {
		return result, convertToAnyEntityErr
	}

	result = &anyEntity
	return result, nil
}

func (a *FindOneGormAdapter[TQuery, TModelDB, TModel]) convertGormErrorToAppError(gormErr error) error {
	return appErrors.BuildUnknownError(gormErr)
}
