package main

import (
	"fmt"
	"strings"
)

func main() {
	var nome string
	var idade int
	var altura float64
	var maioridade bool

	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)
	fmt.Print("Digite a sua idade: ")
	fmt.Scan(&idade)
	fmt.Print("Digite sua altura: ")
	fmt.Scan(&altura)

	maioridade = idade >= 18
	var sb strings.Builder

	fmt.Print("\nDados Pessoais\n")

	sb.WriteString(fmt.Sprintf("Nome: %s \n", nome))
	sb.WriteString(fmt.Sprintf("Idade: %d\n", idade))
	sb.WriteString(fmt.Sprintf("Altura: %.2f\n", altura))
	sb.WriteString(fmt.Sprintf("Maior Idade: %v\n", maioridade))

	//agora vamos printar o resultado da concatenação de strings feita pelo método strings.Builder
	fmt.Println(sb.String())
}
