package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/mohidex/mini-blog/database"
	"github.com/mohidex/mini-blog/model"
)

func main() {
	loadEnv()
	loadDatabase()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Blog{})
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}
