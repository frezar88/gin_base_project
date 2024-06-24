package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// App1Handler обработчик для маршрута "/app1"
func App1Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Handler for MyApp1",
	})
}

// App2Handler обработчик для маршрута "/app2"
func App2Handler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Handler for MyApp2",
	})
}
