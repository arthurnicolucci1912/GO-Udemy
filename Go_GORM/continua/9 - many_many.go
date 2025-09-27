package main

import (
	"fmt" // Import fmt for printing

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Corrected Product struct for Many-to-Many
type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64

	// The field must be a SLICE of the associated struct
	Categories []Category `gorm:"many2many:products_categories;"`
	gorm.Model
}

type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("erro ao abrir banco de dados")
	}

	// Cria as tabelas e a tabela de junção 'products_categories'
	db.AutoMigrate(&Product{}, &Category{})

	// inserindo categoria
	category := Category{Name: "Cozinha"}
	db.FirstOrCreate(&category, Category{Name: "Cozinha"})

	category2 := Category{Name: "Eletrodomésticos"}
	db.FirstOrCreate(&category2, Category{Name: "Eletrodomésticos"})

	// Inserindo Produto (agora funciona com a slice)
	db.Create(&Product{
		Name:  "Fogão",
		Price: 100.00,
		// O tipo de Categories em Product é []Category, combinando com o valor.
		Categories: []Category{category, category2},
	})

	// Listando (Bloco de código que você tinha comentado)
	var categories []Category
	// Preload is essential to load the associated Products
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Println(c.Name, ":")
		for _, product := range c.Products {
			fmt.Println("->", product.Name)
		}
	}
}
