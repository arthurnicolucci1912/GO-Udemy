// gerenciamento de controle
package services

import (
	"estoque/internal/models"
	"fmt"
	"strconv"
	"time"
)

type Estoque struct {
	items map[string]models.Item
	logs  []models.Log
}

// passar informação do estoque por valor
// ponteiro
func NewEstoque() *Estoque {
	return &Estoque{
		items: make(map[string]models.Item),
		logs:  []models.Log{},
	}
}

// Método para adicionar item ao estoque
// primeiro faz referência ao estoque, e depois ao tipo de objeto que será adicionado nele
//
//	-->receptor<--    -->item da struct ITEM<--      , que se localiza dentro de models, e a struct ITEM
func (estoque *Estoque) AddItem(item models.Item) error {
	if item.Quantity <= 0 {
		return fmt.Errorf("erro ao adicionar item: [ID: %d] possui uma quantidade menor insuficiente (0 ou negativa)", item.ID)
	}
	//
	existingItem, exists := estoque.items[strconv.Itoa(item.ID)]
	//                  ^-- A chave foi encontrada? (Booleano)
	// ^-- O item COMPLETO (ID, Nome, Quantidade, Preço, etc.)
	if exists {
		item.Quantity += existingItem.Quantity
	}
	estoque.items[strconv.Itoa(item.ID)] = item

	//simulando logs
	estoque.logs = append(estoque.logs, models.Log{
		TimeStamp: time.Now(),
		Action:    "Entrada de item no estoque",
		User:      "Admin",
		ItemId:    item.ID,
		Quantity:  item.Quantity,
		Reason:    "Adicionando itens ao estoque",
	})
	return nil
}

// metodo de listagem de items do estoque
func (estoque *Estoque) ListItems() []models.Item {
	var List []models.Item
	for _, item := range estoque.items {
		//adiciona cada item a uma lista, para ser exibido
		List = append(List, item)
	}
	return List
}

// metodo de lista de logs
func (estoque *Estoque) ShowLogs() []models.Log {
	return estoque.logs
}

//metodo para calcular preço

func (estoque *Estoque) CalculateTotal() float64 {
	var total float64
	for _, item := range estoque.items {
		total += float64(item.Quantity) * item.Price
	}
	return total
}

func FindBy[element any](data []element, comparator func(element) bool) ([]element, error) {
	var result []element

	for _, value := range data {
		if comparator(value) {
			result = append(result, value)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("nenhum item foi encontrado")
	}
	return result, nil
}

// func FindByName(data []models.Item, name string) ([]models.Item, error) {
// 	var result []models.Item

// 	for _, item := range data {
// 		if item.Name == name {
// 			result = append(result, item)
// 		}
// 	}

// 	if len(result) == 0 {
// 		return nil, fmt.Errorf("Nenhum item com o nome '%s' encontrado.", name)
// 	}
// 	return result, nil
// }
