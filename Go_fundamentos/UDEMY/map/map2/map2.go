package main

import (
	"fmt"
	"strings"
)

func main() {
	//map para contar palavras
	text := "go é uma linguagem muito rapida!"
	words := strings.Fields(text)
	wordsCount := make(map[string]int)

	//contagem da frequesncia de palavras
	//index null, para palavra dentro de palavras = dentro do contador de palvras iterar 1 valor = palavra = palavra+1
	for _, word := range words {
		wordsCount[word]++
	}

	//imprimir a frequencia de palavras
	fmt.Println("Contagem de palavras")

	for word, count := range wordsCount {
		fmt.Printf("Palavra: %s, Frequencia: %d \n", word, count)
	}
}
