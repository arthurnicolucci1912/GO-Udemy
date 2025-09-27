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

	//inserir dados
	db.Create(&Product{Name: "PLUG", Price: 100.00})
}
