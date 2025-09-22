package main

import "fmt"

func main() {
	//criando um map, com a nome do aluno como chave e nota como valor
	estudantes := map[string]float64{
		"Maria": 9.1,
		"João":  8.3,
		"Pedro": 5.5,
	}
	fmt.Println("Classificação dos estudantes: ")
	for nome, nota := range estudantes {
		status := "reprovado"
		if nota >= 6.0 {
			status = "aprovado"
		}
		fmt.Printf("| Aluno: %s| Nota: %.2f| Status: %s| \n", nome, nota, status)
	}
}
