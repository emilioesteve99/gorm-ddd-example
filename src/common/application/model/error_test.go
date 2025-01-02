package appErrors

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAppError_Error(t *testing.T) {
	codeFixture := ErrorCode(0)
	messageFixture := ErrorMessage("error-example")

	err := AppError{
		Code:    codeFixture,
		Message: messageFixture,
	}

	t.Run("when called", func(t *testing.T) {
		res := err.Error()

		t.Run("should return a string", func(t *testing.T) {
			assert.Equal(t, fmt.Sprintf("%s [Code %d]", messageFixture, codeFixture), res)
		})
	})
}

func TestBuildUnknownError(t *testing.T) {
	errFixture := errors.New("error-example")

	t.Run("when called", func(t *testing.T) {
		res := BuildUnknownError(errFixture)

		t.Run("should return an AppError", func(t *testing.T) {
			assert.Equal(t, UnknownCode, res.Code)
			assert.Equal(t, ErrorMessage(fmt.Sprintf("%s: %s", UnknownMsg, errFixture.Error())), res.Message)
		})
	})
}
