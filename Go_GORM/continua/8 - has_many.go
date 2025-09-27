// // ideia, ao final criar um sistema "completo" de cadastro de produtos
// package main

// import (
// 	"fmt"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	ID    int `gorm: primaryKey`
// 	Name  string
// 	Price float64

// 	//variaveis que referenciam "Category"
// 	CategoryID int
// 	Category   Category
// }

// type Category struct {
// 	ID       int `gorm: primaryKey`
// 	Name     string
// 	Products []Product
// }

// func main() {
// 	//cria a conexãa com o BD
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	//caso de erro..
// 	if err != nil {
// 		panic("erro ao abrir banco de dados")
// 	}

// 	//cria as conexões
// 	db.AutoMigrate(&Product{})

// 	// // // //inserindo categoria
// 	// category := Category{
// 	// 	Name: "Cozinha",
// 	// }
// 	// db.Create(&category)

// 	//inserindo Produto
// 	db.Create(&Product{
// 		Name:       "Fogão",
// 		Price:      100.00,
// 		CategoryID: 2,
// 	})

// 	//listando
// 	var categories []Category
// 	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, category := range categories {
// 		fmt.Println(category.Name, ":")
// 		for _, product := range category.Products {
// 			fmt.Println("->", product.Name)
// 		}
// 	}
// }
