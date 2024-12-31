package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonMiddlewares "gorm-ddd-example/src/common/infrastructure/http/middleware"
	commonDependencies "gorm-ddd-example/src/common/infrastructure/ioc"
	"gorm-ddd-example/src/metrics"
	userControllers "gorm-ddd-example/src/user/infrastructure/http/controller"
)

func main() {
	commonDependencies.InitDependencies()

	r := gin.Default()
	r.Use(commonMiddlewares.RequestDurationMiddleware())

	var healthcheckController *httpControllers.HealthcheckController
	container.MustResolve(commonDependencies.Container, &healthcheckController)
	r.GET("/healthcheck", healthcheckController.Healthcheck)

	prometheusRegistry := metrics.NewPrometheusRegistry()
	metricsHandler := promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
	metricsHandler = promhttp.InstrumentMetricHandler(
		prometheusRegistry,
		metricsHandler,
	)
	r.GET("/metrics", gin.WrapH(metricsHandler))

	var insertOneUserController userControllers.InsertOneUserController
	container.MustResolve(commonDependencies.Container, &insertOneUserController)
	r.POST("/v1/users", insertOneUserController.InsertOne)

	var findOneUserController userControllers.FindOneUserController
	container.MustResolve(commonDependencies.Container, &findOneUserController)
	r.GET("/v1/users/:id", findOneUserController.FindOne)

	var paginateFindUserController userControllers.PaginateFindUserController
	container.MustResolve(commonDependencies.Container, &paginateFindUserController)
	r.GET("/v1/users", paginateFindUserController.PaginateFind)

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
