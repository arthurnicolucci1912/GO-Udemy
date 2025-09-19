package main

import "fmt"

func main() {
	var idade int
	fmt.Println("Informe sua idade: ")
	fmt.Scan(&idade)

	if idade >= 18 {
		fmt.Println("Você é maior de idade!")
	} else {
		fmt.Println("Você é menor de idade!")
	}
}
