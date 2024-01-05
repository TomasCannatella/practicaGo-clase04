package grade

import "slices"

func MinFunc(grades ...int) float64 {
	return float64(slices.Min(grades))
}

func AverageFunc(grades ...int) (avarage float64) {
	var totalGrades float64
	for _, grade := range grades {
		totalGrades += float64(grade)
	}
	avarage = totalGrades / float64(len(grades))
	return
}

func MaxFunc(grades ...int) float64 {
	return float64(slices.Max(grades))
}
