package commonControllers

import "github.com/gin-gonic/gin"

type HealthcheckController struct{}

func NewHealthcheckController() *HealthcheckController {
	controller := &HealthcheckController{}
	RegisterController(controller)
	return controller
}

func (controller *HealthcheckController) Control(c *gin.Context) {
	c.JSON(200, "Ok")
}

func (controller *HealthcheckController) Method() string {
	return "GET"
}

func (controller *HealthcheckController) Path() string {
	return "/healthcheck"
}
