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

	fmt.Println("Creado exitosamente!")
}
