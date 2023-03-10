package main

import (
	"fmt"
	"go_gin_gorm_postgres_2/controllers"
	"go_gin_gorm_postgres_2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// Name       string `json:"name" binding:"required"`
	fmt.Println("Starting Application...")
	database.StartDB()

	router := gin.Default()
	router.POST("/books", controllers.CreateBook)
	router.GET("/books", controllers.GetAllBook)
	router.GET("/books/:bookID", controllers.GetBookByID)
	router.PUT("/books/:bookID", controllers.UpdateBookByID)
	router.DELETE("/books/:bookID", controllers.DeleteBook)

	router.Run(":8080")
}