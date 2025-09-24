package main

import (
	"fmt"
	"os"
)

func main() {
	//criando um arquivo
	arquivo, err := os.Create("exemploOS.txt")
	if err != nil {
		fmt.Println("erro ao criar o arquivo")
		return
	}

	defer arquivo.Close()

	//escrevendo no arquivo
	_, err = arquivo.WriteString("---\n -->Esse arquivo de texto foi criando usando funções do package OS<-- \n---")
	if err != nil {
		fmt.Println("erro ao escrever no arquivo: ", err)
		return
	}
	fmt.Println("Arquivo criado e texto escrito com sucesso!")

	//lendo o arquivo
	conteudo, err := os.ReadFile("exemploOS.txt")
	if err != nil {
		fmt.Println("erro ao ler arquivo: ", err)
		return
	}

	fmt.Println("Conteúdo do arquivo: \n", string(conteudo))
}
