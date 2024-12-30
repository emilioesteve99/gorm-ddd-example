package commonDomainManagers

import "context"

type FindOneManager[TQuery any, TOutput any] interface {
	Manage(command TQuery, ctx context.Context) (*TOutput, error)
}
