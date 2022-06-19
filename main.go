package main

import (
	"github.com/nurgali1212/second-project/controllers"
    "github.com/nurgali1212/second-project/models"
    "github.com/gin-gonic/gin"
   
)

func main() {
   route := gin.Default()

   // Подключение к базе данных
   models.ConnectDB()

   // Маршруты
   route.GET("/todos", controllers.GetAllTodo)
   route.POST("/todos", controllers.CreateTodo)
   route.PUT("/todos/:id", controllers.UpdateTodo) 
   route.DELETE("/todos/:id", controllers.DeleteTodo)

   // Запуск сервера
   route.Run()
}



