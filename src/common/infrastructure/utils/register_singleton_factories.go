package commonInfraUtils

import "github.com/golobby/container/v3"

func RegisterSingletonFactories(factories []any, c container.Container) {
	for _, factory := range factories {
		container.MustSingleton(c, factory)
	}
}
