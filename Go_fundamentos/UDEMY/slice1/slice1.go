package main

import "fmt"

func main() {
	var numeros []int

	numeros = append(numeros, 10)
	numeros = append(numeros, 20, 30, 40)
	numeros = append(numeros, 10)
	numeros = append(numeros, 50)

	fmt.Println("Slice de numeros: ", numeros)
}
