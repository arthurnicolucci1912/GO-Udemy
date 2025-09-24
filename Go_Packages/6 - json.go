package main

import (
	"encoding/json"
	"fmt"
)

//JSON: JavaScript object notation

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonString := `{"name": "Arthur", "age":"18 anos"}`
	//referencia para struct
	var person Person

	//decerioaliza os dados
	json.Unmarshal([]byte(jsonString), &person)
	fmt.Printf("Nome: %s, Idade: %d", person.Name, person.Age)

	person.Name = "Joao"
	person.Age = 19

	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}
