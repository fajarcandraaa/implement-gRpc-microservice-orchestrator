package config

import (
	"fmt"
	"log"
	"os"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/config/routers"
	"github.com/joho/godotenv"
)

var server = routers.Serve{}

// init is used to loads values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

// Run is used to settup database connection
func Run() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	server.Initialize(
		os.Getenv("DB_DRIVER"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"))

	// seed.Load(server.DB)

	server.Run()

}
