package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mohidex/mini-blog/controller"
	"github.com/mohidex/mini-blog/database"
	"github.com/mohidex/mini-blog/middleware"
	"github.com/mohidex/mini-blog/model"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.DB.AutoMigrate(&model.User{})
	database.DB.AutoMigrate(&model.Blog{})
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/blogs", controller.AddBlog)
	protectedRoutes.GET("/blogs", controller.GetAllBlog)
	protectedRoutes.GET("/blog/:id", controller.GetBlogById)

	router.Run(":8000")
	fmt.Println("Server running on port 8000")

}
