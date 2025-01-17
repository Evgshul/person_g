package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func InitDb() *gorm.DB {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")
	charset := "utf8mb4"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username, password, host, port, database, charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get SQL DB: %v", err)
	}

	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(10 * time.Minute)

	//createTableQuery := `
	//CREATE TABLE IF NOT EXISTS persons (
	//    id INT AUTO_INCREMENT PRIMARY KEY,
	//    fullname VARCHAR(100),
	//    gender VARCHAR(10),
	//    phone_number VARCHAR(20),
	//    email VARCHAR(100)
	//);`
	//err = db.Create(createTableQuery).Error
	//if err != nil {
	//	log.Fatalf("Failed to create table: %v", err)
	//}
	log.Println("Database connected and initialized successfully!")
	return db
}

func CreatTablePerson(db gorm.DB) *gorm.DB {
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS persons (
        id INT AUTO_INCREMENT PRIMARY KEY,
        fullname VARCHAR(100),
        gender VARCHAR(10),
        phone_number VARCHAR(20),
        email VARCHAR(100)
    );`

	err := db.Create(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	return &db
}

func CloseDatabaseConnection() {
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("Failed to close database connection: %v", err)
		return
	}
	sqlDB.Close()
}
