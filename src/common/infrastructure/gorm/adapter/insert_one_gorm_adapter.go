package commonGormAdapters

import (
	"context"
	"errors"
	appErrors "gorm-ddd-example/src/common/application/model"
	domainConverters "gorm-ddd-example/src/common/domain/converter"
	"gorm.io/gorm"
)

type InsertOneGormAdapter[TCommand any, TModelDB any, TModel any] struct {
	db                                                *gorm.DB
	anyEntityInsertOneCommandToAnyEntityGormConverter domainConverters.Converter[TCommand, TModelDB]
	anyEntityGormToAnyEntityConverter                 domainConverters.Converter[TModelDB, TModel]
}

func NewInsertOneGormAdapter[TCommand any, TModelDB any, TModel any](
	db *gorm.DB,
	anyEntityInsertOneCommandToAnyEntityGormConverter domainConverters.Converter[TCommand, TModelDB],
	anyEntityGormToAnyEntityConverter domainConverters.Converter[TModelDB, TModel],
) *InsertOneGormAdapter[TCommand, TModelDB, TModel] {
	return &InsertOneGormAdapter[TCommand, TModelDB, TModel]{
		db: db,
		anyEntityInsertOneCommandToAnyEntityGormConverter: anyEntityInsertOneCommandToAnyEntityGormConverter,
		anyEntityGormToAnyEntityConverter:                 anyEntityGormToAnyEntityConverter,
	}
}

func (a *InsertOneGormAdapter[TCommand, TModelDB, TModel]) InsertOne(input TCommand, ctx context.Context) (TModel, error) {
	anyEntityGorm, convertingToAnyEntityGormErr := a.anyEntityInsertOneCommandToAnyEntityGormConverter.Convert(input, ctx)
	if convertingToAnyEntityGormErr != nil {
		return *new(TModel), convertingToAnyEntityGormErr
	}

	tx := a.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return *new(TModel), a.convertGormErrorToAppError(err)
	}

	if err := tx.Create(anyEntityGorm).Error; err != nil {
		return *new(TModel), a.convertGormErrorToAppError(err)
	}

	if err := tx.Commit().Error; err != nil {
		return *new(TModel), a.convertGormErrorToAppError(err)
	}

	anyEntity, convertingToAnyEntityErr := a.anyEntityGormToAnyEntityConverter.Convert(anyEntityGorm, ctx)
	if convertingToAnyEntityErr != nil {
		return *new(TModel), convertingToAnyEntityErr
	}

	return anyEntity, nil
}

func (a *InsertOneGormAdapter[TCommand, TModelDB, TModel]) convertGormErrorToAppError(gormErr error) error {
	err := appErrors.BuildUnknownError(gormErr)
	if errors.Is(gormErr, gorm.ErrDuplicatedKey) {
		err = appErrors.AppError{
			Code:    appErrors.InvalidArgumentCode,
			Message: appErrors.DuplicatedEntityMsg,
		}
	} else if errors.Is(gormErr, gorm.ErrForeignKeyViolated) {
		err = appErrors.AppError{
			Code:    appErrors.InvalidArgumentCode,
			Message: appErrors.ForeignKeyViolationMsg,
		}
	}
	return err
}
