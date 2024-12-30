package commonApplicationCommandHandlers

import "context"

type InsertOneCommandHandler[TCommand any, TOutput any] interface {
	Handle(input TCommand, ctx context.Context) (TOutput, error)
}
