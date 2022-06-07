package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SQLconnect() {
	database, err := sql.Open("mysql", "root:test@tcp(127.0.0.1:3306)/karaz_platform")

	if err != nil {
		panic(err.Error())
	}

	DB = database

}
