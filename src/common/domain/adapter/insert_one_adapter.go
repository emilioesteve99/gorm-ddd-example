package commonDomainAdapters

import "context"

type InsertOneAdapter[TCommand any, TOutput any] interface {
	InsertOne(input TCommand, context context.Context) (TOutput, error)
}
