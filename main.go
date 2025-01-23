package main

import "github.com/pallandir/books-api/router"

func main() {
	router := router.ApiRouter()
	router.Run(":8080")
}
