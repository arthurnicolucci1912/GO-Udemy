package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(10) + 1 //número aleatório entre 1 e 10
	fmt.Println("Jogo adivinha")
	fmt.Println("Tente adivinhar um número entre 1 e 10")
	var guess int
	for {

		fmt.Println("Digite seu palpite: ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Tente um número maior")
		} else if guess > target {
			fmt.Println("Tente um número menor")
		} else {
			fmt.Println("Parabéns! Você acertou o número:", target)
			break
		}

		if guess == 0 {
			fmt.Println("Jogo encerrado, o numero era: ", target)
			break
		}

	}
}
