package main

import "fmt"

//definindo uma struct
type Carro struct {
	Modelo string
	Ano    int
	Cor    string
}

func main() {

	//criar a instancia da struct carro
	carro1 := Carro{
		Modelo: "Saveiro",
		Ano:    2025,
		Cor:    "Amarela",
	}

	fmt.Println("Info do carro")
	fmt.Printf("Modelo: %s | Ano: %d | Cor: %s", carro1.Modelo, carro1.Ano, carro1.Cor)

}
