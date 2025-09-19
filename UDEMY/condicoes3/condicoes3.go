package main

import "fmt"

func main() {
	var numero int
	fmt.Print("Digite um Numero: ")
	fmt.Scan(&numero)

	if numero > 0 {
		fmt.Print("O numero e positivo")
	} else if numero < 0 {
		fmt.Print("O numero e negativo")
	} else {
		fmt.Print("O numero e zero")
	}

	if numero%2 == 0 {
		fmt.Print(" e par")
	} else {
		fmt.Print(" e impar")
	}

}
