package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
}

func atualizarCliente(c *Cliente, novoNome string, novaIdade int) {
	//passar "." para ter acesso a informações da struct
	c.Nome = novoNome
	c.Idade = novaIdade
}

func main() {
	cliente := Cliente{Nome: "Arthur", Idade: 30}
	fmt.Println("Antes da alteração: ", cliente)

	//atualizando
	atualizarCliente(&cliente, "Lucas", 18)

	fmt.Println("Após atualização: ", cliente)

}
