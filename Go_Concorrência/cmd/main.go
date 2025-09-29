package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	var price1, price2, price3 float64
	var wg sync.WaitGroup //grupo de espera
	wg.Add(3)             //3 goroutineS

	go func() {
		defer wg.Done() //executa quando função foi executada
		price1 = fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done() //executa quando função foi executada
		price2 = fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done() //executa quando função foi executada
		price3 = fetcher.FetchPriceFromSite3()
	}()

	wg.Wait() //aguarda terminarem

	fmt.Printf("RS %.2f\n", price1)
	fmt.Printf("RS %.2f\n", price2)
	fmt.Printf("RS %.2f\n", price3)
	fmt.Printf("Tempo total: %s", time.Since(start))

}
