package commonDomainAdapters

import "context"

type CountAdapter[TQuery any] interface {
	Count(input TQuery, context context.Context) (int, error)
}
