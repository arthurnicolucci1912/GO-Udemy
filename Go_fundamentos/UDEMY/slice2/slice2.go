package main

import (
	"fmt"
	"strings"
)

func main() {
	frutas := []string{"banana", "maçã", "mamao", "Abacaxi", "uva"}

	//pegando um subslice do indice de 1 a 3
	subslice := frutas[1:4]
	fmt.Println("slice: ", frutas)

	linha := "-"
	fmt.Println(strings.Repeat(linha, 50))

	fmt.Println("subslice de frutas: ", subslice)

	subslice[0] = "Manga"
	fmt.Print(frutas)

}
