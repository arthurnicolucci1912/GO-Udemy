//go get github.com/google/uuid

package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// Struct para modelar, ditar o que vamos precisar para o CRUD de dados
type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

// Inserir novo user
func newUser(name string, email string, age int) *User {
	return &User{
		ID:    uuid.New().String(),
		Name:  name,
		Email: email,
		Age:   age,
	}
}

// Criando tabela
func createTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users(
		id TEXT PRIMARY KEY,
		name TEXT,
		email TEXT,
		age integer
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func insertUser(db *sql.DB, user *User) error {
	//melhor prática de sql injection
	stmt, err := db.Prepare("INSERT INTO users (id, name, email, age) VALUES($1, $2, $3, $4)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}

	return nil

}

func updateUSer(db *sql.DB, user *User) error {
	//melhor prática de sql injection
	stmt, err := db.Prepare("UPDATE users SET name = ?, email = ?, age = ? WHERE id = ? ")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.Age, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func selectUser(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)

	}
	return users, nil

}

func deleteUser(db *sql.DB, id string) error {
	//melhor prática de sql injection
	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	//passa o id como parâmetro
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	//criaando novo banco de dados
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//criar tabela
	err = createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	//inserindo user
	//Instancia os dados da struct User
	user := newUser("Arthur", "arthur@gmail.com", 18)
	//, e ja usa a função inserUser
	err = insertUser(db, user)
	if err != nil {
		log.Fatal(err)
	}

	//atualizando User
	user.Name = "Marcio Antonio"
	user.Email = "marcio@gmail.com"
	user.Age = 46
	err = updateUSer(db, user)
	if err != nil {
		log.Fatal(err)
	}

	//listagem
	users, err := selectUser(db)
	if err != nil {
		log.Fatal(err)
	}

	for _, u := range users {
		fmt.Printf("User -> %s possui email %s e idade %d\n", u.Name, u.Email, u.Age)
	}

	//delete
	err = deleteUser(db, user.ID)
	if err != nil {
		log.Fatal(err)
	}
}
