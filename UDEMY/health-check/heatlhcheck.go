package main

import (
	//para transformar a resposta em JSON
	"encoding/json"
	//foramt para prints e scans
	"fmt"
	//para usar metodos de http
	"net/http"
	//lib de tempo
	"time"
)

// tempo de execução
var tempoAtual = time.Now()

func main() {
	http.HandleFunc("/healthcheck", func(response http.ResponseWriter, request *http.Request) {

		// Configurar header para JSON
		response.Header().Set("Content-Type", "application/json")

		// Criar struct com os dados
		data := map[string]interface{}{
			"data_hora_servidor": time.Now().Format("02/01/2006 15:04:05"),
			"tempo_rodando":      time.Since(tempoAtual).Round(time.Second).String(),
			"mensagem":           "OK",
		}

		// Converter para JSON e enviar
		json.NewEncoder(response).Encode(data)
	})

	http.HandleFunc("/homem-poste", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "OK:)")
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
