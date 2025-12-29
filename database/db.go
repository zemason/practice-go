package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"practice-go/config"
	"practice-go/models"
)

var DB *gorm.DB

func ConnectDB() {
	cfg := config.LoadConfig()

	// Buat connection string PostgreSQL
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	// Connect ke database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	fmt.Println("✅ Database connected successfully!")

	// Auto migrate models
	DB.AutoMigrate(&models.Book{})

	fmt.Println("✅ Database migrated!")

	// Seed data contoh (opsional)
	seedData()
}

func seedData() {
	// Cek apakah sudah ada data
	var count int64
	DB.Model(&models.Book{}).Count(&count)

	if count == 0 {
		books := []models.Book{
			{Title: "Harry Potter", Author: "J.K. Rowling", Year: 1997, Price: 150000},
			{Title: "Laskar Pelangi", Author: "Andrea Hirata", Year: 2005, Price: 85000},
			{Title: "Bumi Manusia", Author: "Pramoedya Ananta Toer", Year: 1980, Price: 95000},
		}

		DB.Create(&books)
		fmt.Println("✅ Sample data added!")
	}
}
