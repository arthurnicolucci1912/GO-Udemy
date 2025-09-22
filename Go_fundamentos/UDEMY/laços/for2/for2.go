package main

import "fmt"

func main() {
	notas := []float64{7.5, 8.0, 6.5, 9.0, 5.5}
	fmt.Println("Calculando media com slice de notas: ", notas)
	var total float64
	for id, nota := range notas {
		fmt.Println("Somando nota de id: ", id, " - ", total, "+", nota)
		total += nota
	}

	//Aqui temos que converter len(notas) que é int para float64
	//para fazer a divisão corretamente e dividir a soma total, pela quantidade de elementos que estao no slice

	media := total / float64(len(notas))

	fmt.Println("Soma: ", total)
	fmt.Println("Média: ", media)
}
