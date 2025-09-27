package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm: primary key`
	Name  string
	Price float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//migrar a tabela
	db.AutoMigrate(&Product{})

	// //inserir dados
	// db.Create(&Product{Name: "PLUG", Price: 100.00})

	//slice de produtos para facilitar a  criação  do  banco  sem código SQL
	products := []Product{
		{Name: "Acer  Nitro V5", Price: 5.500},
		{Name: "Dell  Latitude", Price: 1.500},
		{Name: "PAU IMENSO", Price: 10.000},
	}
	db.Create(&products)
}
