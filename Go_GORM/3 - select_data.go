package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm: primaryKey`
	Name  string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to  connect database")
	}

	//selecionando um produto
	//referencia para struct
	// var product Product
	// // db.First(&product, 1)
	// db.First(&product, "name = ?", "PLUG")
	// fmt.Println(product)

	//selecionando todos

	// var products = []Product{}
	// db.Limit(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	var products = []Product{}
	db.Where("price > ?", 1).Find(&products)
	for _, product := range products {
		fmt.Println(product)
	}
}
