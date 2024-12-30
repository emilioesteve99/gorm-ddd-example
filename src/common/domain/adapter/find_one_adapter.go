package commonDomainAdapters

import "context"

type FindOneAdapter[TQuery any, TOutput any] interface {
	FindOne(input TQuery, context context.Context) (*TOutput, error)
}
