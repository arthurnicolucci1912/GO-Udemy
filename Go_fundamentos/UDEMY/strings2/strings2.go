package main

import (
	"fmt"
	"strings"
)

func main() {
	movieName := "Top Gun"
	movieName2 := "top Gun"

	//go é case sensitive
	fmt.Println(movieName == movieName2)

	movieDescription := `
Top gun é um filme 
muito bom onde o piloto maverick é o piloto
de um jato muito rapido kkk`

	//ao inves de
	// fmt.Println("===================")

	//convertendo para maiusculo
	fmt.Println(strings.ToUpper(movieDescription))
	fmt.Println((strings.ToLower(movieDescription)))

	//primeira letra maiuscula
	fmt.Println(strings.ToTitle(movieDescription))

	//posição do caractere
	fmt.Println(strings.Index(movieName, "G"))

	//quantas vezes aparece aquele caractere
	fmt.Println(strings.Count(movieDescription, "a"))

	//substituir elemento
	fmt.Println(strings.ReplaceAll(movieDescription, "filme", "serie"))

	//dividindo a string
	fmt.Println(strings.Split(movieDescription, "="))

}
