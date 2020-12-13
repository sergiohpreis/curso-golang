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

	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")

	stmt.Exec("Uóxiton Istive", 1)
	stmt.Exec("Sheristone Ualeska", 2)

	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
	stmt2.Exec(4)
}
