package main

import (
	"log"
	// "fmt"
	"os"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/google/uuid"
	"aniby/medods/rest"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found!")
	}
}

func main() {
	// Config
	serverAddress, exists := os.LookupEnv("SERVER_ADDRESS")
	if !exists {
		log.Fatal("SERVER_ADDRESS not found!")
	}

	// Server
	router := gin.Default()
	router.GET("/", rest.GetTokens)
	router.POST("/", rest.PatchTokens)
	router.Run(serverAddress)
}
