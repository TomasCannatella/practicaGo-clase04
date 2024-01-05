package calculateTax_test

import (
	"ejercicio1/internal/calculateTax"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCalculateTax(t *testing.T) {
	t.Run("Gane por debajo de $50000", func(t *testing.T) {
		salary := 49000.0

		result := calculateTax.CalculateTax(salary)

		expectedResult := 0.0

		//Con testing
		// if result != expectedResult {
		// 	t.Fatal("El salario es de ", salary, " y el impuesto es ", result, " y se esperaba un impuesto de: $", expectedResult)
		// }

		//Con testify
		require.Equal(t, result, expectedResult)
	})

	t.Run("Gane por encima de $50000", func(t *testing.T) {
		salary := 51000.0

		result := calculateTax.CalculateTax(salary)

		expectedResult := 8670.0

		//Con testing
		// if result != expectedResult {
		// 	t.Fatal("El salario es de ", salary, " y el impuesto es ", result, " y se esperaba un impuesto de: $", expectedResult)
		// }

		//Con testify
		require.Equal(t, result, expectedResult)
	})

	t.Run("Gane por encima de $150000", func(t *testing.T) {
		salary := 151000.0

		result := calculateTax.CalculateTax(salary)

		expectedResult := 40770.0

		//Con testing
		// if result != expectedResult {
		// 	t.Fatal("El salario es de ", salary, " y el impuesto es ", result, " y se esperaba un impuesto de: $", expectedResult)
		// }

		//Con testify
		require.Equal(t, result, expectedResult)
	})
}
