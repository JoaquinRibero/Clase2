package main

import "fmt"

type producto struct {
	tipo string    `json:"tipo"`
	nombre string `json:"nombre"`
	precio int `json:"precio"`
}

type Producto interface {
	CalcularCosto() float64
}

func (p producto) CalcularCosto() float64 {
	switch p.tipo {
		case "pequeño":
			return float64(p.precio)
		case "mediano":
			return float64(p.precio) * 1.03
		case "grande":
			return float64(p.precio) * 1.06 + 2500.00
		default:
			return 0
	}
}

type tienda struct {
	products []Producto
}

type Ecommerce interface {
	Total() float64
	Agregar(Producto) string
}

func nuevoProducto(tipo string, nombre string, precio int) Producto {
	var prod producto
	prod.tipo = tipo
	prod.nombre = nombre
	prod.precio = precio

	var p Producto
	p = prod

	return p
}

func nuevaTienda (products[] Producto) Ecommerce{
	var t tienda
	t.products = products

	var ec Ecommerce
	ec = &t

	return ec
}

func (t *tienda) Total() float64 {
	sum := 0.0
	for _,p := range t.products {
		sum += p.CalcularCosto()
	}
	return sum
}

func (t *tienda) Agregar(p Producto) string {
	t.products = append(t.products, p)

	return "Producto agregado correctamente"
}

func main() {
	p1 := nuevoProducto("pequeño", "Celular", 5000)
	p2 := nuevoProducto("mediano", "Notebook", 8000)
	p3 := nuevoProducto("grande", "PC", 10000)

	var prods [] Producto = []Producto{p1, p2, p3}

	t := nuevaTienda(prods)
	fmt.Println(t.Total())

	p4 := nuevoProducto("pequeño", "Tablet", 3660)
	t.Agregar(p4)
	fmt.Println(t.Total())
}