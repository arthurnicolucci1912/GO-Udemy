package main

import (
	"estoque/internal/models"
	"estoque/internal/services"
	"fmt"
)

func main() {
	fmt.Println("[---Gerenciamento de estoque---]")

	//varios items
	items := []models.Item{
		{ID: 1, Name: "Xbox One", Quantity: 2, Price: 3000.0},
		{ID: 2, Name: "PC", Quantity: 1, Price: 2000.0},
		{ID: 3, Name: "Ps5", Quantity: 4, Price: 10000.0},
		{ID: 4, Name: "Notebok", Quantity: 2, Price: 100},
	}

	//instanciamoos um novo estoque
	estoque := services.NewEstoque()

	//para exibir os items inseridor precisamos de um laço, para percorrer o slice []models.Item
	for _, item := range items {
		err := estoque.AddItem(item)
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
	// fmt.Println(estoque.ListItems())
	//laço para deixar a listagem ainda melhor
	for _, item := range estoque.ListItems() {
		fmt.Printf("\nID: %d | Item: %s | Quantidade: %d | Preço: %.2f",
			item.ID, item.Name, item.Quantity, item.Price)
	}

	fmt.Println()

	// fmt.Println(estoque.ShowLogs())
	logs := estoque.ShowLogs()

	//laço para deixart a listagem ainda melhor
	for _, log := range logs {
		fmt.Printf("\n[%s] Ação: %s - Usuário: %s - Item ID: %d - Quantidade: %d - Motivo: %s",
			log.TimeStamp.Format("02/01 15:04:05"), log.Action, log.User, log.ItemId, log.Quantity, log.Reason)
	}

	fmt.Println("\nValor total R$: ", estoque.CalculateTotal())

	searchItem, err := services.FindBy(items, func(item models.Item) bool {
		// return item.Name == "Ps5"
		// return item.Price <= 1000.00
		return item.ID <= 4
	})
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(searchItem)

}
