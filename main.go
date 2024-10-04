package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"aniby/medods/database"
	"aniby/medods/rest"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
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
	accessTokenExpiresIn, err := strconv.Atoi(getConfigProperty("ACCESS_TOKEN_EXPIRES_IN_HOURS"))
	if err != nil {
		log.Fatalf("Parsing error:\n%v", err)
	}

	// Database
	db, err := database.ReadyDatabase(&pg.Options{
		Addr: postgresAddress,
        User: postgresUser,
		Password: postgresPassword,
		Database: postgresDatabase,
    })
	if err != nil {
		log.Fatalf("Can't establish database connection!\n%v", err)
	}
	fmt.Println("Database connection established")
	defer db.Close()

	// Server
	rest.JwtSercetKey = jwtSecretKey
	rest.AccessTokenExpiresIn = accessTokenExpiresIn
	router := gin.Default()
	router.GET("/:id", rest.GetTokens)
	router.PATCH("/", rest.PatchTokens)
	router.Run(serverAddress)
}
