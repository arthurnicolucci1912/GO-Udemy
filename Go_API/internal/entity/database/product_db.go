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

func (p *Product) FindByID(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(id string) (*entity.Product, error) {
	_, err := p.FindByID(product.ID.String())
	if err != nil {
		return nil
	}
	return p.DB.Update(product).Error
}

func (p *Product) Delete(id string) (*entity.Product, error) {
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
		a
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created at " + sort).Find(&products).Error
	}
}
