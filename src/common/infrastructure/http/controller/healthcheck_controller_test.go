package commonControllers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHealthcheckController_Method(t *testing.T) {
	controller := NewHealthcheckController()

	t.Run("when called", func(t *testing.T) {
		result := controller.Method()

		t.Run("should return a string", func(t *testing.T) {
			assert.Equal(t, http.MethodGet, result)
		})
	})
}

func TestHealthcheckController_Path(t *testing.T) {
	controller := NewHealthcheckController()

	t.Run("when called", func(t *testing.T) {
		result := controller.Path()

		t.Run("should return a string", func(t *testing.T) {
			assert.Equal(t, "/healthcheck", result)
		})
	})
}
