package main

import "fmt"

func main() {
	var idade int = 18
	var altura float64 = 1.87
	var maiordidade = idade >= 18
	fmt.Println("--Dadaos pessoais em GO--")
	fmt.Println("Idade:", idade)
	fmt.Println("Altura:", altura)
	fmt.Println("Maior de idade:", maiordidade)
}
