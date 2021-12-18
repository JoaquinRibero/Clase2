package main

import "fmt"
import "math"

type Matrix struct {
	Data [][]int 
	Height int
	Width int
	Cuadratic bool
	MaxValue int
}

func (m *Matrix)setMatrix(data [][]int) {
	var max int = int(math.MinInt)
	m.Data = data
	m.Height = len(data)
	m.Width = len(data[0])
	if len(data) == len(data[0]) {
		m.Cuadratic = true
	} else {
		m.Cuadratic = false
	}
	for i := range data {
		for j := range data[i] {
			if data[i][j] > max {
				max = data[i][j]
			}
		}
	}
	m.MaxValue = max

}

func (m Matrix) printMatrix() {
	for i := range m.Data {
		fmt.Printf("%d\n", m.Data[i])
	}
}

func main() {

	var m Matrix
	data := [][]int{
		{1,2,3,4},
		{4,5,8,6},
	}
	m.setMatrix(data)
	m.printMatrix()
	fmt.Printf("El valor màximo de la matriz es: %d\n", m.MaxValue)
	fmt.Printf("La matriz es cuadratica: %t\n", m.Cuadratic)
}