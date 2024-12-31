package commonControllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm-ddd-example/src/metrics"
)

type Controller interface {
	Control(c *gin.Context)
	Method() string
	Path() string
}

var ControllersMapByName = map[string]Controller{}

func RegisterController(controller Controller) {
	key := fmt.Sprintf("%s-%s", controller.Path(), controller.Method())
	ControllersMapByName[key] = controller
}

func RegisterServerRoutes(server *gin.Engine) {
	prometheusRegistry := metrics.NewPrometheusRegistry()
	metricsHandler := promhttp.HandlerFor(prometheusRegistry, promhttp.HandlerOpts{})
	metricsHandler = promhttp.InstrumentMetricHandler(
		prometheusRegistry,
		metricsHandler,
	)
	server.GET("/metrics", gin.WrapH(metricsHandler))
	for _, controller := range ControllersMapByName {
		server.Handle(controller.Method(), controller.Path(), controller.Control)
	}
}
