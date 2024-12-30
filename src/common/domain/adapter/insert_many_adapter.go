package commonDomainAdapters

import "context"

type InsertManyAdapter[TCommand any, TOutput any] interface {
	InsertMany(input TCommand, context context.Context) ([]TOutput, error)
}
