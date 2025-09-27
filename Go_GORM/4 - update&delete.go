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

	//atualizando
	var p Product
	db.First(&p, 1)
	p.Name = "ROLA"
	db.Save(&p)

	var p2 Product
	db.First(&p2, 1)
	p2.Name = "ROLA DOIS"
	db.Save(&p2)

	//atualizando multiplas  colunas
	db.Model(&p).Updates(Product{Price: 100, Name: "Xbox  Series  X"}) //Sem campos  vazios

	//exclusao de produtos
	var p3 Product
	db.First(&p3, 2)
	db.Delete(&p3)
}
