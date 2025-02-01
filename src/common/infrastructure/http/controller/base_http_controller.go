package commonControllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	appErrors "gorm-ddd-example/src/common/application/model"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	commonHttpModels "gorm-ddd-example/src/common/infrastructure/http/model"
	"net/http"
	"reflect"
)

type BaseHttpController struct{}

func NewBaseHttpController() *BaseHttpController {
	return &BaseHttpController{}
}

var httpStatusCodeByErrorCode = map[appErrors.ErrorCode]int{
	appErrors.UnknownCode:         http.StatusInternalServerError,
	appErrors.InvalidArgumentCode: http.StatusConflict,
	appErrors.UnauthorizedCode:    http.StatusUnauthorized,
}

func getJSONFieldName(request any, fieldName string) string {
	result := fieldName
	val := reflect.ValueOf(request).Elem()
	field, _ := val.Type().FieldByName(fieldName)
	jsonTag := field.Tag.Get("json")
	if jsonTag != "" {
		result = jsonTag
	}
	return result
}

func getFormFieldName(request any, fieldName string) string {
	result := fieldName
	val := reflect.ValueOf(request).Elem()
	field, _ := val.Type().FieldByName(fieldName)
	jsonTag := field.Tag.Get("form")
	if jsonTag != "" {
		result = jsonTag
	}
	return result
}

func getUriFieldName(request any, fieldName string) string {
	result := fieldName
	val := reflect.ValueOf(request).Elem()
	field, _ := val.Type().FieldByName(fieldName)
	jsonTag := field.Tag.Get("uri")
	if jsonTag != "" {
		result = jsonTag
	}
	return result
}

func (c *BaseHttpController) ProcessValidationErrors(ctx *gin.Context, instance any, err error, getFieldNameFunc func(instance any, fieldName string) string) *commonHttpModels.HttpErrorResponse {
	var httpError *commonHttpModels.HttpErrorResponse
	httpError = nil
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		message := "Invalid request body"
		for _, validationError := range validationErrors {
			jsonFieldName := getFieldNameFunc(instance, validationError.Field())
			message = "Invalid property " + jsonFieldName
			break
		}
		httpError = &commonHttpModels.HttpErrorResponse{
			Message: message,
		}
		ctx.JSON(http.StatusBadRequest, httpError)
	}
	return httpError
}

func (c *BaseHttpController) BindJSONBody(ctx *gin.Context, instancePointer any) *commonHttpModels.HttpErrorResponse {
	var httpError *commonHttpModels.HttpErrorResponse
	httpError = nil
	if err := ctx.ShouldBindJSON(instancePointer); err != nil {
		httpError = c.ProcessValidationErrors(ctx, instancePointer, err, getJSONFieldName)
	}
	return httpError
}

func (c *BaseHttpController) BindQueryParams(ctx *gin.Context, instancePointer any) *commonHttpModels.HttpErrorResponse {
	var httpError *commonHttpModels.HttpErrorResponse
	httpError = nil
	if err := ctx.ShouldBindQuery(instancePointer); err != nil {
		httpError = c.ProcessValidationErrors(ctx, instancePointer, err, getFormFieldName)
	}
	return httpError
}

func (c *BaseHttpController) BindUri(ctx *gin.Context, instancePointer any) *commonHttpModels.HttpErrorResponse {
	var httpError *commonHttpModels.HttpErrorResponse
	httpError = nil
	if err := ctx.ShouldBindUri(instancePointer); err != nil {
		httpError = c.ProcessValidationErrors(ctx, instancePointer, err, getUriFieldName)
	}
	return httpError
}

func (c *BaseHttpController) ConvertErrorToHttpStatusCode(err error) int {
	httpStatusCode := http.StatusInternalServerError

	var appErr appErrors.AppError
	if errors.As(err, &appErr) {
		if statusCode, exists := httpStatusCodeByErrorCode[appErr.Code]; exists {
			httpStatusCode = statusCode
		}
	}

	return httpStatusCode
}

func (c *BaseHttpController) SendError(ctx *gin.Context, err error) {
	statusCode := c.ConvertErrorToHttpStatusCode(err)
	message := err.Error()
	if statusCode == http.StatusInternalServerError {
		message = "Internal server error"
	}
	httpError := commonHttpModels.HttpErrorResponse{Message: message}
	ctx.JSON(statusCode, httpError)
}

func (c *BaseHttpController) Send404ErrIfEntityNotFound(ctx *gin.Context, entity any) *commonHttpModels.HttpErrorResponse {
	var httpError *commonHttpModels.HttpErrorResponse
	httpError = nil
	if entity == nil {
		httpError = &commonHttpModels.HttpErrorResponse{Message: "Not found"}
		ctx.JSON(http.StatusNotFound, *httpError)
	}
	return httpError
}

func (c *BaseHttpController) BuildPaginationOptions(page *int, limit *int) commonDomainQueries.PaginationOptions {
	result := commonDomainQueries.PaginationOptions{
		Page:  1,
		Limit: 10,
	}
	if page != nil {
		result.Page = *page
	}
	if limit != nil {
		result.Limit = *limit
	}
	return result
}
