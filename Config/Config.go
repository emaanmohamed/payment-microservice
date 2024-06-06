package Config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Port              string
	DatabaseURL       string
	PaymentGatewayURL string
}

func InitConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return &Config{
		Port:              os.Getenv("PORT"),
		DatabaseURL:       os.Getenv("DATABASE_URL"),
		PaymentGatewayURL: os.Getenv("PAYMENT_GATEWAY"),
	}
}
