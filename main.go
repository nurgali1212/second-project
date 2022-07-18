package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nurgali1212/second-project/controllers"
	"github.com/nurgali1212/second-project/models"
)

func main() {
	route := gin.Default()

	// Подключение к базе данных
	models.ConnectDB()

	// Маршруты
	route.GET("/books", controllers.GetAllBooksID)
	route.GET("/books/:id", controllers.GetBookID)
	route.GET("/categories", controllers.GetAllCategory)
	route.POST("/books", controllers.CreateBooks)
	route.PUT("/books/:id", controllers.UpdateBooks)
	route.DELETE("/books/:id", controllers.DeleteBooks)

	// Запуск сервера
	route.Run()
}
