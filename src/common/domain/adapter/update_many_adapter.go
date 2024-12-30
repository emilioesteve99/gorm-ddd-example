package commonDomainAdapters

import "context"

type UpdateManyAdapter[TCommand any] interface {
	UpdateMany(input TCommand, context context.Context) error
}
