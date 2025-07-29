package main

import (
	"fmt"
	"go-task-api/database"
	"go-task-api/models"
	"go-task-api/routes"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	router := gin.Default()

	// ✅ Add CORS middleware here
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Replace * with your frontend domain in production []string{"https://myfrontend.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("JWT_SECRET")
	fmt.Println("Secret as string:", secret)
	fmt.Println("Secret as bytes:", []byte(secret))

	database.Connect()

	database.DB.AutoMigrate(&models.Task{}, &models.User{})

	routes.TaskRoutes(router)
	routes.AuthRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
