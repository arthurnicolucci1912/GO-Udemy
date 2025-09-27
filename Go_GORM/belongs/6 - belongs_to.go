package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm: primaryKey`
	Name string
}

type Product struct {
	ID    int `gorm: primaryKey`
	Name  string
	Price float64

	//variaveis que referenciam "Category"
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	//cria a conexãa com o BD
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	//caso de erro..
	if err != nil {
		panic("erro ao abrir banco de dados")
	}

	db.AutoMigrate(&Product{}, &Category{})

	// //inserindo categoria
	// category := Category{
	// 	Name: "Cozinha",
	// }
	// db.Create(&category)

	// //inserindo Produto
	// db.Create(&Product{
	// 	Name:       "Mesa",
	// 	Price:      6000,
	// 	CategoryID: category.ID,
	// })

	//listando produtos
	var products []Product
	db.Preload("Category").Find(products)
	for _, product := range products {
		fmt.Println(product.Name, "|", product.Category.Name, "|", product.Price)
	}

}
