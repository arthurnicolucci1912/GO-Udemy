package main

import "fmt"

//definindo uma struct
type Pessoa struct {
	Nome     string
	Idade    int
	Endereco string
}

func main() {

	//criar a instancia da struct carro
	carro1 := Pessoa{
		Nome:     "Arthur",
		Idade:    18,
		Endereco: "sumare",
	}

	fmt.Println("Info da Pessoa")
	fmt.Printf("Nome: %s | Idade: %d | Endereco: %s", carro1.Nome, carro1.Idade, carro1.Endereco)

}
