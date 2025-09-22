package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	URL := "https://jsonplaceholder.typicode.com/todos/2"
	response, err := http.Get(URL)
	if err != nil {
		fmt.Println("erro ao fazer a requisição!")
		return
	}

	//indica que só vai ser executado ao termino da funcção
	defer response.Body.Close()

	//usando o Io.Readall para ler o corpo da resposta
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("erro ao ler o corpo da resposta!")
		return
	}

	fmt.Println("Resposta da API: ", string(body))
}
