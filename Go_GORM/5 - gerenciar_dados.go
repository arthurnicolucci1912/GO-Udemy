package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	ID    int `gorm:  primaryKey`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("erro ao conectar no  bnaco de dados")
	}

	db.AutoMigrate(&Product{})

	//isnerindo
	products := []Product{
		{Name: "ACER NITRO V5", Price: 5000},
		{Name: "Xbox Series X", Price: 4500},
	}
	db.Create(&products)

	var p Product
	db.First(&p, 1)
	p.Name = ("Xbox Series SEXO")
	db.Save(&p)
}
