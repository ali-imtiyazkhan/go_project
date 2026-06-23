package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI   string
	MongoDB    string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return nil, err
	}
	mongoURI, err := extractEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}
	mongoDB, err := extractEnv("MONGO_DB")
	if err != nil {
		return nil, err
	}
	serverPort, err := extractEnv("SERVER_PORT")
	if err != nil {
		return nil, err
	}
	return &Config{
		MongoURI:   mongoURI,
		MongoDB:    mongoDB,
		ServerPort: serverPort,
	}, nil
}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("environment variable %s is not set", key)
	}
	return val, nil
}
