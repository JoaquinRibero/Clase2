package main

import "fmt"
import "errors"

func promedio(calificaciones ...int) (error,int) {

	total := len(calificaciones)
	sum := 0
	for _, c := range calificaciones {
		
		if c < 0 {
			return errors.New("Uno de los valores es negativo"), 0
		}
		sum += c
	}
	promedio := sum / total
	return nil, promedio
}

func main() {

	err, prom := promedio(4,5,7,6,8,9,1,5,8,6,-1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(prom)
	}

}