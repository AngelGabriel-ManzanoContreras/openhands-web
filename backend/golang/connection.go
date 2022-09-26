package main

import (
	"database/sql"
	_ "github/lib/pq"
	"log"
)

// getConnection obtiene una conexion a la BD
func getConnection() *sql.DB {
	//debes crear una base de datos
	//dsn := "postgres://<usuario>:<contraseña>@<servidor>:<puerto (el puerto habitual de postgres es "5432")>/<base de datos>?ssl=disable"
	//conectarse a la basede datos
	db, err := sql.Open("postgres", dsn) //db por la conexion que devolvera y err por si hay un error
	//si hay un error nulo, regresa un log fatal y el tipo de error que es
	if err != nil {
		log.Fatal(err)
	}

	//regresa la conexion a la base de datos
	return db
}
