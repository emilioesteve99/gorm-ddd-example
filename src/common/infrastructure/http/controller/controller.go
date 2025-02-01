package commonControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Control(c *gin.Context)
	Method() string
	Path() string
	IsPrivate() bool
}

var ControllersMapByName = map[string]Controller{}
var PrivateControllersByMethodAndPath = map[string]Controller{}

func RegisterController(controller Controller) {
	key := fmt.Sprintf("%s %s", controller.Method(), controller.Path())
	ControllersMapByName[key] = controller
	if controller.IsPrivate() {
		PrivateControllersByMethodAndPath[key] = controller
	}
}

func RegisterServerRoutes(server *gin.Engine) {
	for _, controller := range ControllersMapByName {
		server.Handle(controller.Method(), controller.Path(), controller.Control)
	}
}
