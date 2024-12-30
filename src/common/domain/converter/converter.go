package domainConverters

import "context"

type Converter[TInput any, TOutput any] interface {
	Convert(input TInput, context context.Context) (TOutput, error)
}

type ConverterWithExtraArgs[TInput any, TExtraArgs any, TOutput any] interface {
	Convert(input TInput, extraArgs TExtraArgs, context context.Context) (TOutput, error)
}
