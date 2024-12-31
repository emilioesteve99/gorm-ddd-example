package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	httpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonDependencies "gorm-ddd-example/src/common/infrastructure/ioc"
	userControllers "gorm-ddd-example/src/user/infrastructure/http/controller"
)

func main() {
	commonDependencies.InitDependencies()

	r := gin.Default()

	var healthcheckController *httpControllers.HealthcheckController
	container.MustResolve(commonDependencies.Container, &healthcheckController)
	r.GET("/healthcheck", healthcheckController.Healthcheck)

	var insertOneUserController *userControllers.InsertOneUserController
	container.MustResolve(commonDependencies.Container, &insertOneUserController)
	r.POST("/v1/users", insertOneUserController.InsertOne)

	var findOneUserController *userControllers.FindOneUserController
	container.MustResolve(commonDependencies.Container, &findOneUserController)
	r.GET("/v1/users/:id", findOneUserController.FindOne)

	var paginateFindUserController *userControllers.PaginateFindUserController
	container.MustResolve(commonDependencies.Container, &paginateFindUserController)
	r.GET("/v1/users", paginateFindUserController.PaginateFind)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
