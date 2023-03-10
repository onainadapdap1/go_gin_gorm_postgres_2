package controllers

import (
	"fmt"
	"go_gin_gorm_postgres_2/database"
	"go_gin_gorm_postgres_2/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	db := database.GetDB()
	var newBook models.Book

	// binding json data into struct field
	if err := c.ShouldBindJSON(&newBook); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error to bind json into struct field",
		})
		return
	}

	res := db.Create(&newBook)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error creating book",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"book": newBook,
	})
}

func GetAllBook(c *gin.Context) {
	db := database.GetDB()
	books := []models.Book{}

	res := db.Find(&books)
	if res.Error != nil {
		fmt.Println("error is here")
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no data is founded",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"books": books,
	})
}

func GetBookByID(c *gin.Context) {
	db := database.GetDB()
	id := c.Param("bookID")
	book := models.Book{}

	res := db.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "book is not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"book": book,
	})
}

func UpdateBookByID(c *gin.Context) {
	db := database.GetDB()
	book := models.Book{}
	id := c.Param("bookID")
	
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "error bind json data into slice field",
		})
		return
	}

	updateBook := models.Book{}
	res := db.Model(&updateBook).Where("id = ?", id).Updates(book)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "book not updated",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"newBook": book,
	})
}

func DeleteBook(c *gin.Context) {
	db := database.GetDB()
	book := models.Book{}
	id := c.Param("bookID")
	
	res := db.Find(&book, id)
	if res.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "no data is found",
		})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "book deleted successfully",
	})
	
}