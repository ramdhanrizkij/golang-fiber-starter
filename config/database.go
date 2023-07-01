package config

import (
	"fmt"
	"golang-fiber/app/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupDatabaseConnection() *gorm.DB {
	dbUser := GetEnv("DB_USER")
	dbPass := GetEnv("DB_PASSWORD")
	dbHost := GetEnv("DB_HOST")
	dbName := GetEnv("DB_NAME")
	dbPort := GetEnv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&models.Permission{}, &models.Role{}, &models.PermissionRole{}, &models.User{}, &models.RoleUser{})
	println("Database connected!")
	return db
}

// CloseDatabaseConnection method is closing a connection between your app and your db
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
