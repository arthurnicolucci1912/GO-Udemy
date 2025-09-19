package main

import "fmt"

func main() {
	nomes := []string{"Ana", "Maria", "João", "Pedross"}

	fmt.Println("Os nomes são: ", nomes)

	//quando nao queremos usar o indice colocamos _ no lugar
	for id, nome := range nomes {
		fmt.Println(id, "->", nome)
		if len(nome) > 5 {
			fmt.Println(nome, "Nome grande")
		}
	}

}
