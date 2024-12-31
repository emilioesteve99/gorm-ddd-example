package commonDependencies

import (
	"github.com/golobby/container/v3"
	commonControllers "gorm-ddd-example/src/common/infrastructure/http/controller"
	commonInfraUtils "gorm-ddd-example/src/common/infrastructure/utils"
	userioc "gorm-ddd-example/src/user/infrastructure/ioc"
)

var Container = container.Global

func InitDependencies() {
	controllers := []any{
		commonControllers.NewBaseHttpController,
		commonControllers.NewHealthcheckController,
	}
	commonInfraUtils.RegisterSingletonFactories(controllers, Container)

	InitCommonGormDependencies(Container)
	userioc.InitUserDependencies(Container)
}
