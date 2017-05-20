package main

import (
	"errors"
	"fmt"
	"time"
)

// Estudiante estructura de la bla en la DB
type Estudiante struct {
	ID        int
	Name      string
	Age       int16
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Crear resgistro de un Estudiante en DB
func Crear(e Estudiante) error {
	q := `INSERT INTO colegio.estudiantes(name, age,active) VALUES($1, $2, $3)`

	db := getConnection()
	defer db.Close()
	stmt, err := db.Prepare(q)
	if err != nil {
		return err
	}
	defer stmt.Close()

	r, err := stmt.Exec(e.Name, e.Age, e.Active)
	if err != nil {
		return err
	}

	i, _ := r.RowsAffected()
	fmt.Println(i, " filas afectadas.")

	if i != 1 {
		return errors.New("Error: se esperaba 1 fila afectada")
	}

	return nil
}
