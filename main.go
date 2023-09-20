package main

import (
	"example/gin-webserver/database"
	authHandler "example/gin-webserver/handler/auth"
	entryHandler "example/gin-webserver/handler/entry"
	"example/gin-webserver/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDB()
	serveApp()
}

func loadDB() {
	database.Connect()
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env")
	}
}

func serveApp() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", authHandler.Register)
	publicRoutes.POST("/login", authHandler.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", entryHandler.PostEntry)
	protectedRoutes.GET("/entry", entryHandler.GetAllEntries)

	router.GET("/", serverHealth)

	router.Run(":8080")
	log.Println("Server running on port 8080")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func serverHealth(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Server is running"})
}
