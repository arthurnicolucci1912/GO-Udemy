package database

import (
	"apis/internal/entity"

	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}
func (p *Product) Update(product *entity.Product) error {

	// 1. Opcional, mas recomendado: verificar se o produto existe
	_, err := p.FindById(product.ID.String())
	if err != nil {
		// Se o FindByID retornar um erro (ex: não encontrado), propaga
		return err
	}

	// 2. Usar o GORM Save para atualizar o registro completo.
	// O GORM usa o ID do struct 'product' para saber qual linha atualizar.
	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	//exemplo de implementação
	return p.DB.Delete(&entity.Product{}, "id = ?", id).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created at " + sort).Find(&products).Error
	}
	return products, err
}
