package main

import "fmt"

func main() {

	sal := salario(360, "A")
	fmt.Printf("El salario es %.2 \n", sal)
}

func salario(minutos float32, categoria string) float32 {

	var horas float32 = minutos / 60
	var salario  float32 
	switch categoria {
		case "C":
			salario = horas * 1000
			return salario
		case "B":
			salario = (horas * 1500) * 1.2
			return salario
		case "A":
			salario = (horas * 3000) * 1.5
			return salario
		default:
			return salario
	}
}