package main

import (
	"fmt"
	"net/http"
)

// responsável pela inicialização do servidor interno GO
// função que processa requisições HTTP
// *http.Request é o ponteiro para o objeto que contém as informações da requisição

func handler(response http.ResponseWriter, request *http.Request) {
	//escreve resposta já formatada diretamente no cliente
	// request.URL.Path = caminho de URL que foi solicitado
	fmt.Fprintf(response, "Ola você acessou o endpoint: %s", request.URL.Path)
}

func main() {

	//registra a função handler para todas as requisições no caminho raiz: "/"
	http.HandleFunc("/", handler)
	fmt.Println("O servidor está rodando em http://localhost:8080")

	//define que o servidor "escuta" na porta :8080
	//nil define que usaremos o roteador padrão
	http.ListenAndServe(":8080", nil)
}
