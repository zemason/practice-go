package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"practice-go/config"
	"practice-go/database"
	"practice-go/handlers"
	"practice-go/middleware"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Setup Gin
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	// Set Gin mode
	if gin.Mode() == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to database
	database.ConnectDB()

	// Initialize Gin
	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.Logger())

	// Routes
	setupRoutes(r)

	// Start server
	fmt.Println("\n📚 ==================================")
	fmt.Println("🚀 REST API Buku dengan Gin & PostgreSQL")
	fmt.Println("======================================")
	fmt.Printf("📍 Server running on: http://localhost:%s\n", cfg.Port)
	fmt.Println("📊 Database: PostgreSQL")
	fmt.Println("======================================")

	r.Run(":" + cfg.Port)
}

func setupRoutes(r *gin.Engine) {
	// Home
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Book API",
			"version": "1.0.0",
			"endpoints": []string{
				"GET    /books           - Get all books",
				"GET    /books/:id       - Get book by ID",
				"POST   /books           - Create new book",
				"PUT    /books/:id       - Update book",
				"DELETE /books/:id       - Delete book",
				"GET    /books/search    - Search books",
			},
		})
	})

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"database":  "connected",
			"timestamp": time.Now().Unix(),
		})
	})

	// Book routes
	books := r.Group("/books")
	{
		books.GET("", handlers.GetBooks)
		books.GET("search", handlers.SearchBooks)
		books.GET(":id", handlers.GetBook)
		books.POST("", handlers.CreateBook)
		books.PUT(":id", handlers.UpdateBook)
		books.DELETE(":id", handlers.DeleteBook)
	}
}
