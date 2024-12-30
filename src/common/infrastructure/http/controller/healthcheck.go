package commonHttpControllers

import "github.com/gin-gonic/gin"

type HealthcheckController struct{}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (controller *HealthcheckController) Healthcheck(c *gin.Context) {
	c.JSON(200, "Ok")
}
