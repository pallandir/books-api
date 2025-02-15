package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pallandir/books-api/handlers"
)

func ApiRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:isbn", handlers.GetBookByISBN)
	router.POST("/books", handlers.PostBook)
	router.DELETE("/books/:isbn", handlers.DeleteBookByISBN)
	router.PUT("/books/:isbn", handlers.UpdateBooksByISBN)
	return router
}
