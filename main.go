package main

import (
	"log"
	"os"

	"aniby/medods/rest"
	"aniby/medods/database"
	// "github.com/golang-jwt/jwt/v5"
	// "github.com/google/uuid"
	"github.com/go-pg/pg/v10"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file not found!")
	}
}

func getConfigProperty(name string) string {
	property, exists := os.LookupEnv(name)
	if !exists {
		log.Fatalf("%s not found!", name)
	}
	return property
}

func main() {
	// Config
	serverAddress := getConfigProperty("SERVER_ADDRESS")
	postgresAddress := getConfigProperty("POSTGRES_ADDRESS")
	postgresUser := getConfigProperty("POSTGRES_USER")
	postgresPassword := getConfigProperty("POSTGRES_PASSWORD")
	postgresDatabase := getConfigProperty("POSTGRES_DATABASE")
	jwtSecretKey := getConfigProperty("JWT_SECRET_KEY")

	// Database
	db_error := database.ReadyDatabase(&pg.Options{
		Addr: postgresAddress,
        User: postgresUser,
		Password: postgresPassword,
		Database: postgresDatabase,
    })
	if db_error != nil {
		log.Fatalf("Can't establish database connection!\n%v", db_error)
	}

	// Server
	rest.JwtSercetKey = jwtSecretKey
	router := gin.Default()
	router.GET("/:id", rest.GetTokens)
	router.PATCH("/", rest.PatchTokens)
	router.Run(serverAddress)
}
