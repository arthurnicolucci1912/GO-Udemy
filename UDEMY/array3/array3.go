package main

import "fmt"

func main() {
	//declarando um array de 2 strings com valores iniciais
	var notas = [5]float64{4.5, 3.6, 8.7, 9.0}

	soma := notas[0] + notas[1] + notas[2] + notas[3]

	fmt.Println("Notas", notas)
	fmt.Println("Soma: ", soma)
}
