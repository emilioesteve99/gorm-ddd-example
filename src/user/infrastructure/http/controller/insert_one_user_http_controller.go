package userControllers

import (
	"github.com/gin-gonic/gin"
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	userDomainCommands "gorm-ddd-example/src/user/domain/command"
	userDomainModels "gorm-ddd-example/src/user/domain/model"
	"net/http"
)

type InsertOneUserController struct {
	*commonControllers.BaseHttpController
	userInsertOneCommandHandler commonApplicationCommandHandlers.InsertOneCommandHandler[
		userDomainCommands.UserInsertOneCommand,
		userDomainModels.User,
	]
}

func NewInsertOneUserHttpController(
	baseHttpController *commonControllers.BaseHttpController,
	userInsertOneCommandHandler commonApplicationCommandHandlers.InsertOneCommandHandler[
		userDomainCommands.UserInsertOneCommand,
		userDomainModels.User,
	],
) *InsertOneUserController {
	controller := &InsertOneUserController{
		BaseHttpController:          baseHttpController,
		userInsertOneCommandHandler: userInsertOneCommandHandler,
	}
	commonControllers.RegisterController(controller)
	return controller
}

type insertOneUserBody struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *InsertOneUserController) Control(ctx *gin.Context) {
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
		c.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *InsertOneUserController) Method() string {
	return http.MethodPost
}

func (c *InsertOneUserController) Path() string {
	return "/v1/users"
}

func (c *InsertOneUserController) IsPrivate() bool {
	return false
}
