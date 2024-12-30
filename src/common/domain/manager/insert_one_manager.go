package commonDomainManagers

import "context"

type InsertOneManager[TCommand any, TOutput any] interface {
	Manage(command TCommand, ctx context.Context) (TOutput, error)
}
