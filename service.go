package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB
var driver string = "sqlite3"
var nomeArquivo string = "go_sqlite3.db"

func main() {
	os.Remove(nomeArquivo)

	f, err := os.Create(nomeArquivo)
	if err != nil {
		fmt.Printf("Falha ao gerar arquivo %s: %s \n", nomeArquivo, err.Error())
	}
	defer f.Close()

	db = iniDB()
	defer db.Close()

	u := new(User)
	err = u.CreateTableuser(db)
	if err != nil {
		log.Fatal(err.Error())
	}

	var count int
	for count < 5 {
		v := new(User)
		v.Nome = fmt.Sprintf("User-%d", count+1)
		v.Sobrenome = "Test Insert"
		v.InsertUser(db)
		if err != nil {
			log.Fatal(err)
		}
		count++
	}

	u.ShowAllUsers(db)
	if err != nil {
		log.Fatal(err)
	}
}
