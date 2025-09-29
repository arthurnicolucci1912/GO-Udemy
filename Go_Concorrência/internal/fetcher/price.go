package fetcher

import (
	"math/rand/v2"
	"time"
)

//buscar preços e diferentes sites
//simula um site que demora 1 segundo
//ir na api e buscar preço...futuraente

func FetchPriceFromSite1() float64 {
	time.Sleep(1 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite2() float64 {
	time.Sleep(2 * time.Second)
	return rand.Float64() * 100
}

func FetchPriceFromSite3() float64 {
	time.Sleep(3 * time.Second)
	return rand.Float64() * 100
}
