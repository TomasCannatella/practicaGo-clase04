/*
	Ejercicio 3 - Calcular salario

Una empresa marinera necesita calcular el salario de sus empleados
basándose en la cantidad de horas trabajadas por mes y la categoría.

Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de
minutos trabajados por mes, la categoría y que devuelva su salario.
*/
package main

import (
	"ejercicio3/internal/salary"
	"fmt"
)

func main() {
	workMinutes := 120
	category := "A"

	fmt.Println("El salario del trabajador es: ", salary.CalculateSalary(workMinutes, category))
}
