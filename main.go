package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pallandir/books-api/handlers"
)

func apiRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:isbn", handlers.GetBookByISBN)
	router.POST("/books", handlers.PostBook)

	return router
}

func main() {
	router := apiRouter()
	router.Run(":8080")
}
