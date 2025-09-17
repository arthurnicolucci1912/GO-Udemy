package main

import "fmt"

func main() {
	//declarando um array de 5 numeros e atribuindo valores a cada posição
	var numeros [5]int

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30
	numeros[3] = 40
	numeros[4] = 50

	fmt.Println("Array de numeros em GO", numeros)
}
