package main

import (
	"fmt"
)

func fullName(firstName, lastName string) {
	fmt.Printf("Nome completo: %s %s\n", firstName, lastName)
}

// tipo do param | tipo do retorno
func sumNumbers(a, b float64) float64 {
	return a + b
}

func adress(country string) {
	if country == "" {
		country = "Brasil"
	}
	fmt.Printf("Eu moro no: %s\n", country)
}

func main() {
	fmt.Println("Utilizando função com parametros")
	fullName("arthur", "nicolucci")
	fmt.Printf("Soma: %.2f\n", sumNumbers(1, 2))
	adress("")
	adress("Canadá")
}
