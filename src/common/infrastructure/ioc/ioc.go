package commonDependencies

import (
	"github.com/golobby/container/v3"
	httpControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	userioc "gorm-ddd-example/src/user/infrastructure/ioc"
)

var Container container.Container

func InitDependencies() {
	Container = container.Global

	container.MustSingleton(Container, func() *httpControllers.BaseHttpController {
		return httpControllers.NewBaseHttpController()
	})
	container.MustSingleton(Container, func() *httpControllers.HealthcheckController {
		return httpControllers.NewHealthcheckController()
	})

	InitCommonGormDependencies(Container)
	userioc.InitUserDependencies(Container)
}
