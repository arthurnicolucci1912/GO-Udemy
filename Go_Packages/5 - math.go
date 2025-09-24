package main

import (
	"fmt"
	"math"
)

// 1- Acessar PI
func acessarPI() {
	fmt.Printf("PI arredondado: %.2f\n", math.Pi)
}

// 2 Acessar PI
func acessarEuler() {
	fmt.Printf("Euler: %.2f\n", math.E)
}

// 3 - arredondando numeros ><
func arredondar() {
	num := 10.4
	fmt.Println("Arredondando p cima: ", math.Ceil(num))
	fmt.Println("Arredondando p baixo: ", math.Floor(num))
}

func main() {
	acessarPI()
	acessarEuler()
	arredondar()
	fmt.Println("Potencia de 5^5: ", math.Pow(5, 5))
	fmt.Println("Raiz quadrada de 25: ", math.Sqrt(25))
	fmt.Printf("Logaritmo de 10: %.2f", math.Log(10))
}
