package main

import (
	"apis/configs"
	"apis/internal/database"
	"apis/internal/dto"
	"apis/internal/entity"
	"encoding/json"
	"log"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, float64(product.Price))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func main() {
	// 1. CHECAGEM DE CONFIGURAÇÃO
	_, err := configs.LoadConfig(".")
	if err != nil {
		// Agora você verá a mensagem de erro
		log.Fatal("Erro ao carregar configurações:", err)
	}

	// criar o banco
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		// Agora você verá a mensagem de erro
		log.Fatal("Erro ao conectar/abrir o banco de dados:", err)
	}

	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := NewProductHandler(productDB)
	http.HandleFunc("/products", productHandler.CreateProduct)

	// 3. INICIAR O SERVIDOR
	log.Println("Servidor iniciado na porta :8000")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Erro ao iniciar o servidor HTTP:", err)
	}
}
