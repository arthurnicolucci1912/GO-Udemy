package main

import (
	"buscador/internal/fetcher"
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	priceChannel := make(chan float64)
	var wg sync.WaitGroup //grupo de espera
	wg.Add(3)             //3 goroutineS

	//canal de comunicação das routines
	go func() {
		var totalPrices float64
		countPrice := 0.0
		for price := range priceChannel {
			totalPrices += price
			countPrice++
			fmt.Printf("R$ %.2f\n", price)
			fmt.Printf("Preço médio: R$ %.2f \n", totalPrices/countPrice)
		}
	}()
	go func() {
		defer wg.Done() //executa quando função foi executada
		priceChannel <- fetcher.FetchPriceFromSite1()
	}()

	go func() {
		defer wg.Done() //executa quando função foi executada
		priceChannel <- fetcher.FetchPriceFromSite2()
	}()

	go func() {
		defer wg.Done() //executa quando função foi executada
		priceChannel <- fetcher.FetchPriceFromSite3()
	}()

	wg.Wait() //aguarda terminarem

	close(priceChannel) //fecha o canal para optimização

	fmt.Printf("Tempo total: %s", time.Since(start))

}
