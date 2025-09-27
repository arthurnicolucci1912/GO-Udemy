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
// 	CategoryID   int
// 	Category     Category
// 	SerialNumber SerialNumber
// 	gorm.Model
// }

// type Category struct {
// 	ID   int `gorm: primaryKey`
// 	Name string
// }

// type SerialNumber struct {
// 	ID        int `gorm: primaryKey`
// 	Number    string
// 	ProductID int
// }

// func main() {
// 	//cria a conexãa com o BD
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

// 	//caso de erro..
// 	if err != nil {
// 		panic("erro ao abrir banco de dados")
// 	}

// 	//cria as conexões
// 	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

// 	//inserindo categoria
// 	category := Category{
// 		Name: "Eletronicos",
// 	}
// 	db.Create(&category)

// 	//inserindo Produtoa
// 	db.Create(&Product{
// 		Name:       "PC",
// 		Price:      16000,
// 		CategoryID: category.ID,
// 	})

// 	//Inserindo serialNumber
// 	db.Create(&SerialNumber{
// 		Number:    "123123",
// 		ProductID: ,
// 	})

// 	//listando produtos
// 	//realacionamento de 1 para 1
// 	var products []Product
// 	db.Preload("Category").Preload("Serial Number").Find(&products)
// 	for _, product := range products {
// 		fmt.Println(product.Name, "|", product.Category.Name, "|", product.Price, "|", product.SerialNumber.Number)
// 	}
// }
