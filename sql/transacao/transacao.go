package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:password@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, nome) values (?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")
	_, execErr := stmt.Exec(1, "Tiago") // chave duplicada

	if execErr != nil {
		tx.Rollback()
		log.Fatal(err)
	}
	tx.Commit()
}
