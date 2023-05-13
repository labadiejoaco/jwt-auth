package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/labadiejoaco/jwt-auth/controllers"
	"github.com/labadiejoaco/jwt-auth/database"
	"github.com/labadiejoaco/jwt-auth/middlewares"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()
}

func main() {

	r := gin.Default()

	r.POST("/api/login", controllers.Login)
	r.POST("/api/logout", controllers.Logout)
	r.POST("/api/signup", controllers.Signup)
	r.GET("/api/validate", middlewares.JwtAuthMiddleware, controllers.Validate)

	r.Run()
}
