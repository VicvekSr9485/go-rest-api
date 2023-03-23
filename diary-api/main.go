package main

import (
	"diary-api/database"
	"diary-api/model"
	"diary-api/controller"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
    "log"
	"fmt"

)

func main() {
	loadEnv()
    loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}
