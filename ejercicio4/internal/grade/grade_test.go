package grade_test

import (
	"ejercicio4/internal/grade"
	"testing"
)

func TestGrade(t *testing.T) {
	t.Run("Average Grade", func(t *testing.T) {
		grades := []int{2, 3, 3, 4, 10, 2, 4, 5}

		result := grade.AverageFunc(grades...)

		expectedResult := 4.125

		if result != expectedResult {
			t.Fatal("El promedio esperado era ", expectedResult, " y se obtuvo un promedio de ", result)
		}
	})

	t.Run("Max Grade", func(t *testing.T) {
		grades := []int{2, 3, 3, 4, 10, 2, 4, 5}

		result := grade.MaxFunc(grades...)

		expectedResult := 10.0

		if result != expectedResult {
			t.Fatal("El maximo esperado era ", expectedResult, " y se obtuvo un maximo de ", result)
		}
	})

	t.Run("Min Grade", func(t *testing.T) {
		grades := []int{2, 3, 3, 4, 10, 2, 4, 5}

		result := grade.MinFunc(grades...)

		expectedResult := 2.0

		if result != expectedResult {
			t.Fatal("El minimo esperado era ", expectedResult, " y se obtuvo un minimo de ", result)
		}
	})
}
