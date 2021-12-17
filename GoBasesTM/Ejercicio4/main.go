package main

import ("fmt"
	"math"
	"errors"
)

 func max(values ...int) int {
	var max int = int(math.MinInt)
	for _, v := range values {
		if v > max {
			max = v
		}
	}
	return max
 }

 func min(values ...int) int {
	var min int = int(math.MaxInt)
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
 }

 func prom(values ...int) int {
	 c := len(values)
	 sum := 0
	 for _, v := range values {
		sum += v
	 }
	 return sum/c
 }

 func operacion(tipo string) (func(values ...int) int, error) {
	switch tipo {
		case "minimo":
			return min, nil
		case "maximo":
			return max, nil
		case "promedio":
			return prom, nil
		default:
			return nil, errors.New("invalid operation")
	}
 }

 func main() {

	const (
		minimo = "minimo"
		promedio = "promedio"
		maximo = "maximo"
	 )
	  
	 // minimo
	  
	 minFunc, err := operacion(minimo)
	 if err != nil {
		 fmt.Println(err)
	 } else {
		valorMinimo := minFunc(2,3,3,4,1,2,4,5)
		fmt.Printf("El valor minimo es %d \n", valorMinimo)
	 }

	 // promedio

	 promFunc, err := operacion(promedio)
	 if err != nil {
		fmt.Println(err)
	} else {
		valorPromedio := promFunc(2,3,3,4,1,2,4,5)
	   fmt.Printf("El valor promedio es %d \n", valorPromedio)
	}

	// maximo


	 maxFunc, err := operacion(maximo)
	 if err != nil {
		fmt.Println(err)
	} else {
	   valorMaximo := maxFunc(2,3,3,4,1,2,4,5)
	   fmt.Printf("El valor maximo es %d \n", valorMaximo)
	}

 }