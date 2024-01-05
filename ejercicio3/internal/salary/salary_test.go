package salary_test

import (
	"ejercicio3/internal/salary"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSalaryCalculate(t *testing.T) {
	t.Run("success - case category C", func(t *testing.T) {
		workMinutes := 120
		category := "C"

		result, err := salary.CalculateSalary(workMinutes, category)
		expectedResult := 2000.0

		// if result != expectedResult {
		// 	t.Fatal("El salario obtenido es: ", result, " y se esperaba un salario $", expectedResult)
		// }
		require.NoError(t, err)
		require.Equal(t, result, expectedResult)
	})

	t.Run("success - category B", func(t *testing.T) {
		workMinutes := 120
		category := "B"

		result, err := salary.CalculateSalary(workMinutes, category)
		expectedResult := 3600.0

		// if result != expectedResult {
		// 	t.Fatal("El salario obtenido es: ", result, " y se esperaba un salario $", expectedResult)
		// }
		require.NoError(t, err)
		require.Equal(t, result, expectedResult)
	})

	t.Run("success - category A", func(t *testing.T) {
		workMinutes := 120
		category := "A"

		result, err := salary.CalculateSalary(workMinutes, category)
		expectedResult := 9000.0

		// if result != expectedResult {
		// 	t.Fatal("El salario obtenido es: ", result, " y se esperaba un salario $", expectedResult)
		// }
		require.NoError(t, err)
		require.Equal(t, result, expectedResult)
	})

	t.Run("failure - category not exist", func(t *testing.T) {
		workMinutes := 120
		category := "Z"

		result, err := salary.CalculateSalary(workMinutes, category)
		expectedResult := 0.0
		expectedError := errors.New("this category not exist")
		// if result != expectedResult {
		// 	t.Fatal("El salario obtenido es: ", result, " y se esperaba un salario $", expectedResult)
		// }
		if assert.Error(t, err) {
			require.Equal(t, err, expectedError)
		}
		require.Equal(t, result, expectedResult)
	})
}
