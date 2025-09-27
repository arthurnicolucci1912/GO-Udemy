package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	prod, err := NewProduct("Ps5", 3500.00)
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.NotEmpty(t, prod.ID)
	assert.Equal(t, "Ps5", prod.Name)
	assert.Equal(t, 3500.00, prod.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	prod, err := NewProduct("", 10)
	assert.Nil(t, prod)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	prod, err := NewProduct("Ps5", -10)
	assert.Nil(t, prod)
	assert.Equal(t, ErrInvalidPrice, err)
}

func TestProductWhenPriceIsRequerid(t *testing.T) {
	prod, err := NewProduct("Ps5", 0)
	assert.Nil(t, prod)
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductValidate(t *testing.T) {
	prod, err := NewProduct("Ps5", 10)
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.Nil(t, prod.Validate(), err)
}
