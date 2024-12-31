package userControllers

import (
	"github.com/gin-gonic/gin"
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonDomainQueries "gorm-ddd-example/src/common/domain/query"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonHttpModels "gorm-ddd-example/src/common/infrastructure/http/model"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"net/http"
)

type PaginateFindUserController struct {
	*commonControllers.BaseHttpController
	userPaginateFindQueryHandler commonApplicationQueryHandlers.PaginateFindQueryHandler[userDomainQueries.UserPaginateFindQuery, userDomainModels.User]
}

func NewPaginateFindUserHttpController(
	baseHttpController *commonControllers.BaseHttpController,
	userPaginateFindQueryHandler commonApplicationQueryHandlers.PaginateFindQueryHandler[userDomainQueries.UserPaginateFindQuery, userDomainModels.User],
) *PaginateFindUserController {
	controller := &PaginateFindUserController{
		BaseHttpController:           baseHttpController,
		userPaginateFindQueryHandler: userPaginateFindQueryHandler,
	}
	commonControllers.RegisterController(controller)
	return controller
}

type paginateFindUserQueryParams struct {
	Email *string `form:"email" binding:"omitempty,email"`
	Page  *int    `form:"page" binding:"omitempty,min=1"`
	Limit *int    `form:"limit" binding:"omitempty,min=1"`
}

func (c *PaginateFindUserController) Control(ctx *gin.Context) {
	var queryParams paginateFindUserQueryParams
	if err := c.BindQueryParams(ctx, &queryParams); err != nil {
		return
	}

	query := userDomainQueries.UserPaginateFindQuery{
		BasePaginateFindQuery: commonDomainQueries.BasePaginateFindQuery{
			PaginationOptions: c.BuildPaginationOptions(queryParams.Page, queryParams.Limit),
		},
		Query: userDomainQueries.UserFindQuery{
			Email: queryParams.Email,
		},
	}
	paginatedUsers, err := c.userPaginateFindQueryHandler.Handle(query, ctx)
	if err != nil {
		httpError := commonHttpModels.HttpErrorResponse{Message: err.Error()}
		statusCode := c.ConvertErrorToHttpStatusCode(err)
		ctx.JSON(statusCode, httpError)
		return
	}

	ctx.JSON(http.StatusCreated, paginatedUsers)
}

func (c *PaginateFindUserController) Method() string {
	return http.MethodGet
}

func (c *PaginateFindUserController) Path() string {
	return "/v1/users"
}
