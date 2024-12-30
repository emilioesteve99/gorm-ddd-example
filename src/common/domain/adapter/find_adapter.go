package commonDomainAdapters

import "context"

type FindAdapter[TQuery any, TOutput any] interface {
	Find(input TQuery, context context.Context) ([]TOutput, error)
}
