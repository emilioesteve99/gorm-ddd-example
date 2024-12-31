package userControllers

import (
	"github.com/gin-gonic/gin"
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
	commonHttpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonHttpModels "gorm-ddd-example/src/common/infrastructure/http/model"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	"net/http"
)

type InsertOneUserController struct {
	*commonHttpControllers.BaseHttpController
	userInsertOneCommandHandler commonApplicationCommandHandlers.InsertOneCommandHandler[
		userDomainCommands.UserInsertOneCommand,
		userDomainModels.User,
	]
}

func NewInsertOneUserHttpController(
	baseHttpController *commonHttpControllers.BaseHttpController,
	userInsertOneCommandHandler commonApplicationCommandHandlers.InsertOneCommandHandler[
		userDomainCommands.UserInsertOneCommand,
		userDomainModels.User,
	],
) InsertOneUserController {
	return InsertOneUserController{
		BaseHttpController:          baseHttpController,
		userInsertOneCommandHandler: userInsertOneCommandHandler,
	}
}

type insertOneUserBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *InsertOneUserController) InsertOne(ctx *gin.Context) {
	var body insertOneUserBody
	if err := c.BindJSONBody(ctx, &body); err != nil {
		return
	}

	command := userDomainCommands.UserInsertOneCommand{
		Email:    body.Email,
		Name:     body.Name,
		Password: body.Password,
	}
	user, err := c.userInsertOneCommandHandler.Handle(command, ctx)
	if err != nil {
		httpError := commonHttpModels.HttpErrorResponse{Message: err.Error()}
		statusCode := c.ConvertErrorToHttpStatusCode(err)
		ctx.JSON(statusCode, httpError)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
