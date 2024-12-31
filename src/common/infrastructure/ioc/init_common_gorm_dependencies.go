package commonDependencies

import (
	"github.com/golobby/container/v3"
	gormdb "gorm-ddd-example/src/common/infrastructure/gorm/db"
	commonInfraUtils "gorm-ddd-example/src/common/infrastructure/utils"
)

func InitCommonGormDependencies(c container.Container) {
	dbFactories := []any{
		gormdb.NewGormDB,
	}
	commonInfraUtils.RegisterSingletonFactories(dbFactories, c)
}
