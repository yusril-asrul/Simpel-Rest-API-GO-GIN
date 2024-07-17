package main

import (
	"go-gin-crud/database"
	"go-gin-crud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
    // Inisialisasi database
    database.InitDatabase()

    // Inisialisasi router
    router := gin.Default()

    // Register endpoints
    router.GET("/books", handlers.GetBooks)
    router.GET("/books/:id", handlers.GetBookByID)
    router.POST("/books", handlers.CreateBook)
    router.PUT("/books/:id", handlers.UpdateBook)
    router.DELETE("/books/:id", handlers.DeleteBook)

    // Jalankan server
    router.Run(":8080")
}
