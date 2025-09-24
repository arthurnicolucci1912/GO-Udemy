package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Convertendo string para int
	numeroStr := "123"
	numero, err := strconv.Atoi(numeroStr)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Numero: ", numero)

	//convertendo int para string
	num := 10
	numstring := strconv.Itoa(num)
	fmt.Println(numstring)

	//converter string para float
	floatStr := "10.59"
	valorFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("Erro: ", err)
		return
	}
	fmt.Println("Float: ", valorFloat)
}
