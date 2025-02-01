package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	"gorm-ddd-example/src/common/infrastructure/http/metrics"
	commonMiddlewares "gorm-ddd-example/src/common/infrastructure/http/middleware"
	commonDependencies "gorm-ddd-example/src/common/infrastructure/ioc"
	"gorm-ddd-example/src/config"
)

func startMetricsServer(cfg config.Config) {
	if cfg.Metrics.Enabled {
		r := gin.Default()
		prometheusRegistry := metrics.NewPrometheusRegistry()
		metricsHandler := promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
		metricsHandler = promhttp.InstrumentMetricHandler(
			prometheusRegistry,
			metricsHandler,
		)
		r.GET("/metrics", gin.WrapH(metricsHandler))
		err := r.Run(fmt.Sprintf(":%d", cfg.Metrics.Port))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	commonDependencies.InitDependencies()

	r := gin.Default()
	r.Use(commonMiddlewares.RequestDurationMiddleware())

	commonControllers.RegisterServerRoutes(r)

	var cfg config.Config
	container.MustResolve(commonDependencies.Container, &cfg)

	go func() {
		startMetricsServer(cfg)
	}()
	err := r.Run(fmt.Sprintf(":%d", cfg.Server.Port))
	if err != nil {
		panic(err)
	}
}
