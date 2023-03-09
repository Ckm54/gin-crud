package controllers

import (
	"net/http"

	"github.com/Ckm54/bookstore-go/models"
	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title		string		`json:"title" binding:"required"`
	Author	string		`json:"author" binding:"required"`
}

type UpdateBookInput struct {
	Title		string 		`json:"title"`
	Author	string		`json:"author"`
}

// get all books from database
func FindBooks(ctx *gin.Context) {
	var books []models.Book
	models.DB.Find(&books)

	ctx.JSON(http.StatusOK, gin.H{"data": books})
}

// create a new book
func CreateBook(ctx *gin.Context) {
	// validate inputs
	var input CreateBookInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create book
	book := models.Book{Title: input.Title, Author: input.Author}
	models.DB.Create(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// Find a single book
func FindBook(ctx *gin.Context) {
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record Not Found!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// update a book's information
func UpdateBook(ctx *gin.Context) {
	// check if model exists
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateBookInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// update the book's information in database
	// walkaround to data update error
	if input.Title != "" && input.Author != "" {
		updateBook := models.Book{Author: input.Author, Title: input.Title}
		models.DB.Model(&book).Updates(&updateBook)
	}else if input.Title != "" {
		updateBook := models.Book{Title: input.Title}
		models.DB.Model(&book).Updates(&updateBook)
	} else if input.Author != "" {
		updateBook := models.Book{Author: input.Author}
		models.DB.Model(&book).Updates(&updateBook)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": book})
}

// Delete a book
func DeleteBook(ctx *gin.Context) {
	// get the model if exists
	var book models.Book

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&book).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&book)

	ctx.JSON(http.StatusOK, gin.H{"data": true})
}