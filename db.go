package main

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func iniDB() *sqlx.DB {
	conn, err := sqlx.Connect(driver, nomeArquivo)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := conn.Ping(); err != nil {
		log.Fatal(err.Error())
	}

	return conn
}
