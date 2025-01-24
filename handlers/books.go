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
	context.JSON(http.StatusNotFound, gin.H{"message": "Book not found"})
}

func DeleteBookByISBN(context *gin.Context) {
	isbn := context.Param("isbn")
	for index, book := range database.Books {
		if book.ISBN == isbn {
			database.Books = append(database.Books[:index], database.Books[index+1:]...)
			context.JSON(http.StatusOK, database.Books)
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "No book to delete found"})
}

func UpdateBooksByISBN(context *gin.Context) {
	isbn := context.Param("isbn")
	var updatedBook models.Book
	//BindJSON needs the adress of the struct to directly modify it instead of working on a copy
	if err := context.BindJSON(&updatedBook); err != nil {
		return
	}

	for index, book := range database.Books {
		if book.ISBN == isbn {
			if updatedBook.Author != "" {
				database.Books[index].Author = updatedBook.Author
			}
			if updatedBook.Title != "" {
				database.Books[index].Title = updatedBook.Title
			}
			context.JSON(http.StatusCreated, database.Books[index])
			return
		}
	}
	context.JSON(http.StatusNotFound, gin.H{"message": "No book to update found"})

}
