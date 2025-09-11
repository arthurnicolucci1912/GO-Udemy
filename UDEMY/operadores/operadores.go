package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("n1: ")
	fmt.Scan(&n1)

	fmt.Print("n2: ")
	fmt.Scan(&n2)

	//:= serve para variáveis locais
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := n1 / n2
	mod := n1 % n2

	fmt.Printf("Soma de %d e %d: %d\n", n1, n2, soma)
	fmt.Printf("Sub de %d e %d: %d\n", n1, n2, sub)
	fmt.Printf("Multiplicação de %d e %d: %d\n", n1, n2, mult)
	fmt.Printf("Divisão de %d e %d: %d\n", n1, n2, div)
	fmt.Printf("Resto de %d e %d: %d \n", n1, n2, mod)

}
