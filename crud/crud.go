package main

import (
	"fmt"
	"log"
)

func main() {
	e := Estudiante{
		Name:   "Almendra",
		Age:    23,
		Active: true,
	}

	err := Crear(e)
	if err != nil {
		log.Fatal(err)
	}

	e2 := Estudiante{
		Name:   "Chris",
		Age:    27,
		Active: true,
	}

	err = Crear(e2)
	if err != nil {
		log.Fatal(err)
	}

	e3 := Estudiante{
		Name:   "José",
		Age:    28,
		Active: false,
	}

	err = Crear(e3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Creado exitosamente!!")
}
