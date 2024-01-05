package salary

import "errors"

var categoriesSalary = map[string]int{
	"C": 1000,
	"B": 1800,
	"A": 4500,
}

func CalculateSalary(workMinutes int, category string) (salary float64, err error) {
	salaryPerHour, ok := categoriesSalary[category]

	if !ok {
		err = errors.New("this category not exist")
		return
	}
	salary = (float64(workMinutes) / 60.0) * (float64(salaryPerHour))
	return
}
