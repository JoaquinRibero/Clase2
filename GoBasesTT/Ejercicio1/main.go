package main

import (
	"fmt"
	"time"
	"encoding/json"
)

type Persona struct {
	Nombre    string    `json:"nombre"`
	Apellido  string    `json:"apellido"`
	Dni       int       `json:"dni"`
	Fecha     time.Time `json:"fecha"`
}

func (p Persona) detalle() string {
	miJSON, err := json.Marshal(p)
	if err != nil {
		return "Error marshalling Persona"
	} else {
		return (string(miJSON))
	}
}

func main() {
	p := Persona{
		Nombre: "Joaquin",
		Apellido: "Ribero",
		Dni: 39500508,
		Fecha: time.Date(1996, time.March,14,0,0,0,0, time.UTC),
	}

	fmt.Println(p.detalle())
}