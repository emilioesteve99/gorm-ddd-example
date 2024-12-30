package commonDomainAdapters

import "context"

type UpdateOneAdapter[TCommand any] interface {
	UpdateOne(input TCommand, context context.Context) error
}
