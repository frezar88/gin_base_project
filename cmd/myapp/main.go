package main

import (
	"gin_base/internal/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	// Подключение маршрутов для myapp1
	routes.SetupRoutesForApp1(r)

	// Запуск сервера на порту 8080 для myapp1
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server for myapp1: %v", err)
	}
}
