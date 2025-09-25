package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// faz a conexão do banco
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	//criando instução
	result, err := db.Exec("UPDATE users SET name = ? WHERE id = ?", "Churiripo", 5)
	if err != nil {
		log.Fatal(err)
	}

	affectedRows, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Linhas Atualizadas: %d\n", affectedRows)

	//exclusao de dados
	result, err = db.Exec("DELETE FROM users WHERE id = ?", 1)
	if err != nil {
		log.Fatal(err)
	}
	affectedRows, err = result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Linhas Apagadas: %d", affectedRows)
}
