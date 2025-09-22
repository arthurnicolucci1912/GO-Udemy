package main

import (
	"fmt"
)

// fucao para numeros variadicos
func sum(nums ...int) {
	sumTotal := 0
	for _, n := range nums {
		sumTotal += n
	}
	fmt.Printf("Soma é: %d\n", sumTotal)
}

// funcao para apresentação de cursos com variadicos
func presentation(data map[string]string) {
	for key, value := range data {
		fmt.Printf("%s - %s\n", key, value)
	}
}

func main() {
	sum(7)
	//pode chamar quantos numeros quiser
	sum(7, 10, 5)
	presentation(map[string]string{
		"Name":     "C#",
		"Category": "Backend",
		"Level":    "Avançado",
	})
	presentation(map[string]string{
		"Name":     "Python",
		"Category": "Automação",
		"Level":    "Iniciante",
	})
}
