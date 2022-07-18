package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nurgali1212/second-project/models"
)

type CreateBookInput struct {
	Title      string `json:"title"`
	Author     string `json:"author" `
	CategoryID uint   `json:"category_id"`
}

func GetAllCategory(context *gin.Context) {
	var categories []models.Category
	models.DB.Find(&categories)

	context.JSON(http.StatusOK, gin.H{"categories": categories})
}

// GET  BOOKS / ID

func GetBookID(context *gin.Context) {

	var books models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&books).Error; err != nil {
	   context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
	   return
	}
	models.DB.Find(&books)
	context.JSON(http.StatusOK, gin.H{"books": books})
 }

//GET BOOKS
func GetAllBooksID(context *gin.Context) {
	var books []models.Book
	models.DB.Preload("Category").Find(&books)


	context.JSON(http.StatusOK, gin.H{"books": books})
}

//POST BOOKS
func CreateBooks(context *gin.Context) {

	var input CreateBookInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	books := models.Book{Author: input.Author, Title: input.Title, CategoryID: input.CategoryID}
	models.DB.Create(&books)

	context.JSON(http.StatusOK, gin.H{"books": books})

}

type UpdateBooksInput struct {
	Title      string `json:"title"`
	Author     string `json:"author"`
	CategoryID uint   `json:"category_id"`
}

// PUT BOOKS

func UpdateBooks(context *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateBooksInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&book).Update(input)

	context.JSON(http.StatusOK, gin.H{"books": book})
}

// DELETE BOOKS

func DeleteBooks(context *gin.Context) {

	var book models.Book
	if err := models.DB.Where("id = ?", context.Param("id")).First(&book).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	models.DB.Delete(&book)

	context.JSON(http.StatusOK, gin.H{"books": true})
}
