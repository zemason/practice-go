package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"practice-go/database"
	"practice-go/models"
)

// GetBooks - GET /books
func GetBooks(c *gin.Context) {
	var books []models.Book

	// Query ke database
	result := database.DB.Find(&books)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data buku",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    books,
		"total":   len(books),
	})
}

// GetBook - GET /books/:id
func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	// Cari buku by ID
	result := database.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Buku tidak ditemukan",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    book,
	})
}

// CreateBook - POST /books
func CreateBook(c *gin.Context) {
	var input models.BookRequest

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Data tidak valid",
			"details": err.Error(),
		})
		return
	}

	// Buat book object
	book := models.Book{
		Title:  input.Title,
		Author: input.Author,
		Year:   input.Year,
		Price:  input.Price,
	}

	// Simpan ke database
	result := database.DB.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal menyimpan buku",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Buku berhasil ditambahkan",
		"data":    book,
	})
}

// UpdateBook - PUT /books/:id
func UpdateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book
	var input models.BookRequest

	// Cek apakah buku ada
	result := database.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Buku tidak ditemukan",
		})
		return
	}

	// Bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Data tidak valid",
		})
		return
	}

	// Update book
	book.Title = input.Title
	book.Author = input.Author
	book.Year = input.Year
	book.Price = input.Price

	database.DB.Save(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Buku berhasil diupdate",
		"data":    book,
	})
}

// DeleteBook - DELETE /books/:id
func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	// Cek apakah buku ada
	result := database.DB.First(&book, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Buku tidak ditemukan",
		})
		return
	}

	// Hapus buku
	database.DB.Delete(&book)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Buku berhasil dihapus",
	})
}

// SearchBooks - GET /books/search
func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	minYear := c.Query("min_year")
	maxYear := c.Query("max_year")

	var books []models.Book
	query := database.DB

	// Filter by title (case insensitive, partial match)
	if title != "" {
		query = query.Where("LOWER(title) LIKE LOWER(?)", "%"+title+"%")
	}

	// Filter by author
	if author != "" {
		query = query.Where("LOWER(author) LIKE LOWER(?)", "%"+author+"%")
	}

	// Filter by year range
	if minYear != "" {
		if year, err := strconv.Atoi(minYear); err == nil {
			query = query.Where("year >= ?", year)
		}
	}

	if maxYear != "" {
		if year, err := strconv.Atoi(maxYear); err == nil {
			query = query.Where("year <= ?", year)
		}
	}

	// Eksekusi query
	query.Find(&books)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    books,
		"total":   len(books),
	})
}
