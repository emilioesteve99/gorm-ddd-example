package appErrors

import "fmt"

type AppError struct {
	Code    ErrorCode
	Message ErrorMessage
}

func (a AppError) Error() string {
	return fmt.Sprintf("%s [Code %d]", a.Message, a.Code)
}

func BuildUnknownError(err error) AppError {
	return AppError{
		Code:    UnknownCode,
		Message: ErrorMessage(fmt.Sprintf("%s: %s", UnknownMsg, err.Error())),
	}
}
