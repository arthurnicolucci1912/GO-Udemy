package main

import "fmt"

//Funcao que recebe ponteiro como argumento e altera o valor da variavel original
func alterValue(x *int) {
	*x = *x * 2
}

func main() {
	num := 5
	fmt.Printf("Valor Inicial: %d\n", num)

	//alteração
	alterValue(&num)

	fmt.Printf("Valor alterado: %d\n", num)
}
