package authControllers

import (
	"github.com/gin-gonic/gin"
	authDomainCommands "gorm-ddd-example/src/auth/domain/command"
	commonApplicationCommandHandlers "gorm-ddd-example/src/common/application/command_handler"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	"net/http"
)

type LoginHttpController struct {
	*commonControllers.BaseHttpController
	loginCommandHandler commonApplicationCommandHandlers.ILoginCommandHandler
}

func NewLoginHttpController(
	baseHttpController *commonControllers.BaseHttpController,
	loginCommandHandler commonApplicationCommandHandlers.ILoginCommandHandler,
) *LoginHttpController {
	controller := &LoginHttpController{
		BaseHttpController:  baseHttpController,
		loginCommandHandler: loginCommandHandler,
	}
	commonControllers.RegisterController(controller)
	return controller
}

type loginBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (c *LoginHttpController) Control(ctx *gin.Context) {
	var body loginBody
	if err := c.BindJSONBody(ctx, &body); err != nil {
		return
	}

	command := authDomainCommands.LoginCommand{
		Email:    body.Email,
		Password: body.Password,
	}
	user, err := c.loginCommandHandler.Handle(command, ctx)
	if err != nil {
		c.SendError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *LoginHttpController) Method() string {
	return http.MethodPost
}

func (c *LoginHttpController) Path() string {
	return "/v1/auth/log-ins"
}

func (c *LoginHttpController) IsPrivate() bool {
	return false
}
