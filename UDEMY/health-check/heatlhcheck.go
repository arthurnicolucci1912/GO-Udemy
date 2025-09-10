package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/teste", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "DEU CERTO!!!")
	})

	fmt.Println("Starting server on :8080")
	http.ListenAndServe(":8080", nil)
}
