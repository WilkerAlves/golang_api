package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/wilker/golang_api/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

//SetupDataBaseConnection is creating a new connection to our database
func SetupDataBaseConnection() *gorm.DB {
	errEnv := godotenv.Load()
	if errEnv != nil{
		panic("Failed to load env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, errCon := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errCon != nil{
		panic("Failed to create a connection to database")
	}

	//later we will fill in the model here
	db.AutoMigrate(&entity.Book{}, entity.User{})
	return db
}

//CloseDataBaseConnection method is closing a connection between your app and your db
func CloseDataBaseConnection(db *gorm.DB)  {
	dbSQL, err := db.DB()
	if err != nil{
		panic("Failed to close connection to database")
	}
	dbSQL.Close()
}