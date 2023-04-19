package strategy

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSendRequest(t *testing.T) {
	t.Run("if http can give a response", func(t *testing.T) {
		s, err := SendRequest("https://www.google.com", 3)

		assert.NoError(t, err)
		assert.NotEmpty(t, s)
	})

	t.Run("if http can't give a response", func(t *testing.T) {
		s, err := SendRequest("https://www.dionofrizal.com", 3)

		assert.Error(t, err)
		assert.Empty(t, s)
	})
}
