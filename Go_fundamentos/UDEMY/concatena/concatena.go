package main

import "fmt"

func main() {
	var nome string
	var idade int
	var altura float64
	var maioridade bool
	var maior string

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Print("Digite a sua idade: ")
	fmt.Scan(&idade)

	maioridade = idade >= 18

	fmt.Print("Digite sua altura: ")
	fmt.Scan(&altura)

	fmt.Print("\nDados Pessoais\n")
	fmt.Printf("Nome: %s\n", nome)
	fmt.Printf("Idade: %d\n", idade)
	fmt.Printf("Altura: %.2f\n", altura)

	if maioridade == true {
		maior = "Maior de Idade"
		fmt.Printf("Maior Idade: %s\n", maior)
	}
	fmt.Printf("Menor Idade: %v\n", maioridade)

}
