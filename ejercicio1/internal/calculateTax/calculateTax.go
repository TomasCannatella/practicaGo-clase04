package calculateTax

func CalculateTax(salary float64) (tax float64) {
	switch {
	case salary > 150000:
		tax += (salary * 10) / 100
		fallthrough
	case salary > 50000:
		tax += (salary * 17) / 100
	default:
		tax = 0
	}
	return
}
