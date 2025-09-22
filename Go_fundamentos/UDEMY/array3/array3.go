package main

import "fmt"

func main() {
	//declarando um array de 2 strings com valores iniciais
	var notas = [5]float64{6, 6.6, 8.7, 9.0}

	soma := notas[0] + notas[1] + notas[2] + notas[3]

	fmt.Println("Array de Notas: ", notas)
	fmt.Printf("Soma: %.2f\n", soma)

	media := soma / float64(len(notas))

	fmt.Printf("Média: %.2f\n", media)
	fmt.Printf("Total de Notas: %d\n", len(notas))
}
