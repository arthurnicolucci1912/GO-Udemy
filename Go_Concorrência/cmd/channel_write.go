package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//Iniciando uma goroutine que envia números para op canal
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Enviando numero %d para o canal\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	//lendo numeros
	for num := range ch {
		fmt.Printf("Número recebido %d", num)
	}
}
