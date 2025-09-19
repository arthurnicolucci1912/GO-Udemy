package main

import "fmt"

func main() {
	var nota int
	fmt.Print("Digite a nota do aluno de 0 a 10: ")
	fmt.Scan(&nota)

	switch {
	case nota >= 9:
		fmt.Print("Classificação A: Excelente!")
	case nota >= 7:
		fmt.Print("Classificação B: Bom!")
	case nota >= 5:
		fmt.Print("Classificação C: Regular!")
	case nota >= 3:
		fmt.Print("Classificação D: Ruim!")
	case nota >= 0:
		fmt.Print("Classificação E: Pessimo!")
	default:
		fmt.Print("Nota invalida!")
	}
}
