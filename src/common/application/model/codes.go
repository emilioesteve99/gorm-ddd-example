package appErrors

type ErrorCode int

const (
	UnknownCode         ErrorCode = 0
	InvalidArgumentCode ErrorCode = 1
	UnauthorizedCode    ErrorCode = 2
)
