package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	item := 3
	t.Run("when called", func(t *testing.T) {
		res := Contains(slice, item)

		t.Run("should return true", func(t *testing.T) {
			assert.True(t, res)
		})
	})
}
