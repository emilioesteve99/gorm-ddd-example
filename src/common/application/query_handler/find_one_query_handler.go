package commonApplicationQueryHandlers

import "context"

type FindOneQueryHandler[TQuery any, TOutput any] interface {
	Handle(input TQuery, ctx context.Context) (*TOutput, error)
}
