package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//da mesma forma, faz a conexão do banco

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//criando instução
	//Passando os dados diretamente na cláusula
	_, err = db.Exec("INSERT INTO users (name) VALUES(?)", "Arthur Nicolucci")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Novo usuário inserido com sucesso!")

}
