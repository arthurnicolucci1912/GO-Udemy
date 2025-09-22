package main

import "fmt"

//1- imprimir mensagem de boas vindas
func welcome() {
	fmt.Println("Seja bem-vindo ao cinema!")
}

//cadastrar um dilme
func createMovie() {
	var movieName string
	var yearRelease int
	var moviePrice float64
	var dummy string

	fmt.Println("Nome do filme: ")
	fmt.Scanln(&movieName)
	//limpar buffer de memória
	fmt.Scanln(&dummy)

	fmt.Println("Ano do filme: ")
	fmt.Scanln(&yearRelease)

	fmt.Println("Preço do filme: ")
	fmt.Scanln(&moviePrice)

	fmt.Printf("%s (%d) - %.2f\n ", movieName, yearRelease, moviePrice)
}

//funcao para calcular a media de notas
func calculateAverage() float64 {
	var numRaitings int
	fmt.Println("Quantas avaliações para o filme: ")
	fmt.Scan(&numRaitings)

	var total float64
	for i := 0; i < numRaitings; i++ {
		var note float64
		fmt.Println("Digite a nota do filme:")
		fmt.Scan(&note)
		total += note
	}
	var avarage float64
	if numRaitings > 0 {
		avarage = total / float64(numRaitings)
	} else {
		avarage = 0
	}
	return avarage
}

func main() {
	fmt.Println("Utilizando função")
	welcome()
	// createMovie()
	media := calculateAverage()
	fmt.Printf("A media das avaliações é: %2.f \n", media)

}
