package main

import (
	"FlowerHive/pkg/products"
	"log"
	"net/http"
)

func main() {
	// Инициализация подключения к базе данных
	err := products.InitializeDB("postgres://postgres:postgres@localhost:5432/flower_catalog?sslmode=disable")
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	// Настройка HTTP-обработчиков
	router := products.SetupRouter()

	// Запуск веб-сервера на порту 8080
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
