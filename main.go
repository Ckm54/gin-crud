package main

import (
	"github.com/Ckm54/bookstore-go/models"
	"github.com/Ckm54/bookstore-go/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)

	r.Run()
}