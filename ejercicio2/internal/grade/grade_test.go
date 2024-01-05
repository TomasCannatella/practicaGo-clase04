package grade_test

import (
	"ejercicio2/internal/grade"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAvgGrade(t *testing.T) {
	t.Run("Promedio de notas con notas mayores a 0", func(t *testing.T) {
		grades := []int{10, 7, 8, 8, 8, 7, 8, 9, 10}

		result := grade.AvgGrade(grades...)

		expectedResult := 8.333333333333334

		//Con testing
		// if result != expectedResult {
		// 	t.Fatal("El promedio de nota esperado era ", expectedResult, "y obtuvimos ", result)
		// }

		//Con testify
		require.Equal(t, result, expectedResult)
	})
}
