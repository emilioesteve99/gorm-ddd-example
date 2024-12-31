package commonDependencies

import (
	"github.com/golobby/container/v3"
	commonInfraUtils "gorm-ddd-example/src/common/infrastructure/utils"
	"gorm-ddd-example/src/config"
)

func InitConfigDependencies(c container.Container) {
	configFactories := []any{
		config.GetConfig,
	}
	commonInfraUtils.RegisterSingletonFactories(configFactories, c)
}
