package main

import "fmt"

func main() {
	var n1, n2 int

	fmt.Print("n1: ")
	fmt.Scan(&n1)

	fmt.Print("n2: ")
	fmt.Scan(&n2)

	//:= serve para variáveis locais
	soma := n1 + n2
	sub := n1 - n2
	mult := n1 * n2
	div := n1 / n2
	mod := n1 % n2

	fmt.Printf("Soma de %d e %d: %d\n", n1, n2, soma)
	fmt.Printf("Sub de %d e %d: %d\n", n1, n2, sub)
	fmt.Printf("Multiplicação de %d e %d: %d\n", n1, n2, mult)
	fmt.Printf("Divisão de %d e %d: %d\n", n1, n2, div)
	fmt.Printf("Resto de %d e %d: %d \n", n1, n2, mod)

	//atribuição
	n1 += 1 //n1 = n1+1
	n1 -= 1 //n1 = n1-1
	n1 *= 1 //n1 = n1*1
	n1 /= 1 //n1 = n1/1

	//comparação
	bigger := n1 > n2
	smaller := n1 < n2
	equals := n1 == n2
	diferent := n1 != n2
	biggerEqual := n1 >= n2
	smallerEqual := n1 <= n2

	fmt.Printf("Os numeros %d e %d São iguais?: %v \n", n1, n2, equals)
	fmt.Printf("Os numeros %d e %d São diferentes?: %v \n", n1, n2, diferent)
	fmt.Printf("Os numeros %d e %d menores?: %v \n", n1, n2, smaller)
	fmt.Printf("Os numeros %d e %d maiores?: %v \n", n1, n2, bigger)
	fmt.Printf("Os numeros %d e %d maiores ou igual: %v \n", n1, n2, biggerEqual)
	fmt.Printf("Os numeros %d e %d menores ou igual: %v \n", n1, n2, smallerEqual)

	fmt.Print("Valor final de n1 apos as alterações: ", n1)

}
