package tax

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	//teste para valor 1000
	tax := CalculateTax(1000.00)
	assert.Equal(t, 10.0, tax, "O imposto para 0 deve ser 10")

	//teste para 0
	tax = CalculateTax(0.0)
	assert.Equal(t, 0.0, tax, "O imposto para 0 deve ser 0")

	//negativo
	//teste para 0
	tax = CalculateTax(-100.0)
	assert.Equal(t, 0.0, tax, "O imposto para -100 deve ser 0")

	//teste para valor 500
	tax = CalculateTax(500.00)
	assert.Equal(t, 0.0, tax, "O imposta para 500 deve ser 5")
}
