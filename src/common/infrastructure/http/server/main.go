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

	var healthcheckHttpController *httpControllers.HealthcheckController
	container.MustResolve(commonDependencies.Container, &healthcheckHttpController)
	var userInsertOneHttpController *userControllers.InsertOneUserHttpController
	container.MustResolve(commonDependencies.Container, &userInsertOneHttpController)
	var userFindOneHttpController *userControllers.FindOneUserHttpController
	container.MustResolve(commonDependencies.Container, &userFindOneHttpController)
	var userPaginateFindHttpController *userControllers.PaginateFindUserHttpController
	container.MustResolve(commonDependencies.Container, &userPaginateFindHttpController)

	r.GET("/healthcheck", healthcheckHttpController.Healthcheck)
	r.POST("/v1/users", userInsertOneHttpController.InsertOne)
	r.GET("/v1/users/:id", userFindOneHttpController.FindOne)
	r.GET("/v1/users", userPaginateFindHttpController.PaginateFind)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
