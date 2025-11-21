package cultural_test

import (
	"testing"

	"poc2/back/domain/cultural"

	"github.com/stretchr/testify/assert"
)

func TestNewPrice(t *testing.T) {
	t.Run("valid price", func(t *testing.T) {
		price := cultural.NewPrice("R$10,00")
		assert.Equal(t, "R$10,00", price.String())
	})

	t.Run("empty price returns default", func(t *testing.T) {
		price := cultural.NewPrice("")
		assert.Equal(t, "R$0,00", price.String())
	})
}

func TestNewLocation(t *testing.T) {
	t.Run("valid location", func(t *testing.T) {
		loc, err := cultural.NewLocation("Belo Horizonte")
		assert.NoError(t, err)
		assert.Equal(t, "Belo Horizonte", loc.String())
	})

	t.Run("empty location returns error", func(t *testing.T) {
		_, err := cultural.NewLocation("")
		assert.Error(t, err)
		assert.Equal(t, "location cannot be empty", err.Error())
	})
}
