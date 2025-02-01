package userControllers

import (
	"github.com/gin-gonic/gin"
	commonApplicationQueryHandlers "gorm-ddd-example/src/common/application/query_handler"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	userDomainQueries "gorm-ddd-example/src/user/domain/query"
	"net/http"
)

type FindOneUserController struct {
	*commonControllers.BaseHttpController
	userFindOneQueryHandler commonApplicationQueryHandlers.FindOneQueryHandler[
		userDomainQueries.UserFindOneQuery,
		userDomainModels.User,
	]
}

func NewFindOneUserHttpController(
	baseHttpController *commonControllers.BaseHttpController,
	userFindOneQueryHandler commonApplicationQueryHandlers.FindOneQueryHandler[
		userDomainQueries.UserFindOneQuery,
		userDomainModels.User,
	],
) *FindOneUserController {
	controller := &FindOneUserController{
		BaseHttpController:      baseHttpController,
		userFindOneQueryHandler: userFindOneQueryHandler,
	}
	commonControllers.RegisterController(controller)
	return controller
}

type findOneUserPathParams struct {
	Id string `uri:"id" binding:"required,uuid"`
}

func (c *FindOneUserController) Control(ctx *gin.Context) {
	var pathParams findOneUserPathParams
	if err := c.BindUri(ctx, &pathParams); err != nil {
		return
	}

	ids := []string{pathParams.Id}
	user, err := c.userFindOneQueryHandler.Handle(userDomainQueries.UserFindOneQuery{
		Ids: &ids,
	}, ctx)
	if err != nil {
		c.SendError(ctx, err)
		return
	}

	httpErr := c.Send404ErrIfEntityNotFound(ctx, user)
	if httpErr == nil {
		ctx.JSON(http.StatusOK, user)
	}
}

func (c *FindOneUserController) Method() string {
	return http.MethodGet
}

func (c *FindOneUserController) Path() string {
	return "/v1/users/:id"
}
