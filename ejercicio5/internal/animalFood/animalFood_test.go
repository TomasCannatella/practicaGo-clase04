package animalFood_test

import (
	"ejercicio5/internal/animalFood"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDog(t *testing.T) {
	animalQuantity := 10

	result := animalFood.AnimalDog(animalQuantity)

	expectedResult := 100.0

	//Con la libreria testing nativa de go
	// if result != expectedResult {
	// 	t.Fatal("La comida esperada para el perro era ", expectedResult, " y nos devolvio ", result)
	// }

	//Con testify
	require.Equal(t, result, expectedResult)
}

func TestCat(t *testing.T) {
	animalQuantity := 10

	result := animalFood.AnimalCat(animalQuantity)

	expectedResult := 50.0

	//Con la libreria testing nativa de go
	// if result != expectedResult {
	// 	t.Fatal("La comida esperada para el gato era ", expectedResult, " y nos devolvio ", result)
	// }

	//Con testify
	require.Equal(t, result, expectedResult)
}

func TestHamster(t *testing.T) {
	animalQuantity := 10

	result := animalFood.AnimalHamster(animalQuantity)

	expectedResult := 2.5

	//Con la libreria testing de go
	// if result != expectedResult {
	// 	t.Fatal("La comida esperada para el Hamster era ", expectedResult, " y nos devolvio ", result)
	// }

	//Testify
	require.Equal(t, result, expectedResult)
}

func TestTarantula(t *testing.T) {
	animalQuantity := 10

	result := animalFood.AnimalTarantula(animalQuantity)

	expectedResult := 1.5

	//Con la libreria testing go
	// if result != expectedResult {
	// 	t.Fatal("La comida esperada para el Tarantula era ", expectedResult, " y nos devolvio ", result)
	// }

	//Con testify
	require.Equal(t, result, expectedResult)
}
