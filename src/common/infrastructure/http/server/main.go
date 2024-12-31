package main

import (
	"github.com/gin-gonic/gin"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonMiddlewares "gorm-ddd-example/src/common/infrastructure/http/middleware"
	commonDependencies "gorm-ddd-example/src/common/infrastructure/ioc"
)

func main() {
	commonDependencies.InitDependencies()

	r := gin.Default()
	r.Use(commonMiddlewares.RequestDurationMiddleware())

	commonControllers.RegisterServerRoutes(r)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
