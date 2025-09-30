package models

import "fmt"

type Item struct {
	ID       int
	Name     string
	Quantity int
	Price    float64
}

// receptor para struct Item, objetos para inicar ou instânciar nossa struct (this,self)
func (i Item) Info() string {
	return fmt.Sprintf("ID: %d | Name: %s | Quantity: %d | Price: %.2f",
		i.ID, i.Name, i.Quantity, i.Price)

}
