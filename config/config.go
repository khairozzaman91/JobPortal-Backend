package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configuration *Config

type DBConfig struct {
	Host          string
	Port          string
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	ServiceName string
	Version     string
	HTTPPort    int
	JWTSecret   string
	DB          *DBConfig
}

func load() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	service := os.Getenv("SERVICENAME")
	if service == "" {
		log.Fatal("Service name is required")
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
	}

	portStr := os.Getenv("HTTPPORT")
	if portStr == "" {
		log.Fatal("HTTP Port is required")
	}

	httpPort, err := strconv.ParseInt(portStr, 10, 64)
	if err != nil {
		log.Fatal("Invalid HTTP port")
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET_KEY is required")
	}

	dbHost := os.Getenv("HOST")
	if dbHost == "" {
		log.Fatal("DB Host is required")
	}

	dbPort := os.Getenv("PORT")
	if dbPort == "" {
		log.Fatal("DB Port is required")
	}

	dbName := os.Getenv("NAME")
	if dbName == "" {
		log.Fatal("DB Name is required")
	}

	dbUser := os.Getenv("USER")
	if dbUser == "" {
		log.Fatal("DB User is required")
	}

	dbPass := os.Getenv("PASSWORD")
	if dbPass == "" {
		log.Fatal("DB Password is required")
	}

	enableSSLMode := os.Getenv("ENABLE_SSL_MODE")
	if enableSSLMode == "" {
		log.Fatal("ENABLE_SSL_MODE is required")
	}

	enableSSLMODE, err := strconv.ParseBool(enableSSLMode)
	if err != nil {
		fmt.Println("Invalid ENABLE_SSL_MODE value:", err)
		os.Exit(1)
	}

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          dbPort,
		Name:          dbName,
		User:          dbUser,
		Password:      dbPass,
		EnableSSLMODE: enableSSLMODE,
	}

	configuration = &Config{
		ServiceName: service,
		Version:     version,
		HTTPPort:    int(httpPort),
		JWTSecret:   jwtSecret,
		DB:          dbConfig,
	}
}

func GetConfig() *Config {

	if configuration == nil {
		load()
	}

	return configuration
}