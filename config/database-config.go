package config

import (
	"fmt"
	"os"

	"github.com/avtara/testcov-backend/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabaseConnection is creating new connection to DB
func SetupDatabaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{}, &entity.Hospital{}, &entity.Schedule{}, &entity.Order{})
	return db
}

//CloseDatabaseConnection is closing connection between app and db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err == nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
