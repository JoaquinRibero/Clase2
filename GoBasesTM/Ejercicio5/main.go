package main

import "fmt"
import "errors"

func animalPerro(cant int) float32 {
	return float32(cant * 10)
}
func animalGato(cant int) float32 {
	return float32(cant * 5)
}
func animalHamster(cant int) float32 {
	return float32(cant) * 0.250
}
func animalTarantula(cant int) float32 {
	return float32(cant) * 0.150
}

func animal(tipo string) (func(cant int)float32, error) {

	switch tipo {
		case "Perro":
			return animalPerro, nil
		case "Gato":
			return animalGato, nil
		case "Hamster":
			return animalHamster, nil
		case "Tarantula":
			return animalTarantula, nil
		default:
			return nil, errors.New("Animal no existente")
	}

}
func main() {

	const (
		perro = "Perro"
		gato = "Gato"
		hamster = "Hamster"
		tarantula = "Tarantula"
	)
	var cantidad float32

	animalPerro, err := animal(perro)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalPerro(5)
	}

	animalGato, err := animal(gato)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalGato(8)
	}

	animalHamster, err := animal(hamster)
	if err != nil {
		fmt.Println(err)
	} else {
		cantidad += animalHamster(3)
	}

	fmt.Printf("La cantidad de comida necesaria es: %.3f", cantidad)

}