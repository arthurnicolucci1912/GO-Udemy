package main

import "fmt"

func factorial(num int) int {
	if num == 1 {
		return 1
	} else {
		return num * factorial(num-1)
	}

}

func totalSum(num int) int {
	if num == 1 {
		return num
	}
	return num + totalSum(num-1)
}

func main() {
	var number int
	fmt.Print("Fatorial: ")
	fmt.Scan(&number)
	fmt.Printf("O fatorial de %d: %d\n", number, factorial(number))
	var num int
	fmt.Print("Digite um numero para soma: ")
	fmt.Scan(&num)
	fmt.Printf("A soma total do %d: %d\n", num, totalSum(num))
}
