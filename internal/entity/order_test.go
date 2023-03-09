package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfIdIsBlank(t *testing.T) {
	order := &Order{}

	assert.Error(t, order.Validate(), "Order id is required")
}

func TestIfPriceIsBlank(t *testing.T) {
	order := &Order{}

	order.ID = "ID"

	assert.Error(t, order.Validate(), "Order price is required")
}

func TestIfTaxIsBlank(t *testing.T) {
	order := &Order{}

	order.ID = "ID"
	order.Price = 15.55

	assert.Error(t, order.Validate(), "Order tax is required")
}

func TestValidParams(t *testing.T) {
	order := &Order{
		ID:    "123",
		Price: 10.0,
		Tax:   5.0,
	}

	assert.NoError(t, order.Validate())

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)
}

func TestCalculateFinalPriceValidParams(t *testing.T) {
	order := &Order{
		ID:    "123",
		Price: 10.0,
		Tax:   5.0,
	}

	assert.NoError(t, order.Validate())

	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 5.0, order.Tax)

	order.CalculateFinalPrice()

	assert.Equal(t, 15.0, order.FinalPrice)
}

func TestCalculateFinalPriceError(t *testing.T) {
	order := &Order{
		ID:    "123",
		Price: 0.0,
		Tax:   0.0,
	}

	assert.Error(t, order.CalculateFinalPrice(), "Order price is required")
}
