package commonHttpModels

import "context"

type HttpRequest interface {
	context.Context
	JSON(code int, obj any)
}
