package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID        int    `db:"id"`
	Nome      string `db:"nome"`
	Sobrenome string `db:"sobrenome"`
	Criacao   string `db:"criacao"`
}

func (u *User) CreateTableuser(db *sqlx.DB) error {
	createTbUsers := `create table if not exists users (
		"id" integer not null primary key autoincrement,
		"nome" text,
		"sobrenome" text,
		criacao datetime default current_timestamp
		)`
	stm, err := db.Prepare(createTbUsers)
	if err != nil {
		return fmt.Errorf("Falha ao criar tabela de usuários: Erro - %s", err.Error())
	}

	_, err = stm.Exec()
	if err != nil {
		return fmt.Errorf("Falha ao criar tabela de usuários: Erro - %s", err.Error())
	}

	return nil
}

func (u *User) InsertUser(db *sqlx.DB) error {
	insertSQL := `
		insert into users (
			nome, sobrenome
		) values (
			?, ?
		)
	`

	stm, err := db.Prepare(insertSQL)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Falha ao criar usuário: Erro - %s", err.Error())
	}

	_, err = stm.Exec(u.Nome, u.Sobrenome)
	if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("Falha ao criar usuário: Erro - %s", err.Error())
	}

	return nil
}

func (u *User) ShowAllUsers(db *sqlx.DB) error {
	var users []User
	err := db.Select(&users, "SELECT * FROM users ORDER BY nome")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	for _, v := range users {
		fmt.Printf("ID: %d - Nome: %s - Sobrenome: %s - Inclusão: %s \n", v.ID, v.Nome, v.Sobrenome, v.Criacao)
	}

	return nil
}
