package routes

import (
	"gin_base/internal/handlers"
	"github.com/gin-gonic/gin"
)

// SetupRoutesForApp1 настраивает маршруты для приложения myapp1
func SetupRoutesForApp1(r *gin.Engine) {
	// Пример маршрута GET на "/app1"
	r.GET("/app1", handlers.App1Handler)

	// Добавление других маршрутов для myapp1...
}

// SetupRoutesForApp2 настраивает маршруты для приложения myapp2
func SetupRoutesForApp2(r *gin.Engine) {
	// Пример маршрута GET на "/app2"
	r.GET("/app2", handlers.App2Handler)

	// Добавление других маршрутов для myapp2...
}
