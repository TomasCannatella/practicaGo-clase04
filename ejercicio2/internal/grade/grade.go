package grade

func AvgGrade(grades ...int) (avarage float64) {
	var totalGrades int
	for _, grade := range grades {
		if grade < 0 {
			continue
		}
		totalGrades += grade
	}
	avarage = float64(totalGrades) / float64(len(grades))
	return
}
