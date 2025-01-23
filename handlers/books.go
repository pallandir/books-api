package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pallandir/books-api/database"
	"github.com/pallandir/books-api/models"
)

func GetBooks(context *gin.Context) {
	context.JSON(http.StatusOK, database.Books)
}

func PostBook(context *gin.Context) {
	var newBook models.Book

	//BindJSON needs the adress of the struct to directly modify it instead of working on a copy
	if err := context.BindJSON(&newBook); err != nil {
		return
	}
	database.Books = append(database.Books, newBook)
	context.JSON(http.StatusCreated, newBook)
}

func GetBookByISBN(context *gin.Context) {
	isbn := context.Param("isbn")
	for _, books := range database.Books {
		if books.ISBN == isbn {
			context.JSON(http.StatusOK, books)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
}
