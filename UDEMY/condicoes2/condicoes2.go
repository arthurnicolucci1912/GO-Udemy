package main

import "fmt"

func main() {
	var nota float64
	fmt.Println("Digite sua nota: ")
	fmt.Scan(&nota)

	if nota < 0 || nota > 10 {
		fmt.Println("nota invalida, digite uma nota entre 0 e 10")
	} else if nota >= 9 {
		fmt.Println("Execelente")
	} else if nota >= 7 {
		fmt.Println("Muito bom")
	} else if nota >= 5 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado!")
	}
}
