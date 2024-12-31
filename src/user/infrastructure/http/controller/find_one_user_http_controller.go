package userControllers

import (
	"github.com/gin-gonic/gin"
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonHttpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonHttpModels "gorm-ddd-example/src/common/infrastructure/http/model"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"net/http"
)

type FindOneUserController struct {
	*commonHttpControllers.BaseHttpController
	userFindOneQueryHandler commonApplicationQueryHandlers.FindOneQueryHandler[
		userDomainQueries.UserFindOneQuery,
		userDomainModels.User,
	]
}

func NewFindOneUserHttpController(
	baseHttpController *commonHttpControllers.BaseHttpController,
	userFindOneQueryHandler commonApplicationQueryHandlers.FindOneQueryHandler[
		userDomainQueries.UserFindOneQuery,
		userDomainModels.User,
	],
) FindOneUserController {
	return FindOneUserController{
		BaseHttpController:      baseHttpController,
		userFindOneQueryHandler: userFindOneQueryHandler,
	}
}

type findOneUserPathParams struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (c *FindOneUserController) FindOne(ctx *gin.Context) {
	var pathParams findOneUserPathParams
	if err := c.BindUri(ctx, &pathParams); err != nil {
		return
	}

	ids := []string{pathParams.Id}
	user, err := c.userFindOneQueryHandler.Handle(userDomainQueries.UserFindOneQuery{
		Ids: &ids,
	}, ctx)
	if err != nil {
		httpError := commonHttpModels.HttpErrorResponse{Message: err.Error()}
		statusCode := c.ConvertErrorToHttpStatusCode(err)
		ctx.JSON(statusCode, httpError)
		return
	}

	httpErr := c.Send404ErrIfEntityNotFound(ctx, user)
	if httpErr == nil {
		ctx.JSON(http.StatusOK, user)
	}
}
