package main

import "fmt"

func main() {

	var salario float32 = 200000

	imp := impuesto(salario)
	fmt.Println(imp)

}

func impuesto(salario float32) float32 { 
		
	var imp float32
	if salario > 50000 {
		
		if salario < 150000 {
			
			imp = salario * 0.17
			
			return imp
		} else {
			imp = salario * 0.27

			return imp
		}
	} 
	return 0
}