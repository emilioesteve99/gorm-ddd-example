package commonDomainAdapters

import "context"

type DeleteAdapter[TCommand any] interface {
	Delete(input TCommand, context context.Context) error
}
