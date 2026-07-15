package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type Config struct {
	ServiceName string
	Version     string
	HTTPPort    int
	JWTSecret   string
}

func load() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	service := os.Getenv("SERVICENAME")

	if service == "" {
		fmt.Println("service name required")
		os.Exit(1)
	}

	version := os.Getenv("VERSION")

	if version == "" {
		fmt.Println("version name required")
		os.Exit(1)
	}

	portStr := os.Getenv("HTTPPORT")
	if portStr == "" {
		fmt.Println("Port name required")
		os.Exit(1)
	}

	httpPort, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal("Invalid HTTP port")
		os.Exit(1)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET is required")
	}

	configuration = &Config{
		ServiceName: service,
		Version:     version,
		HTTPPort:    int(httpPort),
		JWTSecret: jwtSecret,
	}

}

func GetConfig() *Config {

	if configuration == nil {
		load()
	}
	return configuration
}
