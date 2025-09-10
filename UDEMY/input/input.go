package main

import "fmt"

func main() {
	var nome string
	var idade int
	var cpf int

	fmt.Print("Nome:")
	fmt.Scan(&nome)

	fmt.Print("Idade: ")
	fmt.Scan(&idade)

	fmt.Print("CPF: ")
	fmt.Scan(&cpf)

	fmt.Print("Bem vindo ", nome, "\nVocê tem ", idade, " anos, e seu CPF:", cpf)
}
