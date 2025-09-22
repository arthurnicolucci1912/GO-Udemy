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

	//faça
	line := "="

	fmt.Print(strings.Repeat(line, 30))
	fmt.Println(movieDescription)
	fmt.Print(strings.Repeat(line, 30))

	//verifica se a palavra existe dentro da String
	fmt.Println(strings.Contains(movieDescription, "top"))
	fmt.Println(strings.Contains(movieName, "Top"))

}
