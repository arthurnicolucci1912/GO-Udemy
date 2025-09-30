package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string)

	var wg sync.WaitGroup

	//criando 3 goroutines q	ue vao ler o canal
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for msg := range ch {
				fmt.Printf("Goroutine %d recebeu %s\n", id, msg)
				time.Sleep(time.Second * 500)
			}
		}(i)
	}

	//ebnviando msg para canal
	messages := []string{"Hello", "World", "Go", "concorrencia", "é", "incrivel"}
	for _, msg := range messages {
		ch <- msg
		time.Sleep(time.Microsecond * 500)
	}

	close(ch)

	wg.Wait()
	fmt.Printf("Todas as Goroutines terminaram")

}
