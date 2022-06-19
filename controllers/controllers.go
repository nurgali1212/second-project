package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nurgali1212/second-project/models"
)

type CreateTodoinput struct {
	Description string `json:"description"  binding:"required"`
	CompletedAt string `json:"completed_at"  binding:"required"`
}
// GET /todos
// Получаем список всех todo
func GetAllTodo(context *gin.Context) {
	var todos []models.Todo
	models.DB.Find(&todos)

	context.JSON(http.StatusOK, gin.H{"todos": todos})
}

// POST /todos
// Создание file
func CreateTodo(context *gin.Context) {
	var input CreateTodoinput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	completedAt, err := time.Parse(time.RFC3339, input.CompletedAt)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": completedAt})

		return
	}

	todos := models.Todo{Description: input.Description, CreatedAt: time.Now(), CompletedAt: completedAt}
	models.DB.Create(&todos)

	context.JSON(http.StatusOK, gin.H{"todos": todos})


}


type UpdateTodoInput struct {
	Description string `json:"description"`
	
 }


 // PUT /todos/:id
// Изменения информации
func UpdateTodo(context *gin.Context) {
	// Проверяем имеется ли такая запись перед тем как её менять
	var todo models.Todo
	if err := models.DB.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
	   context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
	   return
	}
 
	var input UpdateTodoInput
	if err := context.ShouldBindJSON(&input); err != nil {
	   context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	   return
	}
 
	models.DB.Model(&todo).Update(input)
 
	context.JSON(http.StatusOK, gin.H{"todos": todo})
}



// DELETE /todos/:id
// Удаление
func DeleteTodo(context *gin.Context) {
	// Проверка прежде чем удалять, есть ли такой файл
	var todo models.Todo
	if err := models.DB.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
	   context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
	   return
	}
 
	models.DB.Delete(&todo)
 
	context.JSON(http.StatusOK, gin.H{"todos": true})
 }